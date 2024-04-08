package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
)

// Environments ...
type Environments interface {
	// NewEnvironment creates a new profile.
	NewEnvironment(ctx context.Context, environment *models.Environment) error
	// ListEnvironment lists all profiles.
	ListEnvironment(ctx context.Context, teamSlug string, pagination models.Pagination[*models.Environment]) (*models.Pagination[*models.Environment], error)
	// GetEnvironment by ID.
	GetEnvironment(ctx context.Context, id uuid.UUID) (*models.Environment, error)
	// UpdateEnvironment updates a profile.
	UpdateEnvironment(ctx context.Context, profile *models.Environment) error
	// DeleteEnvironment deletes a profile.
	DeleteEnvironment(ctx context.Context, id uuid.UUID) error
}
