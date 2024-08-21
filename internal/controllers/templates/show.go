package templates

import (
	"bytes"
	"context"

	"github.com/zeiss/service-lens/internal/builder"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/templates"
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

// ShowTemplateControllerImpl ...
type ShowTemplateControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewShowTemplateController ...
func NewShowTemplateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowTemplateControllerImpl {
	return &ShowTemplateControllerImpl{
		store: store,
	}
}

// Get ...
func (l *ShowTemplateControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{},
			func() htmx.Node {
				template := models.Template{}

				err := l.BindParams(&template)
				if err != nil {
					panic(err)
				}

				err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTemplate(ctx, &template)
				})
				if err != nil {
					panic(err)
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
				err = markdown.Convert([]byte(template.Body), &b)
				if err != nil {
					panic(err)
				}

				template.Body = b.String()

				return htmx.Fragment(
					templates.TemplateBodyCard(
						templates.TemplateBodyCardProps{
							Template: template,
							Markdown: template.Body,
						},
					),
				)
			},
		),
	)
}
