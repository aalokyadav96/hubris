package usecase

import (
	"context"
	"naevis/internal/cart/domain"
)

// CartUseCase defines the application-level operations for carts.
type CartUseCase interface {
	Health(ctx context.Context) (string, error)
}

type cartUseCase struct{}

func NewCartUseCase() CartUseCase {
	return &cartUseCase{}
}

func (u *cartUseCase) Health(ctx context.Context) (string, error) {
	_ = ctx
	return "ok", nil
}

var _ = domain.Cart{}
