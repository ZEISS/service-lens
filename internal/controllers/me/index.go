package me

import (
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
)

// MeController ...
type MeController struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewMeIndexController ...
func NewMeController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *MeController {
	return &MeController{
		store: store,
	}
}

// Get ...
func (m *MeController) Get() error {
	return m.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        m.Path(),
				User:        m.Session().User,
				Development: m.IsDevelopment(),
			},
			func() htmx.Node {
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
							htmx.Text("Profile"),
						),
						htmx.Form(
							htmx.HxPost("/me"),
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
										Name:     "username",
										Value:    m.Session().User.Name,
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
										htmx.Text("Your full nane as it will appear in the system."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Email"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:     "email",
										Value:    m.Session().User.Email,
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
										htmx.Text("Your email address. This is where we will send notifications."),
									),
								),
							),

							cards.Actions(
								cards.ActionsProps{},
								buttons.OutlinePrimary(
									buttons.ButtonProps{
										Disabled: true,
									},
									htmx.Attribute("type", "submit"),
									htmx.Text("Update Profile"),
								),
							),
						),
					),
				)
			},
		),
	)
}
