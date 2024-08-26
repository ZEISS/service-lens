package designs

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/cards"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
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
	return &ListDesignsControllerImpl{
		results: tables.Results[models.Design]{SearchFields: []string{"Title"}},
		store:   store,
	}
}

// Prepare ...
func (l *ListDesignsControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
			},
			func() htmx.Node {
				errorx.Panic(l.BindQuery(&l.results))
				errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListDesigns(ctx, &l.results)
				}))

				templates := tables.Results[models.Template]{}
				errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListTemplates(ctx, &templates)
				}))

				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						designs.DesignsTable(
							designs.DesignsTableProps{
								Designs:   l.results.GetRows(),
								Templates: templates.GetRows(),
								Offset:    l.results.GetOffset(),
								Limit:     l.results.GetLimit(),
								Total:     l.results.GetLen(),
								Search:    l.results.GetSearch(),
								URL:       l.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
