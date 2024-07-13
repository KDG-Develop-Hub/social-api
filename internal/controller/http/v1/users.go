package v1

import (
	"github.com/kdg-develop-hub/api/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type userRoutes struct {
	u service.UserService
	l *zerolog.Logger
}

func NewUserRoutes(e *echo.Group, u service.UserService, l *zerolog.Logger) {
	r := &userRoutes{u: u, l: l}
	g := e.Group("/users")
	g.POST("", r.create)
}

func (r *userRoutes) create(c echo.Context) error {
	//TODO: Implement crete user
	return nil
}
