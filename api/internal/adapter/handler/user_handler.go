package handler

import (
	"context"
	"net/http"

	"github.com/garguelles/archpass/internal/adapter/repository"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var input dto.CreateUserInput
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}
	ctx := context.Background()
	userRepo := repository.NewUserRepository(&ctx)

	user, err := userRepo.Create(input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}
