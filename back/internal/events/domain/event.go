package domain

import "errors"

// Event represents a scheduled or operational event.
type Event struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrEventNotFound = errors.New("event not found")
	ErrInvalidEvent  = errors.New("invalid event data")
)
