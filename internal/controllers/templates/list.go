package templates

import (
	"github.com/zeiss/fiber-htmx/components/cards"
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
	results tables.Results[models.Template]
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
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
				Path: l.Path(),
				User: l.Session().User,
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
						templates.TemplatesTable(
							templates.TemplatesTableProps{
								Templates: l.results.GetRows(),
								Offset:    l.results.GetOffset(),
								Limit:     l.results.GetLimit(),
								Total:     l.results.GetLen(),
							},
						),
					),
				)
			},
		),
	)
}
