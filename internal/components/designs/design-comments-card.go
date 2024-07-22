package designs

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignCommentsCardProps ...
type DesignCommentsCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
}

// DesignCommentsCard ...
func DesignCommentsCard(props DesignCommentsCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
			),
		},
		cards.Body(
			cards.BodyProps{},
			htmx.Div(
				htmx.ID("comments"),
				htmx.Group(htmx.ForEach(tables.RowsPtr(props.Design.Comments), func(c *models.DesignComment, choiceIdx int) htmx.Node {
					return cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"my-4": true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							htmx.Text(c.Comment),
							cards.Actions(
								cards.ActionsProps{},
								avatars.AvatarRoundSmall(
									avatars.AvatarProps{},
									htmx.Img(
										htmx.Attribute("src", utils.PtrStr(c.Author.Image)),
									),
								),
							),
						),
					)
				})...),
			),
			htmx.FormElement(
				htmx.HxPost(fmt.Sprintf(utils.CreateDesignCommentUrlFormat, props.Design.ID)),
				htmx.HxTarget("#comments"),
				htmx.HxSwap("beforeend"),
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						forms.FormControl(
							forms.FormControlProps{
								ClassNames: htmx.ClassNames{},
							},
							forms.TextareaBordered(
								forms.TextareaProps{
									ClassNames: htmx.ClassNames{
										"h-32": true,
									},
									Name:        "comment",
									Placeholder: "Add a comment...",
								},
							),
							forms.FormControlLabel(
								forms.FormControlLabelProps{},
								forms.FormControlLabelText(
									forms.FormControlLabelTextProps{
										ClassNames: htmx.ClassNames{
											"text-neutral-500": true,
										},
									},
									htmx.Text("Supports Markdown."),
								),
							),
						),
						cards.Actions(
							cards.ActionsProps{},
							buttons.Outline(
								buttons.ButtonProps{},
								htmx.Attribute("type", "submit"),
								htmx.Text("Comment"),
							),
						),
					),
				),
			),
		),
	)
}
