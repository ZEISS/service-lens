package tags

import (
	"context"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/tags"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
	seed "github.com/zeiss/gorm-seed"
)

// TagsListControllerImpl ...
type TagsListControllerImpl struct {
	tags  tables.Results[models.Tag]
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTagsListController ...
func NewTagsListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TagsListControllerImpl {
	return &TagsListControllerImpl{store: store}
}

// Prepare ...
func (w *TagsListControllerImpl) Prepare() error {
	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTags(ctx, &w.tags)
	})
}

// Get ...
func (w *TagsListControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        w.Path(),
				User:        w.Session().User,
				Development: w.IsDevelopment(),
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
						tags.TagsTable(
							tags.TagsTableProps{
								Tags:   w.tags.GetRows(),
								Offset: w.tags.GetOffset(),
								Limit:  w.tags.GetLimit(),
								Total:  w.tags.GetTotalRows(),
							},
						),
					),
				)
			},
		),
	)
}
