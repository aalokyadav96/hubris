package usecase

import (
	"context"
	"naevis/internal/auth/domain"
)

// AuthUseCase defines the application-level operations for auth.
type AuthUseCase interface {
	Health(ctx context.Context) (string, error)
}

type authUseCase struct{}

func NewAuthUseCase() AuthUseCase {
	return &authUseCase{}
}

func (u *authUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.Auth{}
