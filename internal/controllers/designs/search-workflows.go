package designs

import (
	"context"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/conv"
)

var _ = htmx.Controller(&SearchWorkflowsControllerImpl{})

// Search ...
type SearchWorkflowsControllerImpl struct {
	workflows tables.Results[models.Workflow]
	store     seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewSearchWorkflowsController ...
func NewSearchWorkflowsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *SearchWorkflowsControllerImpl {
	return &SearchWorkflowsControllerImpl{
		workflows: tables.Results[models.Workflow]{SearchFields: []string{"name"}},
		store:     store,
	}
}

// Error ...
func (l *SearchWorkflowsControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *SearchWorkflowsControllerImpl) Prepare() error {
	var params struct {
		Workflow string `json:"workflow" form:"workflow" query:"workflow" validate:"required"`
	}

	err := l.BindQuery(&params)
	if err != nil {
		return err
	}
	l.workflows.Search = params.Workflow

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListWorkflows(ctx, &l.workflows)
	})
}

// Get ...
func (l *SearchWorkflowsControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.ForEach(l.workflows.GetRows(), func(e *models.Workflow, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(conv.String(e.ID)),
					htmx.Text(conv.String(e.Name)),
				)
			})...,
		),
	)
}
