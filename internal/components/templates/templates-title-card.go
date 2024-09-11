package templates

import (
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/fiber-htmx/components/typography"
)

// TemplateTitleCardProps ...
type TemplateTitleCardProps struct {
	ClassNames htmx.ClassNames
	Template   models.Template
	Markdown   string
}

// TemplateTitleCard ...
func TemplateTitleCard(props TemplateTitleCardProps) htmx.Node {
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
		htmx.ID("name"),
		cards.Body(
			cards.BodyProps{},
			typography.H2(
				typography.Props{},
				htmx.Text(props.Template.Name),
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxGet(fmt.Sprintf(utils.EditTemplateTitleUrlFormat, props.Template.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
