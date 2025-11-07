package asset

import (
	"time"
)

// TrackedAsset is the domain model representing a row in tracked_assets.
// Keep this model simple and independent of frameworks (Clean Architecture).
// Database-specific details belong in the repository implementation.

type TrackedAsset struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Interval  int       `json:"interval"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
