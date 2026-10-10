package domain

import placesdomain "naevis/internal/places/domain"

// Place is the compatibility export for the canonical place domain model.
type Place = placesdomain.Place

var (
	ErrInvalidPlaceName = placesdomain.ErrInvalidPlaceName
	ErrPlaceNotFound    = placesdomain.ErrPlaceNotFound
)
