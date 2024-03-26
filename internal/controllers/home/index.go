package home

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// HomeIndexController ...
type HomeIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewHomeIndexController ...
func NewHomeIndexController(db ports.Repository) *HomeIndexController {
	return &HomeIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (h *HomeIndexController) Get(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		h.Hx,
		components.PageProps{},
		components.Layout(
			h.Hx,
			components.LayoutProps{},
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
