package usecase

import (
	"context"
	"naevis/internal/tickets/domain"
)

// TicketUseCase defines the application-level operations for tickets.
type TicketUseCase interface {
	Health(ctx context.Context) (string, error)
}

type ticketUseCase struct{}

func NewTicketUseCase() TicketUseCase {
	return &ticketUseCase{}
}

func (u *ticketUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.Ticket{}
