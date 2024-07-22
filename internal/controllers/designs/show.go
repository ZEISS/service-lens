package designs

import (
	"bytes"
	"context"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
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

	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			extension.GFM,
		),
	)

	var b bytes.Buffer
	err = markdown.Convert([]byte(l.Design.Body), &b)
	if err != nil {
		return err
	}

	l.Body = b.String()

	return nil
}

// Get ...
func (l *ShowDesignControllerImpl) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				designs.DesignTitleCard(
					designs.DesignTitleCardProps{
						Design: l.Design,
					},
				),
				designs.DesignBodyCard(
					designs.DesignBodyCardProps{
						Design:   l.Design,
						Markdown: l.Body,
					},
				),
				designs.DesignMetadataCard(
					designs.DesignMetadataCardProps{
						Design: l.Design,
					},
				),
				designs.DesignCommentsCard(
					designs.DesignCommentsCardProps{
						Design: l.Design,
					},
				),
			),
		),
	)
}
