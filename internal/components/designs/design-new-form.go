package designs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
)

// DesignNewFormProps ...
type DesignNewFormProps struct {
	ClassNames htmx.ClassNames
}

// DesignNewForm ...
func DesignNewForm(props DesignNewFormProps) htmx.Node {
	return htmx.FormElement(
		htmx.HxPost(""),
		htmx.HxTarget("this"),
		htmx.HxSwap("outerHTML"),
		htmx.ID("body"),
		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
			},
			cards.Body(
				cards.BodyProps{},
				cards.Title(
					cards.TitleProps{},
					htmx.Text("Create Design"),
				),
				forms.FormControl(
					forms.FormControlProps{
						ClassNames: htmx.ClassNames{},
					},
					forms.TextInputBordered(
						forms.TextInputProps{
							Name:        "title",
							Placeholder: "Title",
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
							htmx.Text("The title must be from 3 to 2048 characters."),
						),
					),
				),
				forms.FormControl(
					forms.FormControlProps{
						ClassNames: htmx.ClassNames{},
					},
					forms.TextareaBordered(
						forms.TextareaProps{
							ClassNames: htmx.ClassNames{
								"h-[50vh]": true,
							},
							Name:        "body",
							Placeholder: "Start typing...",
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
				cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"w-full": true,
							"my-4":   true,
						},
					},

					htmx.Input(
						htmx.Attribute("type", "hidden"),
						htmx.ID("environment"),
						htmx.Attribute("name", "environment_id"),
						htmx.Value(""),
					),
					dropdowns.Dropdown(
						dropdowns.DropdownProps{},
						htmx.HyperScript("on click from (closest <a/>) set (previous <input/>).value to 'test'"),
						dropdowns.DropdownButton(
							dropdowns.DropdownButtonProps{},
							htmx.Text("Select Environment"),
							htmx.HxGet("/workloads/partials/environments"),
							htmx.HxTarget("#environments-list"),
							htmx.HxSwap("innerHTML"),
							htmx.ID("environments-button"),
						),
						dropdowns.DropdownMenuItems(
							dropdowns.DropdownMenuItemsProps{
								TabIndex: 1,
							},
							htmx.ID("environments-list"),
							dropdowns.DropdownMenuItem(
								dropdowns.DropdownMenuItemProps{},
							),
						),
					),
				),
				cards.Actions(
					cards.ActionsProps{},
					buttons.Outline(
						buttons.ButtonProps{},
						htmx.Attribute("type", "submit"),
						htmx.Text("Save Design"),
					),
				),
			),
		),
		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
			},
			cards.Body(
				cards.BodyProps{},
				cards.Title(
					cards.TitleProps{},
					htmx.Text("Tags - Optional"),
				),
			),
		),
	)
}
