package workloads

import (
	"context"
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

var validate *validator.Validate

// CreateWorkloadControllerImpl ...
type CreateWorkloadControllerImpl struct {
	workload models.Workload
	store    ports.Datastore
	htmx.DefaultController
}

// NewCreateWorkloadController ...
func NewCreateWorkloadController(store ports.Datastore) *CreateWorkloadControllerImpl {
	return &CreateWorkloadControllerImpl{store: store}
}

// Prepare ...
func (l *CreateWorkloadControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.workload)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.workload)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateWorkload(ctx, &l.workload)
	})
}

// Post ...
func (l *CreateWorkloadControllerImpl) Post() error {
	return l.Redirect(fmt.Sprintf("/workloads/%s", l.workload.ID))
}
