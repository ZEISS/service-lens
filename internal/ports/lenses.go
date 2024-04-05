package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
)

// Lenses ...
type Lenses interface {
	// GetLensByID is a method that retrieves a lens from the database by its ID.
	GetLensByID(ctx context.Context, id uuid.UUID) (*models.Lens, error)
	// GetPillarById is a method that retrieves a pillar from the database by its ID.
	GetPillarById(ctx context.Context, teamSlug string, lensId uuid.UUID, id int) (*models.Pillar, error)
	AddLens(ctx context.Context, lens *models.Lens) (*models.Lens, error)
	ListLenses(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Lens, error)
	// DestroyLens is a method that deletes a lens from the database.
	DestroyLens(ctx context.Context, id uuid.UUID) error
}
