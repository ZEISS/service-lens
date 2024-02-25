package utils

import (
	"github.com/gofiber/fiber/v2"
)

// TeamFromContext returns the team from the context.
func TeamFromContext(c *fiber.Ctx) (string, error) {
	return c.Params("team", ""), nil
}
