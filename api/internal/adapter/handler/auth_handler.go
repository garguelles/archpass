package handler

import (
	// "bytes"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	// "encoding/hex"
	"fmt"

	"log"
	"net/http"
	"os"
	// "strings"
	// "time"

	// "github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/log"
	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/internal/adapter/repository"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/spruceid/siwe-go"
)

func init() {
	godotenv.Load(".env")
}

// ERC1271ABI is the ABI for the isValidSignature function
const ERC1271ABI = `[{"inputs":[{"internalType":"bytes32","name":"_hash","type":"bytes32"},{"internalType":"bytes","name":"_signature","type":"bytes"}],"name":"isValidSignature","outputs":[{"internalType":"bytes4","name":"magicValue","type":"bytes4"}],"stateMutability":"view","type":"function"}]`

// ERC1271MagicValue is the expected return value for a valid signature
var ERC1271MagicValue = [4]byte{0x16, 0x26, 0xba, 0x7e} // 0x1626ba7e

func createMessageHash(message string) (common.Hash, error) {
	// Create the EIP-191 version byte + version-specific data
	prefix := "\x19Ethereum Signed Message:\n"
	messageLength := fmt.Sprintf("%d", len(message))

	// Concatenate the prefix, message length, and message
	prefixedMessage := fmt.Sprintf("%s%s%s", prefix, messageLength, message)

	// Hash the prefixed message
	return crypto.Keccak256Hash([]byte(prefixedMessage)), nil
}

// SmartWallet is the interface for the ERC-4337 compatible smart wallet
type SmartWallet interface {
	IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error)
}

type VerifyInput struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Nonce     string `json:"nonce"`
}

func Nonce(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"nonce": siwe.GenerateNonce()})
}

func Verify(c echo.Context) error {
	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		log.Fatal("Cannot load configuration")
	}

	var input dto.VerifyInput
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	message, err := siwe.ParseMessage(input.Message)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	domain := message.GetDomain()
	publicKey, err := message.Verify(input.Signature, &domain, input.Nonce, nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}
	walletAddress := crypto.PubkeyToAddress(*publicKey).Hex()
	ctx := context.Background()
	userRepo := repository.NewUserRepository(&ctx)

	var user ent.User
	user, err = userRepo.FindByWalletAddress(walletAddress)
	if ent.IsNotFound(err) {
		user, err = userRepo.Create(dto.CreateUserInput{WalletAddress: walletAddress})
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		}
	}

	claims := &dto.JWTCustomClaims{
		WalletAddress: user.WalletAddress,
		Id:            user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"accessToken": accessToken})
}
