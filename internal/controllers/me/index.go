package me

import (
	goth "github.com/zeiss/fiber-goth"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// MeIndexController ...
type MeIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewMeIndexController ...
func NewMeIndexController(db ports.Repository) *MeIndexController {
	return &MeIndexController{db, htmx.UnimplementedController{}}
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
			m.Hx(),
			components.PageProps{},
			components.Layout(
				m.Hx(),
				components.LayoutProps{},
				htmx.Form(
					htmx.HxPost("/me"),
					htmx.Label(
						htmx.ClassNames{
							"form-control": true,
							"w-full":       true,
							"max-w-xs":     true,
						},
						forms.TextInput(
							forms.TextInputProps{
								Name:     "username",
								Value:    user.Name,
								Disabled: true,
							},
						),
					),
					htmx.Label(
						htmx.ClassNames{
							"form-control": true,
							"w-full":       true,
							"max-w-xs":     true,
						},
						forms.TextInput(
							forms.TextInputProps{
								Name:     "email",
								Value:    user.Email,
								Disabled: true,
							},
						),
					),
					buttons.OutlineAccent(
						buttons.ButtonProps{},
						htmx.Attribute("type", "submit"),
						htmx.Text("Update Profile"),
					),
				),
			),
		),
	)
}
