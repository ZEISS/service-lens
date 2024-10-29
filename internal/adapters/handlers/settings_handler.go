package handlers

import (
	"github.com/gofiber/fiber/v2"
	goth "github.com/zeiss/fiber-goth"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	reload "github.com/zeiss/fiber-reload"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/components"
)

type SettingsHandler struct{}

func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{}
}

func (h *SettingsHandler) ListSettings(c *fiber.Ctx) (htmx.Node, error) {
	return components.DefaultLayout(
		components.DefaultLayoutProps{
			Path:        c.Path(),
			User:        errorx.Ignore(goth.SessionFromContext(c)).User,
			Development: reload.IsDevelopment(c.UserContext()),
		},
		func() htmx.Node {
			return htmx.Fragment(
				cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.Merge(
							htmx.ClassNames{
								tailwind.M2: true,
							},
						),
					},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Settings"),
						),
					),
				),
			)
		},
	), nil
}
