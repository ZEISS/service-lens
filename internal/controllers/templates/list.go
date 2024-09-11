package templates

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/templates"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
	seed "github.com/zeiss/gorm-seed"
)

var _ = htmx.Controller(&ListTemplatesControllerImpl{})

// ListTemplatesControllerImpl ...
type ListTemplatesControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewListTemplatesController ...
func NewListTemplatesController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ListTemplatesControllerImpl {
	return &ListTemplatesControllerImpl{store: store}
}

// Prepare ...
func (l *ListTemplatesControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
			},
			func() htmx.Node {
				results := tables.Results[models.Template]{SearchFields: []string{"Name"}}

				errorx.Panic(l.BindQuery(&results))
				errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListTemplates(ctx, &results)
				}))

				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						templates.TemplatesTable(
							templates.TemplatesTableProps{
								Templates: results.GetRows(),
								Offset:    results.GetOffset(),
								Limit:     results.GetLimit(),
								Total:     results.GetLen(),
								URL:       l.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
