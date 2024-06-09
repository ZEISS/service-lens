package environments

import (
	"context"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/environments"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// EnvironmentListControllerImpl ...
type EnvironmentListControllerImpl struct {
	environments tables.Results[models.Environment]
	store        ports.Datastore
	htmx.UnimplementedController
}

// NewEnvironmentListController ...
func NewEnvironmentListController(store ports.Datastore) *EnvironmentListControllerImpl {
	return &EnvironmentListControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *EnvironmentListControllerImpl) Prepare() error {
	err := w.BindQuery(&w.environments)
	if err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListEnvironments(ctx, &w.environments)
	})
}

// Get ...
func (w *EnvironmentListControllerImpl) Get() error {
	return w.Render(
		components.Page(
			components.PageProps{},
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
						environments.EnvironmentsTable(
							environments.EnvironmentsTableProps{
								Environments: w.environments.GetRows(),
								Offset:       w.environments.GetOffset(),
								Limit:        w.environments.GetLimit(),
								Total:        w.environments.GetLen(),
							},
						),
					),
				),
			),
		),
	)
}
