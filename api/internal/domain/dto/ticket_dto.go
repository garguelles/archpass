package dto

import "time"

type CreateTicketInput struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	EventId         int    `json:"eventId"`
	Quantity        int    `json:"quantity"`
	MintPrice       string `json:"mintPrice"`
	ImageUrl        string `json:"imageUrl"`
	ContractAddress string `json:"contractAddress"`
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
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	MintPrice string `json:"mintPrice"`
	ImageUrl  string `json:"imageUrl"`
}

type PublicTicket struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	TicketSlug      string `json:"ticketSlug"`
	MintPrice       string `json:"mintPrice"`
	ContractAddress string `json:"contractAddress"`
}

type SimplePublicTicket struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TicketSlug  string `json:"ticketSlug"`
	MintPrice   string `json:"mintPrice"`
	EventSlug   string `json:"eventSlug"`
}

type AttendeeSimpleTicket struct {
	TicketSlug string `json:"ticketSlug"`
	EventName  string `json:"eventName"`
}

type AttendeeTicket struct {
	DisplayName   string                 `json:"displayName"`
	Bio           string                 `json:"bio"`
	WalletAddress string                 `json:"walletAddress"`
	Tickets       []AttendeeSimpleTicket `json:"tickets"`
}
