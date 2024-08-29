package lenses

import (
	"context"
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/lenses"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
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
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
			},
			func() htmx.Node {
				return htmx.Fragment(

					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Overview"),
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
									htmx.Text("Name"),
								),
								htmx.H3(
									htmx.Text(l.lens.Name),
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
									htmx.Text("Version"),
								),
								htmx.H3(
									htmx.Text(conv.String(l.lens.Version)),
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
									htmx.Text("Description"),
								),
								htmx.H3(
									htmx.Text(l.lens.Description),
								),
							),
							lenses.LensesStatus(
								lenses.LensesStatusProps{
									IsDraft: l.lens.IsDraft,
								},
								htmx.ID("status"),
							),
							cards.Actions(
								cards.ActionsProps{},
								lenses.LensesPublishButton(
									lenses.LensesPublishButtonProps{
										ID:      l.lens.ID,
										IsDraft: l.lens.IsDraft,
									},
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(fmt.Sprintf(utils.DeleteLensUrlFormat, l.lens.ID)),
									htmx.HxConfirm("Are you sure you want to delete this lens?"),
									htmx.Text("Delete"),
								),
							),
						),
					),
					lenses.LensMetadataCard(
						lenses.LensMetadataCardProps{
							Lens: l.lens,
						},
					),
				)
			},
		),
	)
}
