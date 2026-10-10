package usecase

import (
	"context"
	"naevis/internal/farms/domain"
)

// FarmUseCase defines the application-level operations for farms.
type FarmUseCase interface {
	Health(ctx context.Context) (string, error)
}

type farmUseCase struct{}

func NewFarmUseCase() FarmUseCase {
	return &farmUseCase{}
}

func (u *farmUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.Farm{}
