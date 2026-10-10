package usecase

import (
	"context"
	"naevis/internal/places/domain"
)

// PlaceRepository defines persistence operations for places.
type PlaceRepository interface {
	List(ctx context.Context) ([]domain.Place, error)
	GetByID(ctx context.Context, id int64) (domain.Place, error)
	Create(ctx context.Context, place domain.Place) (domain.Place, error)
	Update(ctx context.Context, id int64, place domain.Place) (domain.Place, error)
	Delete(ctx context.Context, id int64) error
}

// PlaceUseCase defines application-level operations for places.
type PlaceUseCase interface {
	List(ctx context.Context) ([]domain.Place, error)
	GetByID(ctx context.Context, id int64) (domain.Place, error)
	Create(ctx context.Context, place domain.Place) (domain.Place, error)
	Update(ctx context.Context, id int64, place domain.Place) (domain.Place, error)
	Delete(ctx context.Context, id int64) error
}

type placeUseCase struct{ repo PlaceRepository }

func NewPlaceUseCase(repo PlaceRepository) PlaceUseCase {
	return &placeUseCase{repo: repo}
}

func (u *placeUseCase) List(ctx context.Context) ([]domain.Place, error) {
	return u.repo.List(ctx)
}

func (u *placeUseCase) GetByID(ctx context.Context, id int64) (domain.Place, error) {
	if id <= 0 {
		return domain.Place{}, domain.ErrPlaceNotFound
	}
	return u.repo.GetByID(ctx, id)
}

func (u *placeUseCase) Create(ctx context.Context, place domain.Place) (domain.Place, error) {
	if err := place.Validate(); err != nil {
		return domain.Place{}, err
	}
	return u.repo.Create(ctx, place)
}

func (u *placeUseCase) Update(ctx context.Context, id int64, place domain.Place) (domain.Place, error) {
	if id <= 0 {
		return domain.Place{}, domain.ErrPlaceNotFound
	}
	if err := place.Validate(); err != nil {
		return domain.Place{}, err
	}
	return u.repo.Update(ctx, id, place)
}

func (u *placeUseCase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return domain.ErrPlaceNotFound
	}
	return u.repo.Delete(ctx, id)
}
