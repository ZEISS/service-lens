package controllers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// Login ...
type Login struct {
	db ports.Repository
}

// NewLoginController ...
func NewLoginController(db ports.Repository) *Login {
	return &Login{db}
}

// Show ...
func (l *Login) Show(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		components.PageProps{}.WithContext(c),
		components.Wrap(
			components.WrapProps{},
			htmx.A(
				htmx.Attribute("href", "/login/github"),
				htmx.Text("Login with GitHub"),
			),
		),
	), nil
}
