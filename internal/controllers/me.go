package controllers

import (
	"github.com/gofiber/fiber/v2"
	goth "github.com/zeiss/fiber-goth"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// Home ...
type Me struct {
	db ports.Repository
}

// NewMeController ...
func NewMeController(db ports.Repository) *Me {
	return &Me{db}
}

// Index ...
func (m *Me) Index(c *fiber.Ctx) (htmx.Node, error) {
	ctx := htmx.FromContext(c)

	session, err := goth.SessionFromContext(c)
	if err != nil {
		return nil, err
	}

	user, err := m.db.GetUserByID(c.Context(), session.UserID)
	if err != nil {
		return nil, err
	}

	ctx.Locals("user", user)

	return components.Page(
		ctx,
		components.PageProps{},
		components.Layout(
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
	), nil
}
