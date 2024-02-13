package handlers

import (
	"github.com/zeiss/service-lens/internal/components"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

type indexHandler struct{}

// NewIndexHandler returns a new IndexHandler.
func NewIndexHandler() *indexHandler {
	return &indexHandler{}
}

// Index is the handler for the index page.
func (h *indexHandler) Index() fiber.Handler {
	return htmx.NewCompHandler(components.Page(components.PageProps{}))
}
