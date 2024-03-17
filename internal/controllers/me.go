package controllers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// Home ...
type Me struct {
	db ports.Repository
}

// NewMeController ...
func NewMeController(db ports.Repository) *Home {
	return &Home{db}
}

// Index ...
func (m *Me) Index(c *fiber.Ctx) (htmx.Node, error) {
	ctx := htmx.DefaultCtx()
	ctx.Context(c)

	return components.Page(
		components.PageProps{
			Ctx: ctx,
		},
		components.Layout(
			components.LayoutProps{
				Ctx: ctx,
			},
			components.Wrap(
				components.WrapProps{},
				htmx.Div(
					htmx.H1(
						htmx.Text("Me Page"),
					),
				),
			),
		),
	), nil
}
