package usecase

import (
	"context"
	"naevis/internal/events/domain"
)

// EventUseCase defines the application-level operations for events.
type EventUseCase interface {
	Health(ctx context.Context) (string, error)
}

type eventUseCase struct{}

func NewEventUseCase() EventUseCase {
	return &eventUseCase{}
}

func (u *eventUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.Event{}
