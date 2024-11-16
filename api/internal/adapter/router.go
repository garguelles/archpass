package adapter

import (
	"log"
	"net/http"
	"os"

	"github.com/garguelles/archpass/internal/adapter/handler"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load(".env")
}

func Router() *echo.Echo {
	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		log.Fatal("Cannot load configuration")
	}

	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Gzip())
	r.Use(middleware.CORS())

	// JWT Config
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(dto.JWTCustomClaims)
		},
		SigningKey: []byte(secretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unauthorized"})
		},
	}

	// Public routes
	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"Status": "ok"})
	})

	r.GET("/health", handler.Health)

	// Auth handlers
	r.GET("/nonce", handler.Nonce)
	r.POST("/verify", handler.Verify)

	// Public event
	r.GET("/event.get", handler.GetEvent)

	// Private routes
	p := r.Group("")
	p.Use(echojwt.WithConfig(config))

	// Events
	p.POST("/admin/event.create", handler.CreateEvent)
	p.GET("/admin/event.list", handler.ListDashboardEvents)
	p.GET("/admin/event.get", handler.GetDashboardEvent)

	// Tickets
	p.POST("/admin/ticket.create", handler.CreateTicket)
	p.GET("/admin/ticket.list", handler.ListDashboardTickets)

	return r
}
