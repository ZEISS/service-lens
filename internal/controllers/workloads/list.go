package workloads

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/workloads"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// WorkloadListControllerImpl ...
type WorkloadListControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadListController ...
func NewWorkloadListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadListControllerImpl {
	return &WorkloadListControllerImpl{store: store}
}

// Get ...
func (w *WorkloadListControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        w.Path(),
				User:        w.Session().User,
				Development: w.IsDevelopment(),
			},
			func() htmx.Node {
				results := tables.Results[models.Workload]{SearchFields: []string{"name"}}

				errorx.Panic(w.BindQuery(&results))
				errorx.Panic(w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListWorkloads(ctx, &results)
				}))

				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						workloads.WorkloadsTable(
							workloads.WorkloadsTableProps{
								Workloads: results.GetRows(),
								Offset:    results.GetOffset(),
								Limit:     results.GetLimit(),
								Total:     results.GetTotalRows(),
								URL:       w.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
