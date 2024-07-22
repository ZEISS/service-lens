package designs

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignTitleCardProps ...
type DesignTitleCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
	Markdown   string
}

// DesignTitleCard ...
func DesignTitleCard(props DesignTitleCardProps) htmx.Node {
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
		htmx.ID("title"),
		cards.Body(
			cards.BodyProps{},
			htmx.H1(htmx.Text(props.Design.Title)),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Outline(
					buttons.ButtonProps{},
					htmx.HxGet(fmt.Sprintf(utils.EditTitleUrlFormat, props.Design.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
