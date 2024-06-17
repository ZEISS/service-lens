package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeiss/service-lens/internal/models"
	"gorm.io/gorm"
)

type teamContextKey int

const teamKey teamContextKey = iota

// ResolveTeam ...
func ResolveTeam(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		team := models.Team{
			Slug: c.Params("t_slug"),
		}

		err := db.WithContext(c.Context()).Where(&team).First(&team).Error
		if err != nil {
			return err
		}

		c.Locals(teamKey, team)

		return c.Next()
	}
}

// FromContextTeam ...
func FromContextTeam(c *fiber.Ctx) models.Team {
	return c.Locals(teamKey).(models.Team)
}
