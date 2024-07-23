package environments

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/environments"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// EnvironmentListControllerImpl ...
type EnvironmentListControllerImpl struct {
	environments tables.Results[models.Environment]
	store        seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.UnimplementedController
}

// NewEnvironmentListController ...
func NewEnvironmentListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EnvironmentListControllerImpl {
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
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Environments",
				Path:  w.Path(),
			},
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"my-2": true,
						"mx-2": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
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
	)
}
