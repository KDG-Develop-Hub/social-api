package response

import "errors"

var (
	ErrResourceNotFound      = errors.New("resource not found")
	ErrResourceAlreadyExists = errors.New("resource already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrInternalServer        = errors.New("internal server error")
)
