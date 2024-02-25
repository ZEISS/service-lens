package services

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/controllers"
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
	db  ports.Repository
}

// New returns a new instance of NoopSrv.
func New(cfg *configs.Config, db ports.Repository) *WebSrv {
	return &WebSrv{cfg, db}
}

// Start starts the server.
func (a *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		homeController := controllers.NewHomeController(a.db)
		profilesController := controllers.NewProfilesController(a.db)
		lensesController := controllers.NewLensesController(a.db)

		workloadController := controllers.NewWorkloadsController(a.db)
		workloadLensController := controllers.NewLensController(a.db)
		settingsController := controllers.NewSettingsController(a.db)

		teamsController := controllers.NewTeamsController(a.db)

		app.Get("/", htmx.NewCompFuncHandler(homeController.Index))

		teams := app.Group("/teams")
		teams.Get("/new", htmx.NewCompFuncHandler(teamsController.New))
		teams.Post("/new", htmx.NewHtmxHandler(teamsController.Store))
		teams.Get("/:id", htmx.NewCompFuncHandler(teamsController.Show))

		team := app.Group("/:team")

		profiles := team.Group("/profiles")
		profiles.Get("/list", htmx.NewCompFuncHandler(profilesController.List))
		profiles.Get("/new", htmx.NewCompFuncHandler(profilesController.New))
		profiles.Post("/new", htmx.NewHtmxHandler(profilesController.Store))
		profiles.Get("/:id", htmx.NewCompFuncHandler(profilesController.Show))

		lenses := team.Group("/lenses")
		lenses.Get("/list", htmx.NewCompFuncHandler(lensesController.List))
		lenses.Get("/new", htmx.NewCompFuncHandler(lensesController.New))
		lenses.Post("/new", htmx.NewHtmxHandler(lensesController.Store))
		lenses.Get("/:id", htmx.NewCompFuncHandler(lensesController.Show))

		workloads := team.Group("/workloads")
		workloads.Get("/list", htmx.NewCompFuncHandler(workloadController.List))
		workloads.Post("/search", htmx.NewHtmxHandler(workloadController.Search))
		workloads.Get("/new", htmx.NewCompFuncHandler(workloadController.New))
		workloads.Post("/new", htmx.NewHtmxHandler(workloadController.Store))
		workloads.Get("/:id", htmx.NewCompFuncHandler(workloadController.Show))
		workloads.Delete("/:id", htmx.NewHtmxHandler(workloadController.Destroy))

		workloadLens := workloads.Group("/:id/lens/:lens")
		workloadLens.Get("/list", htmx.NewCompFuncHandler(workloadLensController.List))

		settings := app.Group("/settings")
		settings.Get("/list", htmx.NewCompFuncHandler(settingsController.List))

		err := app.Listen(a.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
