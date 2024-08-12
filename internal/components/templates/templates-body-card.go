package templates

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/service-lens/internal/models"
)

// TemplateBodyCardProps ...
type TemplateBodyCardProps struct {
	ClassNames htmx.ClassNames
	Template   models.Template
	Markdown   string
}

// TemplateBodyCard ...
func TemplateBodyCard(props TemplateBodyCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
			),
		},
		htmx.HxTarget("this"),
		htmx.HxSwap("outerHTML"),
		htmx.ID("body"),
		cards.Body(
			cards.BodyProps{},
			htmx.Div(
				htmx.Raw(props.Markdown),
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Outline(
					buttons.ButtonProps{},
					htmx.HxGet(""),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
