package designs

import (
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// NewDesignControllerImpl ...
type NewDesignControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewDesignController ...
func NewDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewDesignControllerImpl {
	return &NewDesignControllerImpl{store: store}
}

// Get ...
func (l *NewDesignControllerImpl) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: l.Path(),
				},
				htmx.FormElement(
					htmx.HxPost("/designs/new"),
					htmx.HxTarget("this"),
					htmx.HxSwap("outerHTML"),
					htmx.ID("body"),
					cards.CardBordered(
						cards.CardProps{},
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
											"h-64": true,
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
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Tags - Optional"),
							),
						),
					),
				),
			),
		),
	)
}
