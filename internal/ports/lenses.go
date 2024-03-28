package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
)

// Lenses ...
type Lenses interface {
	GetLensByID(ctx context.Context, id uuid.UUID) (*models.Lens, error)
	AddLens(ctx context.Context, lens *models.Lens) (*models.Lens, error)
	ListLenses(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Lens, error)
}
