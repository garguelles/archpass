package dto

import "time"

type CreateEventInput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Location    string    `json:"location"`
	ImageUrl    *string   `json:"imageUrl"`
}

type DashboardEvent struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	ImageUrl    string    `json:"imageUrl"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	EventSlug   string    `json:"eventSlug"`
	CreatedAt   time.Time `json:"createdAt"`
	ModifiedAt  time.Time `json:"modifiedAt"`
}

type SimpleDashboardEvent struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Location   string    `json:"location"`
	ImageUrl   string    `json:"imageUrl"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	EventSlug  string    `json:"eventSlug"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}