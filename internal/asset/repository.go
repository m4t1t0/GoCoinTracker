package asset

import (
	"context"
)

// Repository defines the required persistence operations for tracked assets.
// Keep it small and focused on the use cases we need now.

type Repository interface {
	Create(ctx context.Context, name string, interval int) (TrackedAsset, error)
}
