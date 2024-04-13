package utils

import (
	"context"

	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	goth "github.com/zeiss/fiber-goth"
	htmx "github.com/zeiss/fiber-htmx"
)

const (
	// ValuesKeyUser ...
	ValuesKeyUser = "user"
)

// User ...
func User(c *fiber.Ctx, db ports.Repository) htmx.ContextFunc {
	return func(ctx context.Context) (interface{}, interface{}, error) {
		session, err := goth.SessionFromContext(c)
		if err != nil {
			return nil, nil, err
		}

		user, err := db.GetUserByID(c.Context(), session.UserID)
		if err != nil {
			return err, nil, err
		}

		return ValuesKeyUser, user, nil
	}
}
