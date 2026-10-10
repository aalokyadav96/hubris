package domain

import "errors"

// Auth is the domain model for authentication concerns.
type Auth struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
)
