package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/garguelles/archpass/internal/adapter/repository"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateEvent(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "JWT Token missing or invalid."})
	}

	claims, ok := token.Claims.(*dto.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unable to cast claims."})
	}

	var input dto.CreateEventInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}
	ctx := context.Background()
	eventRepo := repository.NewEventRepository(&ctx)

	event, err := eventRepo.Create(input, claims.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.DashboardEvent{Id: event.ID,
		Name: event.Name, Description: event.Description, Location: event.Location,
		ImageUrl: event.ImageURL, EventSlug: event.EventSlug, CreatedAt: event.CreatedAt,
		ModifiedAt: event.ModifiedAt, Date: event.Date,
	})
}

func ListDashboardEvents(c echo.Context) error {
	ctx := context.Background()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "JWT Token missing or invalid."})
	}

	claims, ok := token.Claims.(*dto.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unable to cast claims."})
	}

	defaultLimit := 10
	defaultOffset := 0

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset <= 0 {
		offset = defaultOffset
	}

	eventRepo := repository.NewEventRepository(&ctx)
	events, err := eventRepo.ListByOrganizerId(limit, offset, claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	eventDtos := make([]dto.SimpleDashboardEvent, 0)

	for _, event := range events {
		eventDto := dto.SimpleDashboardEvent{
			Id:        event.ID,
			Name:      event.Name,
			ImageUrl:  event.ImageURL,
			EventSlug: event.EventSlug,
			Location:  event.Location,
			// StartDate:  event.StartDate,
			// EndDate:    event.EndDate,
			Date:       event.Date,
			CreatedAt:  event.CreatedAt,
			ModifiedAt: event.ModifiedAt,
		}
		eventDtos = append(eventDtos, eventDto)
	}

	return c.JSON(http.StatusOK, eventDtos)
}

func GetDashboardEvent(c echo.Context) error {
	ctx := context.Background()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "JWT Token missing or invalid."})
	}

	claims, ok := token.Claims.(*dto.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unable to cast claims."})
	}

	slug := c.QueryParam("slug")
	if slug == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'id' parameter is required."})
	}

	eventRepo := repository.NewEventRepository(&ctx)
	event, err := eventRepo.GetBySlugAndOrganizerId(slug, claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.DashboardEvent{Id: event.ID,
		Name: event.Name, Description: event.Description, Location: event.Location,
		ImageUrl: event.ImageURL, EventSlug: event.EventSlug, CreatedAt: event.CreatedAt,
		ModifiedAt: event.ModifiedAt, Date: event.Date,
	})

}

func GetEvent(c echo.Context) error {
	ctx := context.Background()

	slugParam := c.QueryParam("slug")
	if len(slugParam) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'slug' parameter is required."})
	}

	eventRepo := repository.NewEventRepository(&ctx)
	event, err := eventRepo.GetBySlug(slugParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	eventDto := dto.PublicEvent{
		Name:     event.Name,
		Location: event.Location,
		ImageUrl: event.ImageURL,
		// StartDate:       event.StartDate,
		// EndDate:         event.EndDate,
		Date:            event.Date,
		EventSlug:       event.EventSlug,
		ContractAddress: event.ContractAddress,
		Organizer: dto.Organizer{
			WalletAddress: event.Edges.User.WalletAddress,
			Bio:           event.Edges.User.Bio,
			DisplayName:   event.Edges.User.DisplayName,
		},
		Tickets: []dto.PublicTicket{},
	}

	for _, ticket := range event.Edges.Tickets {
		eventDto.Tickets = append(eventDto.Tickets, dto.PublicTicket{
			Name:            ticket.Name,
			Description:     ticket.Description,
			TicketSlug:      ticket.TicketSlug,
			MintPrice:       ticket.MintPrice,
			ContractAddress: ticket.ContractAddress,
		})
	}

	return c.JSON(http.StatusOK, eventDto)
}
