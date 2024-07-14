package v1

import (
	"github.com/kdg-develop-hub/api/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type userRoutes struct {
	u services.UserService
	l *zerolog.Logger
}

func NewUserRoutes(e *echo.Group, u services.UserService, l *zerolog.Logger) {
	r := &userRoutes{u: u, l: l}
	g := e.Group("/users")
	g.POST("", r.create)
}

func (r *userRoutes) create(c echo.Context) error {
	//TODO: Implement crete user
	return nil
}
