package environments

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
)

// NewFormProps ...
type NewFormProps struct{}

// NewForm ...
func NewForm(props NewFormProps) htmx.Node {
	return htmx.FormElement(
		htmx.HxPost(""),
		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					"m-2": true,
				},
			},
			cards.Body(
				cards.BodyProps{},
				cards.Title(
					cards.TitleProps{},
					htmx.Text("Properties"),
				),
				forms.FormControl(
					forms.FormControlProps{},
					forms.FormControlLabel(
						forms.FormControlLabelProps{},
						forms.FormControlLabelText(
							forms.FormControlLabelTextProps{},
							htmx.Text("Name"),
						),
					),

					forms.TextInputBordered(
						forms.TextInputProps{
							Name: "name",
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
							htmx.Text("The name must be from 3 to 100 characters. At least 3 characters must be non-whitespace."),
						),
					),
					forms.FormControl(
						forms.FormControlProps{},
						forms.FormControlLabel(
							forms.FormControlLabelProps{},
							forms.FormControlLabelText(
								forms.FormControlLabelTextProps{},
								htmx.Text("Description"),
							),
						),
						forms.TextareaBordered(
							forms.TextareaProps{
								Name: "description",
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
								htmx.Text("The description must be from 3 to 1024 characters."),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Attribute("type", "submit"),
							htmx.Text("Create Environment"),
						),
					),
				),
			),
		),
		components.AddTags(
			components.AddTagsProps{},
		),
	)
}
