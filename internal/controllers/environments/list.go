package environments

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/environments"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// EnvironmentListControllerImpl ...
type EnvironmentListControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.UnimplementedController
}

// NewEnvironmentListController ...
func NewEnvironmentListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EnvironmentListControllerImpl {
	return &EnvironmentListControllerImpl{
		store: store,
	}
}

// Get ...
func (c *EnvironmentListControllerImpl) Get() error {
	return c.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       "Environments",
				Path:        c.Path(),
				User:        c.Session().User,
				Development: c.IsDevelopment(),
			},
			func() htmx.Node {
				results := tables.Results[models.Environment]{SearchFields: []string{"Name"}}

				errorx.Panic(c.BindQuery(&results))
				errorx.Panic(c.store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListEnvironments(ctx, &results)
				}))

				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						environments.EnvironmentsTable(
							environments.EnvironmentsTableProps{
								Environments: results.GetRows(),
								Offset:       results.GetOffset(),
								Limit:        results.GetLimit(),
								Total:        results.GetLen(),
								URL:          c.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
