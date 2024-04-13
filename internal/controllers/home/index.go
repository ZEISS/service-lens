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
	db  ports.Repository
	ctx htmx.Ctx

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
	ctx, err := htmx.NewDefaultContext(h.Hx().Ctx(), utils.Team(h.Hx().Ctx(), h.db), utils.User(h.Hx().Ctx(), h.db))
	if err != nil {
		return err
	}
	h.ctx = ctx

	return nil
}

// Get ...
func (h *HomeIndexController) Get(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		h.ctx,
		components.PageProps{},
		components.Layout(
			h.ctx,
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
