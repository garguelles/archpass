package dto

type CreateAttendeeInput struct {
	TokenId         int    `json:"tokenId"`
	UserId          int    `json:"userId"`
	EventId         int    `json:"eventId"`
	TicketId        int    `json:"ticketId"`
	TransactionHash string `json:"transactionHash"`
	BlockNumber     int64  `json:"blockNumber"`
}
