package workflows

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/cards"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/workflows"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

var _ = htmx.Controller(&ListWorkflowsControllerImpl{})

// ListWorkflowsControllerImpl ...
type ListWorkflowsControllerImpl struct {
	results tables.Results[models.Workflow]
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewListWorkflowsController ...
func NewListWorkflowsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ListWorkflowsControllerImpl {
	return &ListWorkflowsControllerImpl{store: store}
}

// Prepare ...
func (l *ListWorkflowsControllerImpl) Prepare() error {
	if err := l.BindQuery(&l.results); err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListWorkflows(ctx, &l.results)
	})
}

// Prepare ...
func (l *ListWorkflowsControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						workflows.WorkflowsTable(
							workflows.WorkflowsTableProps{
								Workflows: l.results.GetRows(),
								Offset:    l.results.GetOffset(),
								Limit:     l.results.GetLimit(),
								Total:     l.results.GetLen(),
							},
						),
					),
				)
			},
		),
	)
}
