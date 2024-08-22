package designs

import (
	"fmt"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/service-lens/internal/builder"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
	"go.abhg.dev/goldmark/mermaid"
)

// DesignBodyCardProps ...
type DesignBodyCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
	Markdown   string
}

// DesignBodyCard ...
func DesignBodyCard(props DesignBodyCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
				},
			),
		},
		htmx.HxTarget("this"),
		htmx.HxSwap("outerHTML"),
		htmx.ID("body"),
		cards.Body(
			cards.BodyProps{},
			htmx.Div(
				htmx.Markdown(
					conv.Bytes(props.Design.Body),
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
				),
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxGet(fmt.Sprintf(utils.EditBodyUrlFormat, props.Design.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
