package domain

import "errors"

// TimetableSlot represents a scheduled time slot.
type TimetableSlot struct {
	ID string `json:"id,omitempty"`
}

var (
	ErrSlotNotFound = errors.New("time slot not found")
	ErrInvalidSlot  = errors.New("invalid timetable slot")
)
