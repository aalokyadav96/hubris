package domain

import "time"

// Place is a domain entity for a named location in the system.
type Place struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Latitude    float64   `json:"latitude,omitempty"`
	Longitude   float64   `json:"longitude,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (p Place) Validate() error {
	if p.Name == "" {
		return ErrInvalidPlaceName
	}
	return nil
}
