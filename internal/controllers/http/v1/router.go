package v1

import (
	"github.com/kdg-develop-hub/api/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func NewRouter(e *echo.Echo, log *zerolog.Logger, u services.UserService) {
	g := e.Group("/v1")
	NewUserRoutes(g, u, log)
	g.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
}
