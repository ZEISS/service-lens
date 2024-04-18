package utils

import (
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
func Team(db ports.Repository) htmx.BindFunc {
	return func(ctx *fiber.Ctx) (interface{}, interface{}, error) {
		slug, err := TeamFromContext(ctx)
		if err != nil {
			return nil, nil, err
		}

		team, err := db.GetTeamBySlug(ctx.Context(), slug)
		if err != nil {
			return nil, nil, err
		}

		return ValuesKeyTeam, team, nil
	}
}
