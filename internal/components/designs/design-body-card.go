package designs

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
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
					htmx.HxGet(fmt.Sprintf(utils.EditBodyUrlFormat, props.Design.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
