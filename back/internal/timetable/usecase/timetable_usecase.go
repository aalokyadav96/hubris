package usecase

import (
	"context"
	"naevis/internal/timetable/domain"
)

// TimetableUseCase defines the application-level operations for timetable slots.
type TimetableUseCase interface {
	Health(ctx context.Context) (string, error)
}

type timetableUseCase struct{}

func NewTimetableUseCase() TimetableUseCase {
	return &timetableUseCase{}
}

func (u *timetableUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.TimetableSlot{}
