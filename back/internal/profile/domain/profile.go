package domain

import "errors"

// Profile represents a user profile.
type Profile struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrProfileNotFound = errors.New("profile not found")
	ErrInvalidProfile  = errors.New("invalid profile data")
)
