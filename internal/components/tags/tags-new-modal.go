package tags

import (
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/modals"
)

// NewTagModalProps ...
type NewTagModalProps struct{}

// NewTagModal ...
func NewTagModal() htmx.Node {
	return modals.Modal(
		modals.ModalProps{
			ID: "new_tag_modal",
		},
		htmx.FormElement(
			htmx.HxPost(utils.CreateTagUrlFormat),
			forms.FormControl(
				forms.FormControlProps{
					ClassNames: htmx.ClassNames{},
				},
				forms.TextInputBordered(
					forms.TextInputProps{
						Name:        "name",
						Placeholder: "Provide tag name ...",
					},
					htmx.AutoComplete("off"),
				),
				forms.FormControlLabel(
					forms.FormControlLabelProps{},
					forms.FormControlLabelText(
						forms.FormControlLabelTextProps{
							ClassNames: htmx.ClassNames{
								"text-neutral-500": true,
							},
						},
						htmx.Text("Use a unique name to identify the tag that has between 3 and 255 characters."),
					),
				),
			),
			forms.FormControl(
				forms.FormControlProps{
					ClassNames: htmx.ClassNames{},
				},
				forms.TextInputBordered(
					forms.TextInputProps{
						Name:        "value",
						Placeholder: "Provide a tag value ...",
					},
					htmx.AutoComplete("off"),
				),
				forms.FormControlLabel(
					forms.FormControlLabelProps{},
					forms.FormControlLabelText(
						forms.FormControlLabelTextProps{
							ClassNames: htmx.ClassNames{
								"text-neutral-500": true,
							},
						},
						htmx.Text("Use a unique value of the tag that has between 3 and 255 characters."),
					),
				),
			),
			modals.ModalAction(
				modals.ModalActionProps{},
				buttons.Button(
					buttons.ButtonProps{
						Type: "submit",
					},
					htmx.Text("Create"),
				),
			),
		),
	)
}
