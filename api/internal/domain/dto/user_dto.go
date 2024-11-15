package dto

type CreateUserInput struct {
	WalletAddress string `json:"walletAddress"`
	Bio           string `json:"bio"`
}

type Organizer struct {
	WalletAddress string `json:"walletAddress"`
	Bio           string `json:"bio"`
	DisplayName   string `json:"displayName"`
}
