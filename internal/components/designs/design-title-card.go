package designs

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
					tailwind.M2: true,
				},
			),
		},
		htmx.HxTarget("this"),
		htmx.HxSwap("outerHTML"),
		htmx.ID("title"),
		cards.Body(
			cards.BodyProps{},
			typography.H2(
				typography.Props{},
				htmx.Text(props.Design.Title),
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxGet(fmt.Sprintf(utils.EditTitleUrlFormat, props.Design.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
