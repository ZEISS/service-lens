package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
)

// Lenses ...
type Lenses interface {
	GetLensByID(ctx context.Context, teamSlug string, id uuid.UUID) (*models.Lens, error)
	GetPillarById(ctx context.Context, teamSlug string, lensId uuid.UUID, id int) (*models.Pillar, error)
	AddLens(ctx context.Context, lens *models.Lens) (*models.Lens, error)
	ListLenses(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Lens, error)
	// DestroyLens is a method that deletes a lens from the database.
	DestroyLens(ctx context.Context, id uuid.UUID) error
}
