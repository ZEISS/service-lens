package handlers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/login"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Login() htmx.CompFunc {
	return func(c *fiber.Ctx) (htmx.Node, error) {
		return components.Page(
			components.PageProps{},
			components.Wrap(
				components.WrapProps{},
				login.NewLogin(),
			),
		), nil
	}
}
