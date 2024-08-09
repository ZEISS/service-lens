package lenses

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/modals"
	"github.com/zeiss/fiber-htmx/components/progress"
	"github.com/zeiss/service-lens/internal/utils"
)

// NewLensModalProps ...
type NewLensModalProps struct{}

// NewLensModal ...
func NewLensModal() htmx.Node {
	return modals.Modal(
		modals.ModalProps{
			ID: "new_lens_modal",
		},
		htmx.FormElement(
			htmx.ID("new-lens-form"),
			htmx.HxEncoding("multipart/form-data"),
			htmx.HxPost(utils.CreateLensUrlFormat),
			htmx.Attribute("_", "on htmx:xhr:progress(loaded, total) set #new-lens-progress.value to (loaded/total)*100'"),
			htmx.Div(
				forms.FileInputBordered(
					forms.FileInputProps{},
					htmx.Attribute("name", "spec"),
				),
			),
			progress.Progress(
				progress.ProgressProps{
					ClassNames: htmx.ClassNames{
						"block": true,
						"my-4":  true,
					},
				},
				htmx.ID("new-lens-progress"),
				htmx.Value("0"),
				htmx.Max("100"),
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
