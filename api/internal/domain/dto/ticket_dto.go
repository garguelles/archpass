package dto

import "time"

type CreateTicketInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EventId     int    `json:"eventId"`
	Quantity    int    `json:"quantity"`
	MintPrice   string `json:"mintPrice"`
}

type DashboardTicket struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TicketSlug  string    `json:"ticketSlug"`
	MintPrice   string    `json:"mintPrice"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
	ModifiedAt  time.Time `json:"modifiedAt"`
}

type SimpleDashboardTicket struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
