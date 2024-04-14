package me

import (
	goth "github.com/zeiss/fiber-goth"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// MeIndexController ...
type MeIndexController struct {
	db ports.Repository

	htmx.DefaultController
}

// NewMeIndexController ...
func NewMeIndexController(db ports.Repository) *MeIndexController {
	return &MeIndexController{
		db: db,
	}
}

// Prepare ...
func (m *MeIndexController) Prepare() error {
	if err := m.BindValues(utils.User(m.db), utils.Team(m.db)); err != nil {
		return err
	}

	return nil
}

// Get ...
func (m *MeIndexController) Get() error {
	session, err := goth.SessionFromContext(m.Hx().Context())
	if err != nil {
		return err
	}

	user, err := m.db.GetUserByID(m.Hx().Context().Context(), session.UserID)
	if err != nil {
		return err
	}

	return m.Hx().RenderComp(
		components.Page(
			m.DefaultCtx(),
			components.PageProps{},
			components.Layout(
				m.DefaultCtx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					cards.CardBordered(
						cards.CardProps{},
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
											Value:    user.Name,
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
											Value:    user.Email,
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
			),
		),
	)
}
