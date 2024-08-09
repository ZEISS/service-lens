package workloads

import (
	"context"
	"net/http"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadUpdateControllerImpl ...
type WorkloadUpdateControllerImpl struct {
	workload models.Workload
	store    ports.Datastore
	htmx.DefaultController
}

// NewWorkloadUpdateController ...
func NewWorkloadUpdateController(store ports.Datastore) *WorkloadUpdateControllerImpl {
	return &WorkloadUpdateControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *WorkloadUpdateControllerImpl) Prepare() error {
	err := w.BindParams(&w.workload)
	if err != nil {
		return err
	}

	err = w.BindBody(&w.workload)
	if err != nil {
		return err
	}

	return w.store.ReadWriteTx(w.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateWorkload(ctx, &w.workload)
	})
}

// Put ...
func (w *WorkloadUpdateControllerImpl) Put() error {
	return w.Ctx().SendStatus(http.StatusNoContent)
}
