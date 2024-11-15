package dto

type CreateUserInput struct {
	WalletAddress string `json:"walletAddress"`
	Bio           string `json:"bio"`
}
