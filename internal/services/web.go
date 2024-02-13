package services

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/adapters/handlers"
	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
)

var _ server.Listener = (*WebSrv)(nil)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg *configs.Config
	pc  ports.Profiles
}

// New returns a new instance of NoopSrv.
func New(cfg *configs.Config, pc ports.Profiles) *WebSrv {
	return &WebSrv{cfg, pc}
}

// Start starts the server.
func (a *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		indexHandler := handlers.NewIndexHandler()
		profilesHandler := handlers.NewProfilesHandler(a.pc)

		app.Get("/", indexHandler.Index())
		app.Get("/profiles", profilesHandler.Index())
		app.Get("/profiles/:id", profilesHandler.GetProfile())
		app.Get("/profiles/new", profilesHandler.New())
		app.Post("/profiles", htmx.NewHtmxHandler(profilesHandler.NewProfile()))

		err := app.Listen(a.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
