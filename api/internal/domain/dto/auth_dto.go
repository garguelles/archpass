package dto

import "github.com/golang-jwt/jwt/v5"

type VerifyInput struct {
	Message   string  `json:"message"`
	Signature string  `json:"signature"`
	Nonce     *string `json:"nonce"`
}

type JWTCustomClaims struct {
	Id            int    `json:"id"`
	WalletAddress string `json:"walletAddress"`
	jwt.RegisteredClaims
}
