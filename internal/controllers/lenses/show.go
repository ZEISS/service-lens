package lenses

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// LensShowControllerImpl ...
type LensShowControllerImpl struct {
	lens  models.Lens
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensShowController ...
func NewLensShowController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *LensShowControllerImpl {
	return &LensShowControllerImpl{
		store: store,
	}
}

// Prepare ...
func (l *LensShowControllerImpl) Prepare() error {
	err := l.BindParams(&l.lens)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetLens(ctx, &l.lens)
	})
}

// Get ...
func (l *LensShowControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: l.Path(),
				User: l.Session().User,
			},
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"m-2": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text("Overview"),
					),
					htmx.Div(
						htmx.H1(
							htmx.Text(l.lens.Name),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("Description"),
							),
							htmx.H3(
								htmx.Text(l.lens.Description),
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("ID"),
							),
							htmx.H3(htmx.Text(l.lens.ID.String())),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("Created at"),
							),
							htmx.H3(
								htmx.Text(
									l.lens.CreatedAt.Format("2006-01-02 15:04:05"),
								),
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("Updated at"),
							),
							htmx.H3(
								htmx.Text(
									l.lens.UpdatedAt.Format("2006-01-02 15:04:05"),
								),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.OutlinePrimary(
							buttons.ButtonProps{},
							htmx.HxDelete(""),
							htmx.HxConfirm("Are you sure you want to delete this lens?"),
							htmx.Text("Delete"),
						),
					),
				),
			),
		),
	)
}

// // Delete ...
// func (l *LensIndexController) Delete() error {
// 	err := l.db.DestroyLens(l.Context(), l.params.ID)
// 	if err != nil {
// 		return err
// 	}

// 	l.Hx().Redirect(fmt.Sprintf("/teams/%s/lenses/list", l.params.Team))

// 	return nil
// }
