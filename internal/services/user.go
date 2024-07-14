package services

import (
	"github.com/kdg-develop-hub/api/internal/entities"
	"github.com/kdg-develop-hub/api/internal/repositories"
)

type (
	userService struct {
		ur repositories.UserRepository
	}
	UserService interface {
		FindByID(id int) (entities.User, error)
		SearchByUsername(username string) ([]entities.User, error)
		Register(uname, displayName, email, password string) error
	}
)

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{ur: repo}
}

func (u userService) FindByID(id int) (entities.User, error) {
	return u.ur.FindByID(id)
}

func (u userService) SearchByUsername(username string) ([]entities.User, error) {
	return u.ur.SearchByUsername(username)
}

func (u userService) Register(username, displayName, email, password string) error {
	return u.ur.Create(username, displayName, email, password)
}
