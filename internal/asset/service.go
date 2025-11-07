package asset

import (
	"context"
	"strings"
)

// Service contains business logic around tracked assets.
// Keep it minimal; validations beyond HTTP-level can go here later.

type Service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return Service{repo: r}
}

// Create registers a new tracked asset.
// It normalizes the name to uppercase and trims spaces.
func (s Service) Create(ctx context.Context, name string, interval int) (TrackedAsset, error) {
	clean := strings.TrimSpace(strings.ToUpper(name))
	return s.repo.Create(ctx, clean, interval)
}
