package workloads

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// WorkloadDeleteControllerImpl ...
type WorkloadDeleteControllerImpl struct {
	workload models.Workload
	store    ports.Datastore
	htmx.DefaultController
}

// NewWorkloadDeleteController ...
func NewWorkloadDeleteController(store ports.Datastore) *WorkloadDeleteControllerImpl {
	return &WorkloadDeleteControllerImpl{
		store: store,
	}
}

// Prepare ...
func (p *WorkloadDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.workload)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteWorkload(ctx, &p.workload)
	})
}

// Delete ...
func (p *WorkloadDeleteControllerImpl) Delete() error {
	return p.Redirect("/workloads")
}
