package services

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/adapters/handlers"
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
	pc  ports.Profiles
	lc  ports.Lenses
	db  ports.Repository
}

// New returns a new instance of NoopSrv.
func New(cfg *configs.Config, pc ports.Profiles, lc ports.Lenses, db ports.Repository) *WebSrv {
	return &WebSrv{cfg, pc, lc, db}
}

// Start starts the server.
func (a *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		indexHandler := handlers.NewIndexHandler()
		profilesHandler := handlers.NewProfilesHandler(a.pc, a.db)
		lensesHandler := handlers.NewLensesHandler(a.lc, a.db)

		workloadController := controllers.NewWorkloadsController(a.db)

		app.Get("/", indexHandler.Index())

		profiles := app.Group("/profiles")
		profiles.Get("/list", htmx.NewCompFuncHandler(profilesHandler.List))
		profiles.Get("/new", profilesHandler.New())
		profiles.Get("/:id", profilesHandler.GetProfile())
		profiles.Post("/new", htmx.NewHtmxHandler(profilesHandler.NewProfile()))

		lenses := app.Group("/lenses")
		lenses.Get("/list", htmx.NewCompFuncHandler(lensesHandler.List))
		lenses.Get("/new", lensesHandler.New())
		lenses.Post("/new", htmx.NewHtmxHandler(lensesHandler.NewLens()))
		lenses.Get("/:id", lensesHandler.GetLens())

		workloads := app.Group("/workloads")
		workloads.Get("/list", htmx.NewCompFuncHandler(workloadController.List))
		workloads.Post("/search", htmx.NewHtmxHandler(workloadController.Search))
		workloads.Get("/new", htmx.NewCompFuncHandler(workloadController.New))
		workloads.Post("/new", htmx.NewHtmxHandler(workloadController.Store))
		workloads.Get("/:id", htmx.NewCompFuncHandler(workloadController.Show))
		workloads.Delete("/:id", htmx.NewHtmxHandler(workloadController.Destroy))

		err := app.Listen(a.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
