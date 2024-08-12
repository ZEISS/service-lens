package designs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/models"
)

// DesignTagsCardProps ...
type DesignTagsCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
}

// DesignTagsCard ...
func DesignTagsCard(props DesignTagsCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.ClassNames{
				"m-2": true,
			},
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text("Tags - Optional"),
			),
			htmx.Group(
				htmx.ForEach(props.Design.Tags, func(tag models.Tag, idx int) htmx.Node {
					return htmx.Div(
						htmx.ClassNames{
							tailwind.Flex:    true,
							tailwind.WFull:   true,
							tailwind.SpaceX4: true,
						},
						forms.FormControl(
							forms.FormControlProps{
								ClassNames: htmx.ClassNames{},
							},
							forms.TextInputBordered(
								forms.TextInputProps{
									Value:    tag.Name,
									Disabled: true,
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
									htmx.Text("Key in a tag."),
								),
							),
						),
						forms.FormControl(
							forms.FormControlProps{
								ClassNames: htmx.ClassNames{},
							},
							forms.TextInputBordered(
								forms.TextInputProps{
									Value:    tag.Value,
									Disabled: true,
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
									htmx.Text("Value in a tag."),
								),
							),
						),
					)
				},
				)...,
			),
		),
	)
}
