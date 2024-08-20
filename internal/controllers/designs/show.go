package designs

import (
	"bytes"
	"context"

	"github.com/zeiss/service-lens/internal/builder"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
	"go.abhg.dev/goldmark/mermaid"
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
			renderer.WithNodeRenderers(util.Prioritized(builder.NewMarkdownBuilder(), 1)),
		),
		goldmark.WithExtensions(
			extension.GFM,
			emoji.Emoji,
			&mermaid.Extender{},
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
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: l.Design.Title,
				Path:  l.Ctx().Path(),
				User:  l.Session().User,
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
							Design:   l.Design,
							Markdown: l.Body,
						},
					),
					designs.DesignMetadataCard(
						designs.DesignMetadataCardProps{
							Design: l.Design,
						},
					),
					designs.DesignTagsCard(
						designs.DesignTagsCardProps{
							Design: l.Design,
						},
					),
					designs.DesignCommentsCard(
						designs.DesignCommentsCardProps{
							Design: l.Design,
						},
					),
				)
			},
		),
	)
}
