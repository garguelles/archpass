package handler

import (
	"context"

	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
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
		user, err = userRepo.Create(dto.CreateUserInput{WalletAddress: walletAddress, Bio: "Test Bio"})
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
