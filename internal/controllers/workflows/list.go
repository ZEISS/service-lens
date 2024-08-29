package workflows

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
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
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewListWorkflowsController ...
func NewListWorkflowsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ListWorkflowsControllerImpl {
	return &ListWorkflowsControllerImpl{
		store: store,
	}
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
				results := tables.Results[models.Workflow]{SearchFields: []string{"Name"}}

				errorx.Panic(l.BindQuery(&results))
				errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListWorkflows(ctx, &results)
				}))

				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						workflows.WorkflowsTable(
							workflows.WorkflowsTableProps{
								Workflows: results.GetRows(),
								Offset:    results.GetOffset(),
								Limit:     results.GetLimit(),
								Total:     results.GetLen(),
								URL:       l.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
