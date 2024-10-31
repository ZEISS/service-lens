package handlers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

type AssetsHandler struct{}

func NewAssetsHandler() *PreviewHandler {
	return &PreviewHandler{}
}

func (h *PreviewHandler) Upload(c *fiber.Ctx) (htmx.Node, error) {
	return nil, nil
}
