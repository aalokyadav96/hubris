package domain

import "errors"

var (
	ErrInvalidPlaceName = errors.New("place name is required")
	ErrPlaceNotFound    = errors.New("place not found")
)
