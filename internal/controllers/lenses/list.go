package lenses

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/lenses"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// LensListController ...
type LensListController struct {
	lenses tables.Results[models.Lens]
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensListController ...
func NewLensListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *LensListController {
	return &LensListController{store: store}
}

// Prepare ...
func (w *LensListController) Prepare() error {
	err := w.BindQuery(&w.lenses)
	if err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListLenses(ctx, &w.lenses)
	})
}

// Get ...
func (w *LensListController) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: w.Path(),
				User: w.Session().User,
			},
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"m-2": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					lenses.LensesTable(
						lenses.LensesTableProps{
							Lenses: w.lenses.GetRows(),
							Offset: w.lenses.GetOffset(),
							Limit:  w.lenses.GetLimit(),
							Total:  w.lenses.GetTotalRows(),
						},
					),
				),
			),
		),
	)
}
