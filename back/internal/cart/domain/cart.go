package domain

import "errors"

// Cart represents a shopping cart item collection.
type Cart struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrCartNotFound = errors.New("cart not found")
	ErrCartEmpty    = errors.New("cart is empty")
)
