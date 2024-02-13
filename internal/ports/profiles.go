package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
)

// Profiles ...
type Profiles interface {
	NewProfile(ctx context.Context, profile *models.Profile) error
	FetchProfile(ctx context.Context, id uuid.UUID) (*models.Profile, error)
}
