package designs

import (
	"fmt"

	"github.com/zeiss/service-lens/internal/builders"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"

	"go.abhg.dev/goldmark/mermaid"
)

// DesignBodyCardProps ...
type DesignBodyCardProps struct {
	// ClassNames ...
	ClassNames htmx.ClassNames
	// Design ...
	Design models.Design
	// Markdown ...
	Markdown string
	// User ...
	User adapters.GothUser
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
		cards.Body(
			cards.BodyProps{},
			htmx.Div(
				htmx.ID("body"),
				htmx.Markdown(
					conv.Bytes(props.Design.Body),
					goldmark.WithRendererOptions(
						html.WithXHTML(),
						html.WithUnsafe(),
						renderer.WithNodeRenderers(util.Prioritized(builders.NewMarkdownBuilder(), 1)),
					),
					goldmark.WithExtensions(
						extension.GFM,
						emoji.Emoji,
						&mermaid.Extender{},
					),
				),
			),
			cards.Actions(
				cards.ActionsProps{
					ClassNames: htmx.ClassNames{
						tailwind.JustifyEnd:     false,
						tailwind.JustifyBetween: true,
					},
				},
				DesignReactions(
					DesignReactionsProps{
						User:   props.User,
						Design: props.Design,
					},
				),
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxSwap("outerHTML"),
					htmx.HxGet(fmt.Sprintf(utils.EditBodyUrlFormat, props.Design.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
