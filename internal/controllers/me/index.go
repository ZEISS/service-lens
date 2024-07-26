package me

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
)

// MeController ...
type MeController struct {
	user  adapters.GothUser
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewMeIndexController ...
func NewMeController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *MeController {
	return &MeController{
		store: store,
	}
}

// Prepare ...
func (m *MeController) Prepare() error {
	return m.store.ReadTx(m.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetUser(ctx, &m.user)
	})
}

// Get ...
func (m *MeController) Get() error {
	return m.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: m.Path(),
			},
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
									Value:    m.user.Name,
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
									Value:    m.user.Email,
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
			),
		),
	)
}
