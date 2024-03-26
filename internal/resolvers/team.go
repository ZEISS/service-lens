package resolvers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

const (
	// ValuesKeyTeam ...
	ValuesKeyTeam = "team"
)

// Team ...
func Team(db ports.Repository) htmx.ResolveFunc {
	return func(c *fiber.Ctx) (interface{}, interface{}, error) {
		slug, err := utils.TeamFromContext(c)
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
