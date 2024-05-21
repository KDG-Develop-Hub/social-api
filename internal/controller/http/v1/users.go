package v1

import (
	"errors"
	"github.com/kdg-develop-hub/api/internal/domain"
	"github.com/kdg-develop-hub/api/internal/service"
	"github.com/kdg-develop-hub/api/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
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
	request := domain.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
		Email:    c.FormValue("email"),
	}
	err := r.u.Register(request)
	if err != nil {
		if errors.Is(response.ErrResourceAlreadyExists, err) {
			res := response.WithErrors{Errors: []string{err.Error()}}
			res.Message = "User Already Exists"
			return c.JSON(http.StatusConflict, res)
		}
		r.l.Error().Err(err).Msg("Failed to register user")
		return c.JSON(500, response.Response{Message: "Internal Server Error"})
	}
	return c.JSON(201, response.Response{Message: "User Created"})
}
