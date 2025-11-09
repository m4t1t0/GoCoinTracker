package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/m4t1t0/GoCoinTracker/internal/asset"
	"gorm.io/gorm"
)

// Repository implements asset.Repository using PostgreSQL via GORM.
 type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return Repository{db: db}
}

type trackedAssetEntity struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Interval  int       `gorm:"column:interval;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (trackedAssetEntity) TableName() string { return "tracked_assets" }

func (r Repository) Create(ctx context.Context, name string, interval int) (asset.TrackedAsset, error) {
	ent := trackedAssetEntity{
		ID:       uuid.NewString(),
		Name:     name,
		Interval: interval,
	}
	if err := r.db.WithContext(ctx).Create(&ent).Error; err != nil {
		return asset.TrackedAsset{}, err
	}
	return asset.TrackedAsset{
		ID:        ent.ID,
		Name:      ent.Name,
		Interval:  ent.Interval,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}, nil
}
