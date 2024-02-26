package controllers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// Home ...
type Home struct {
	db ports.Repository
}

// NewHomeController ...
func NewHomeController(db ports.Repository) *Home {
	return &Home{db}
}

// Index ...
func (h *Home) Index(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		components.PageProps{},
		components.Layout(
			components.LayoutProps{}.WithContext(c),
			components.Wrap(
				components.WrapProps{},
				htmx.Div(
					htmx.H1(
						htmx.Text("Welcome to Service Lens"),
					),
				),
			),
		),
	), nil
}
