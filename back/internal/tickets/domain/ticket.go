package domain

import "errors"

// Ticket represents a support or ticket record.
type Ticket struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrTicketNotFound = errors.New("ticket not found")
	ErrInvalidTicket  = errors.New("invalid ticket data")
)
