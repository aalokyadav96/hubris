package usecase

import (
	"context"
	"naevis/internal/crops/domain"
)

// CropUseCase defines the application-level operations for crops.
type CropUseCase interface {
	Health(ctx context.Context) (string, error)
}

type cropUseCase struct{}

func NewCropUseCase() CropUseCase {
	return &cropUseCase{}
}

func (u *cropUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.Crop{}
