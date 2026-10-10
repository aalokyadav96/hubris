package domain

import "errors"

// Farm represents a farm entity.
type Farm struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrFarmNotFound = errors.New("farm not found")
	ErrInvalidFarm  = errors.New("invalid farm data")
)
