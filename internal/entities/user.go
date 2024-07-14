package entities

import "time"

type User struct {
	ID          int
	Username    string
	DisplayName string
	Password    string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
