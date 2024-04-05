package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
)

// Profiles ...
type Profiles interface {
	// NewProfile creates a new profile.
	NewProfile(ctx context.Context, profile *models.Profile) error
	// FetchProfile fetches a profile by its ID.
	FetchProfile(ctx context.Context, id uuid.UUID) (*models.Profile, error)
	// ListProfiles lists all profiles.
	ListProfiles(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Profile, error)
	// GetProfileByID fetches a profile by its ID.
	GetProfileByID(ctx context.Context, id uuid.UUID) (*models.Profile, error)
	// UpdateProfile updates a profile.
	UpdateProfile(ctx context.Context, profile *models.Profile) error
	// DestroyProfile deletes a profile.
	DestroyProfile(ctx context.Context, id uuid.UUID) error
}
