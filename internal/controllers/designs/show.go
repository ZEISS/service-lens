package designs

import (
	"context"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// ShowDesignControllerImpl ...
type ShowDesignControllerImpl struct {
	Design models.Design
	Body   string
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewShowDesignController ...
func NewShowDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowDesignControllerImpl {
	return &ShowDesignControllerImpl{
		store: store,
	}
}

// Prepare ...
func (l *ShowDesignControllerImpl) Prepare() error {
	err := l.BindParams(&l.Design)
	if err != nil {
		return err
	}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &l.Design)
	})
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (l *ShowDesignControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       l.Design.Title,
				Path:        l.Ctx().Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
				Head: []htmx.Node{
					htmx.Link(
						htmx.Attribute("href", "https://cdn.jsdelivr.net/simplemde/1.11/simplemde.min.css"),
						htmx.Rel("stylesheet"),
						htmx.Type("text/css"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://cdn.jsdelivr.net/simplemde/1.11/simplemde.min.js"),
						htmx.Type("text/javascript"),
					),
				},
			},
			func() htmx.Node {
				return htmx.Fragment(
					designs.DesignTitleCard(
						designs.DesignTitleCardProps{
							Design: l.Design,
						},
					),
					designs.DesignBodyCard(
						designs.DesignBodyCardProps{
							User:   l.Session().User,
							Design: l.Design,
						},
					),
					designs.DesignMetadataCard(
						designs.DesignMetadataCardProps{
							Design: l.Design,
						},
					),
					designs.DesignRevisionsCard(
						designs.DesignRevisionsCardProps{
							DesignID: l.Design.ID,
						},
					),
					designs.DesignTagsCard(
						designs.DesignTagsCardProps{
							Design: l.Design,
						},
					),
					designs.DesignCommentsCard(
						designs.DesignCommentsCardProps{
							User:   l.Session().User,
							Design: l.Design,
						},
					),
				)
			},
		),
	)
}
