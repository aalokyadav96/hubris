package domain

import "errors"

// Crop represents a crop entity.
type Crop struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrCropNotFound = errors.New("crop not found")
	ErrInvalidCrop  = errors.New("invalid crop data")
)
