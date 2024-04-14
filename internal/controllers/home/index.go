package home

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// HomeIndexController ...
type HomeIndexController struct {
	db ports.Repository

	htmx.DefaultController
}

// NewHomeIndexController ...
func NewHomeIndexController(db ports.Repository) *HomeIndexController {
	return &HomeIndexController{
		db: db,
	}
}

// Prepare ...
func (h *HomeIndexController) Prepare() error {
	if err := h.BindValues(utils.User(h.db), utils.Team(h.db)); err != nil {
		return err
	}

	return nil
}

// Get ...
func (h *HomeIndexController) Get(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		h.DefaultCtx(),
		components.PageProps{},
		components.Layout(
			h.DefaultCtx(),
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
