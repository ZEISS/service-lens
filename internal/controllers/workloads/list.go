package workloads

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/workloads"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// WorkloadListControllerImpl ...
type WorkloadListControllerImpl struct {
	workloads tables.Results[models.Workload]
	store     seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadListController ...
func NewWorkloadListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadListControllerImpl {
	return &WorkloadListControllerImpl{store: store}
}

// Prepare ...
func (w *WorkloadListControllerImpl) Prepare() error {
	if err := w.BindQuery(&w.workloads); err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListWorkloads(ctx, &w.workloads)
	})
}

// Get ...
func (w *WorkloadListControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: w.Path(),
				User: w.Session().User,
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
						workloads.WorkloadsTable(
							workloads.WorkloadsTableProps{
								Workloads: w.workloads.GetRows(),
								Offset:    w.workloads.GetOffset(),
								Limit:     w.workloads.GetLimit(),
								Total:     w.workloads.GetTotalRows(),
							},
						),
					),
				)
			},
		),
	)
}
