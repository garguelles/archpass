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

func CreateTicket(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "JWT Token missing or invalid."})
	}

	claims, ok := token.Claims.(*dto.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unable to cast claims."})
	}

	var input dto.CreateTicketInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}
	ctx := context.Background()
	ticketRepo := repository.NewTicketRepository(&ctx)

	ticket, err := ticketRepo.Create(input, claims.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.DashboardTicket{Id: ticket.ID,
		Name: ticket.Name, Description: ticket.Description, TicketSlug: ticket.TicketSlug,
		CreatedAt: ticket.CreatedAt, ModifiedAt: ticket.UpdatedAt, MintPrice: ticket.MintPrice,
		Quantity: ticket.Quantity,
	})
}

func ListDashboardTickets(c echo.Context) error {
	ctx := context.Background()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "JWT Token missing or invalid."})
	}

	claims, ok := token.Claims.(*dto.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unable to cast claims."})
	}

	eventIdParam := c.QueryParam("eventId")
	if eventIdParam == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'eventId' parameter is required."})
	}
	id, err := strconv.Atoi(eventIdParam)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'eventId' parameter is invalid"})
	}

	ticketRepo := repository.NewTicketRepository(&ctx)
	tickets, err := ticketRepo.ListByEventId(id, claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	ticketDtos := make([]dto.SimpleDashboardTicket, 0)

	for _, ticket := range tickets {
		ticketDto := dto.SimpleDashboardTicket{
			Id:   ticket.ID,
			Name: ticket.Name,
		}
		ticketDtos = append(ticketDtos, ticketDto)
	}

	return c.JSON(http.StatusOK, ticketDtos)
}
