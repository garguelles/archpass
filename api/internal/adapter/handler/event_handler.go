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
		ModifiedAt: event.ModifiedAt, StartDate: event.StartDate, EndDate: event.EndDate,
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
			Id:         event.ID,
			Name:       event.Name,
			ImageUrl:   event.ImageURL,
			EventSlug:  event.EventSlug,
			Location:   event.Location,
			StartDate:  event.StartDate,
			EndDate:    event.EndDate,
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

	idParam := c.QueryParam("id")
	if idParam == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'id' parameter is required."})
	}
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'id' parameter is invalid"})
	}

	eventRepo := repository.NewEventRepository(&ctx)
	event, err := eventRepo.GetByIdAndOrganizerId(id, claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.DashboardEvent{Id: event.ID,
		Name: event.Name, Description: event.Description, Location: event.Location,
		ImageUrl: event.ImageURL, EventSlug: event.EventSlug, CreatedAt: event.CreatedAt,
		ModifiedAt: event.ModifiedAt, StartDate: event.StartDate, EndDate: event.EndDate,
	})

}
