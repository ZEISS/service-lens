package lenses

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/loading"
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
		htmx.H2(
			htmx.ClassNames{
				"text-xl":         true,
				tailwind.Mb2:      true,
				tailwind.FontBold: true,
			},
			htmx.Text("New Lens"),
		),
		htmx.FormElement(
			htmx.ID("new-lens-form"),
			htmx.HxEncoding("multipart/form-data"),
			htmx.HxPost(utils.CreateLensUrlFormat),
			htmx.HxIndicator(".htmx-indicator"),
			htmx.HxDisabledElt("find button, find input"),
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
						forms.FileInputProps{},
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
				buttons.Button(
					buttons.ButtonProps{
						Type: "submit",
					},
					loading.Spinner(
						loading.SpinnerProps{
							ClassNames: htmx.ClassNames{
								"htmx-indicator": true,
							},
						},
					),
					htmx.Text("Create Lens"),
				),
			),
		),
	)
}
