package workloads

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/modals"
	"github.com/zeiss/service-lens/internal/utils"
)

// AddTagModalProps ...
type AddTagModalProps struct {
	// WorkloadID ...
	WorkloadID uuid.UUID
}

// AddTagModal ...
func AddTagModal(props AddTagModalProps) htmx.Node {
	return modals.Modal(
		modals.ModalProps{
			ID: "add_tag_modal",
		},
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.AddWorkloadTagUrlFormat, props.WorkloadID)),
			htmx.HxTrigger("submit"),
			htmx.HxOn("htmx:after-settle", "event.target.closest('dialog').close(), event.target.reset()"),
			htmx.HxSwap("none"),
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
				buttons.Ghost(
					buttons.ButtonProps{
						Type: "button",
					},
					htmx.Text("Cancel"),
					htmx.Attribute("formnovalidate", ""),
					htmx.OnClick("event.target.closest('dialog').close()"),
				),
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
