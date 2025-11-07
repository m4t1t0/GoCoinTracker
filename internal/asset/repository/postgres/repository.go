package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/m4t1t0/GoCoinTracker/internal/asset"
)

// Repository implements asset.Repository using PostgreSQL via database/sql.
type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(ctx context.Context, name string, interval int) (asset.TrackedAsset, error) {
	id := uuid.NewString()
	const q = `INSERT INTO tracked_assets (id, name, interval)
		VALUES ($1, $2, $3)
		RETURNING id, name, interval, created_at, updated_at`

	var out asset.TrackedAsset
	var createdAt, updatedAt time.Time
	row := r.db.QueryRowContext(ctx, q, id, name, interval)
	if err := row.Scan(&out.ID, &out.Name, &out.Interval, &createdAt, &updatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return asset.TrackedAsset{}, nil
		}
		return asset.TrackedAsset{}, err
	}
	out.CreatedAt = createdAt
	out.UpdatedAt = updatedAt
	return out, nil
}
