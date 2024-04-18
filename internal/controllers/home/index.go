package home

import (
	"github.com/gofiber/fiber/v2"
	authz "github.com/zeiss/fiber-authz"
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
		components.PageProps{},
		components.Layout(
			components.LayoutProps{
				User: h.Values(utils.ValuesKeyUser).(*authz.User),
				Team: h.Values(utils.ValuesKeyTeam).(*authz.Team),
			},
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
