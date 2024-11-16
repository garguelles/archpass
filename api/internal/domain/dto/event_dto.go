package dto

import (
	"time"
)

type CreateEventInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// StartDate   time.Time `json:"startDate"`
	// EndDate     time.Time `json:"endDate"`
	Date            string  `json:"date"`
	Location        string  `json:"location"`
	ImageUrl        *string `json:"imageUrl"`
	ContractAddress string  `json:"contractAddress"`
}

type DashboardEvent struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Location        string `json:"location"`
	ImageUrl        string `json:"imageUrl"`
	ContractAddress string `json:"contractAddress"`
	// StartDate   time.Time `json:"startDate"`
	// EndDate     time.Time `json:"endDate"`
	Date       string    `json:"date"`
	EventSlug  string    `json:"eventSlug"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type SimpleDashboardEvent struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Location        string `json:"location"`
	ImageUrl        string `json:"imageUrl"`
	ContractAddress string `json:"contractAddress"`
	// StartDate  time.Time `json:"startDate"`
	// EndDate    time.Time `json:"endDate"`
	Date       string    `json:"date"`
	EventSlug  string    `json:"eventSlug"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type PublicEvent struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description"`
	// StartDate       time.Time      `json:"startDate"`
	// EndDate         time.Time      `json:"endDate"`
	Date            string         `json:"date"`
	EventSlug       string         `json:"eventSlug"`
	ContractAddress string         `json:"contractAddress"`
	Tickets         []PublicTicket `json:"tickets"`
	Organizer       Organizer      `json:"organizer"`
}
