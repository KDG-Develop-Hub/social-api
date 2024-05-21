package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/kdg-develop-hub/api/internal/domain"
	"github.com/kdg-develop-hub/api/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	db *sqlx.DB
}

type UserRepository interface {
	Create(username, displayName, email, password string) error
	Authenticate(email, password string) error
	FindByID(id int) (domain.User, error)
	SearchByUsername(username string) ([]domain.User, error)
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(username, displayName, password, email string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrInternalServer
	}
	_, err = r.db.Exec("INSERT INTO users (username, displayName ,password, email) VALUES ($1, $2, $3, $4)", username, displayName, hashed, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrResourceAlreadyExists
		}
		return response.ErrInternalServer
	}
	return nil
}

func (r *userRepo) Authenticate(email, password string) error {
	var u domain.User
	err := r.db.Get(&u, "SELECT password FROM users WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrInvalidCredentials
		}
		return response.ErrInternalServer
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return response.ErrInvalidCredentials
		}
		return response.ErrInternalServer
	}
	return nil
}

func (r *userRepo) FindByID(id int) (domain.User, error) {
	var u domain.User
	err := r.db.Get(&u, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, response.ErrResourceNotFound
		}
		return u, response.ErrInternalServer
	}
	return u, nil
}

func (r *userRepo) SearchByUsername(username string) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Select(&users, "SELECT * FROM users WHERE username LIKE $1", "%"+username+"%")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, response.ErrResourceNotFound
		}
		return users, response.ErrInternalServer
	}
	return users, nil
}
