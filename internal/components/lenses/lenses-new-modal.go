package lenses

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/modals"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/utils"
)

// NewLensModalProps ...
type NewLensModalProps struct{}

// NewLensModal ...
func NewLensModal(props NewLensModalProps) htmx.Node {
	return modals.Modal(
		modals.ModalProps{
			ID: "new_lens_modal",
		},
		htmx.FormElement(
			htmx.ID("new-lens-form"),
			htmx.HxEncoding("multipart/form-data"),
			htmx.HxTrigger("submit"),
			htmx.HxPost(utils.CreateLensUrlFormat),
			htmx.HxDisabledElt("find button, find input"),
			htmx.HxOn("htmx:after-settle", "event.target.closest('dialog').close(), event.target.reset()"),
			htmx.HxSwap("none"),
			htmx.Div(
				forms.FormControl(
					forms.FormControlProps{},
					forms.FormControlLabel(
						forms.FormControlLabelProps{},
						forms.FormControlLabelText(
							forms.FormControlLabelTextProps{
								ClassNames: htmx.ClassNames{
									"text-neutral-500": true,
								},
							},
							htmx.Text("Select the file to upload."),
						),
					),
					forms.FileInputBordered(
						forms.FileInputProps{
							ClassNames: htmx.ClassNames{
								tailwind.MaxWXs: false,
							},
						},
						htmx.Attribute("name", "spec"),
					),
					forms.FormControlLabel(
						forms.FormControlLabelProps{},
						forms.FormControlLabelText(
							forms.FormControlLabelTextProps{
								ClassNames: htmx.ClassNames{
									"text-neutral-500": true,
								},
							},
							htmx.Text("Needs to conform the lens format specification."),
						),
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
					htmx.Text("Add Lens"),
				),
			),
		),
	)
}
