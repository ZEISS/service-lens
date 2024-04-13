package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/ports"
)

const (
	// ValuesKeyTeam ...
	ValuesKeyTeam = "team"
)

// TeamFromContext returns the team from the context.
func TeamFromContext(c *fiber.Ctx) (string, error) {
	return c.Params("team", ""), nil
}

// Team ...
func Team(c *fiber.Ctx, db ports.Repository) htmx.ContextFunc {
	return func(ctx context.Context) (interface{}, interface{}, error) {
		slug, err := TeamFromContext(c)
		if err != nil {
			return nil, nil, err
		}

		team, err := db.GetTeamBySlug(c.Context(), slug)
		if err != nil {
			return nil, nil, err
		}

		return ValuesKeyTeam, team, nil
	}
}
