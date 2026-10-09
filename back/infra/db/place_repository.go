package db

import (
	"context"
	"errors"
	"naevis/internal/domain"
	"naevis/internal/usecase"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PlaceRepository persists places in PostgreSQL.
type PlaceRepository struct { pool *pgxpool.Pool }

func NewPlaceRepository(pool *pgxpool.Pool) usecase.PlaceRepository {
	if pool == nil {
		return nil
	}
	return &PlaceRepository{pool: pool}
}

func (r *PlaceRepository) List(ctx context.Context) ([]domain.Place, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, name, description, latitude, longitude, created_at, updated_at
		FROM places
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	places := make([]domain.Place, 0)
	for rows.Next() {
		var p domain.Place
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Latitude, &p.Longitude, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		places = append(places, p)
	}
	return places, rows.Err()
}

func (r *PlaceRepository) GetByID(ctx context.Context, id int64) (domain.Place, error) {
	var p domain.Place
	if err := r.pool.QueryRow(ctx, `
		SELECT id, name, description, latitude, longitude, created_at, updated_at
		FROM places
		WHERE id = $1
	`, id).Scan(&p.ID, &p.Name, &p.Description, &p.Latitude, &p.Longitude, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Place{}, domain.ErrPlaceNotFound
		}
		return domain.Place{}, err
	}
	return p, nil
}

func (r *PlaceRepository) Create(ctx context.Context, place domain.Place) (domain.Place, error) {
	if place.Name == "" {
		return domain.Place{}, domain.ErrInvalidPlaceName
	}

	p := domain.Place{
		Name:        place.Name,
		Description: place.Description,
		Latitude:    place.Latitude,
		Longitude:   place.Longitude,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := r.pool.QueryRow(ctx, `
		INSERT INTO places (name, description, latitude, longitude, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`, p.Name, p.Description, p.Latitude, p.Longitude, p.CreatedAt, p.UpdatedAt).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return domain.Place{}, err
	}
	return p, nil
}

func (r *PlaceRepository) Update(ctx context.Context, id int64, place domain.Place) (domain.Place, error) {
	if place.Name == "" {
		return domain.Place{}, domain.ErrInvalidPlaceName
	}

	p := domain.Place{
		ID:          id,
		Name:        place.Name,
		Description: place.Description,
		Latitude:    place.Latitude,
		Longitude:   place.Longitude,
		UpdatedAt:   time.Now().UTC(),
	}

	if err := r.pool.QueryRow(ctx, `
		UPDATE places
		SET name = $1, description = $2, latitude = $3, longitude = $4, updated_at = $5
		WHERE id = $6
		RETURNING id, name, description, latitude, longitude, created_at, updated_at
	`, p.Name, p.Description, p.Latitude, p.Longitude, p.UpdatedAt, id).Scan(&p.ID, &p.Name, &p.Description, &p.Latitude, &p.Longitude, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Place{}, domain.ErrPlaceNotFound
		}
		return domain.Place{}, err
	}
	return p, nil
}

func (r *PlaceRepository) Delete(ctx context.Context, id int64) error {
	commandTag, err := r.pool.Exec(ctx, `DELETE FROM places WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return domain.ErrPlaceNotFound
	}
	return nil
}

func (r *PlaceRepository) EnsureTable(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS places (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			latitude DOUBLE PRECISION,
			longitude DOUBLE PRECISION,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	return err
}

var _ usecase.PlaceRepository = (*PlaceRepository)(nil)
