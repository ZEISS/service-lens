package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
)

// Workloads ...
type Workloads interface {
	ListWorkloads(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Workload, error)
	IndexWorkload(ctx context.Context, id uuid.UUID) (*models.Workload, error)
	StoreWorkload(ctx context.Context, workload *models.Workload) error
	DestroyWorkload(ctx context.Context, id uuid.UUID) error
}
