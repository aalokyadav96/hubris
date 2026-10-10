package usecase

import (
	"context"
	"naevis/internal/profile/domain"
)

// ProfileUseCase defines the application-level operations for profiles.
type ProfileUseCase interface {
	Health(ctx context.Context) (string, error)
}

type profileUseCase struct{}

func NewProfileUseCase() ProfileUseCase {
	return &profileUseCase{}
}

func (u *profileUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.Profile{}
