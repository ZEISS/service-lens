package designs

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/cards"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

var _ = htmx.Controller(&ListDesignsControllerImpl{})

// ListDesignsControllerImpl ...
type ListDesignsControllerImpl struct {
	results tables.Results[models.Design]
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewListDesignsController ...
func NewListDesignsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ListDesignsControllerImpl {
	return &ListDesignsControllerImpl{store: store}
}

// Prepare ...
func (l *ListDesignsControllerImpl) Prepare() error {
	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListDesigns(ctx, &l.results)
	})
}

// Prepare ...
func (l *ListDesignsControllerImpl) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{
				Title: "Designs",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						designs.DesignsTable(
							designs.DesignsTableProps{
								Designs: l.results.GetRows(),
								Offset:  l.results.GetOffset(),
								Limit:   l.results.GetLimit(),
								Total:   l.results.GetLen(),
							},
						),
					),
				),
			),
		),
	)
}
