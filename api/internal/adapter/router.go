package adapter

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pragma-collective/0xStarter-api/internal/adapter/handler"
)

func Router() *echo.Echo {
	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Gzip())

	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"Status": "ok"})
	})

	r.GET("/health", handler.Health)
	r.POST("/user.create", handler.CreateUser)

	return r
}
