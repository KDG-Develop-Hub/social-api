package service

import (
	"github.com/kdg-develop-hub/api/internal/domain"
	"github.com/kdg-develop-hub/api/internal/repository"
)

type (
	userService struct {
		ur repository.UserRepository
	}
	UserService interface {
		FindByID(id int) (domain.User, error)
		SearchByUsername(username string) ([]domain.User, error)
		Register(uname, displayName, email, password string) error
	}
)

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{ur: repo}
}

func (u userService) FindByID(id int) (domain.User, error) {
	return u.ur.FindByID(id)
}

func (u userService) SearchByUsername(username string) ([]domain.User, error) {
	return u.ur.SearchByUsername(username)
}

func (u userService) Register(username, displayName, email, password string) error {
	return u.ur.Create(username, displayName, email, password)
}
