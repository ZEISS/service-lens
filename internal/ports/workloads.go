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
	// CreateWorkload creates a new workload.
	CreateWorkload(ctx context.Context, workload *models.Workload) error
	// Destroy a workload by its ID.
	DestroyWorkload(ctx context.Context, id uuid.UUID) error
	ListAnswers(ctx context.Context, id uuid.UUID, lensID uuid.UUID, questionID int) (*models.WorkloadLensQuestionAnswer, error)
	UpdateAnswers(ctx context.Context, id uuid.UUID, lensID uuid.UUID, questionID int, choices []int, doesNotApply bool, notes string) error
}
