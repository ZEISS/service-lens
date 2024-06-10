package workloads

import (
	"context"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/workloads"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// WorkloadListControllerImpl ...
type WorkloadListControllerImpl struct {
	workloads tables.Results[models.Workload]
	store     ports.Datastore
	htmx.DefaultController
}

// NewWorkloadListController ...
func NewWorkloadListController(store ports.Datastore) *WorkloadListControllerImpl {
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
		components.Page(
			components.PageProps{
				Title: "Workloads",
			},
			components.Layout(
				components.LayoutProps{
					Path: w.Path(),
				},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						workloads.WorkloadsTable(
							workloads.WorkloadsTableProps{
								Workloads: w.workloads.GetRows(),
								Offset:    w.workloads.GetOffset(),
								Limit:     w.workloads.GetLimit(),
								Total:     w.workloads.GetTotalRows(),
							},
						),
					),
				),
			),
		),
	)
}
