package dto

type CreateEventInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	ImageUrl    string `json:"imageUrl"`
}
