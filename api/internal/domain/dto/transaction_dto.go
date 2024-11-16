package dto

type CreateTransactionInput struct {
	EventType       string `json:"name"`
	WalletAddress   string `json:"walletAddress"`
	TransactionHash string `json:"transactionHash"`
	BlockNumber     int64  `json:"blockNumber"`
}
