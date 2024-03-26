package services

import (
	"context"
	"os"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/controllers"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	goth "github.com/zeiss/fiber-goth"
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-goth/providers"
	"github.com/zeiss/fiber-goth/providers/github"
)

var _ server.Listener = (*WebSrv)(nil)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg     *configs.Config
	db      ports.Repository
	adapter adapters.Adapter
}

// New returns a new instance of NoopSrv.
func New(cfg *configs.Config, db ports.Repository, adapter adapters.Adapter) *WebSrv {
	return &WebSrv{cfg, db, adapter}
}

// Start starts the server.
func (a *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		providers.RegisterProvider(github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:8080/auth/github/callback"))

		gothConfig := goth.Config{
			Adapter:        a.adapter,
			Secret:         goth.GenerateKey(),
			CookieHTTPOnly: true,
		}

		teamConfig := htmx.Config{
			Resolvers: []htmx.ResolveFunc{
				resolvers.Team(a.db),
			},
		}

		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		app.Use(goth.NewProtectMiddleware(gothConfig))
		app.Use(authz.SetAuthzHandler(authz.NewNoopObjectResolver(), authz.NewNoopActionResolver(), authz.NewGothAuthzPrincipalResolver()))

		app.Get("/", htmx.NewHxControllerHandler(controllers.NewDashboardController(a.db)))
		app.Get("/login", htmx.NewHxControllerHandler(controllers.NewLoginIndexController(a.db)))
		app.Get("/login/:provider", goth.NewBeginAuthHandler(gothConfig))
		app.Get("/auth/:provider/callback", goth.NewCompleteAuthHandler(gothConfig))
		app.Get("/logout", goth.NewLogoutHandler(gothConfig))

		me := app.Group("/me")
		me.Get("/index", htmx.NewHxControllerHandler(controllers.NewMeIndexController(a.db)))

		teams := app.Group("/teams")
		teams.Get("/new", htmx.NewHxControllerHandler(controllers.NewTeamsNewController(a.db)))
		teams.Post("/new", htmx.NewHxControllerHandler(controllers.NewTeamsNewController(a.db)))
		teams.Get("/:id", htmx.NewHxControllerHandler(controllers.NewTeamIndexController(a.db)))

		team := app.Group("/:team")
		team.Get("/index", htmx.NewHxControllerHandler(controllers.NewTeamDashboardController(a.db), teamConfig))

		profiles := team.Group("/profiles")
		profiles.Get("/list", htmx.NewHxControllerHandler(controllers.NewProfileListController(a.db), teamConfig))
		profiles.Get("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), teamConfig))
		profiles.Post("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), teamConfig))
		profiles.Get("/:id", htmx.NewHxControllerHandler(controllers.NewProfileIndexController(a.db), teamConfig))

		// lenses := team.Group("/lenses")
		// lenses.Get("/list", htmx.NewHxControllerHandler(lensesController.List))
		// lenses.Get("/new", htmx.NewHxControllerHandler(lensesController.New))
		// lenses.Post("/new", htmx.NewHxControllerHandler(lensesController.Store))
		// lenses.Get("/:id", htmx.NewHxControllerHandler(lensesController.Show))

		// workloads := team.Group("/workloads")
		// workloads.Get("/list", htmx.NewHxControllerHandler(workloadController.List))
		// workloads.Post("/search", htmx.NewHxControllerHandler(workloadController.Search))
		// workloads.Get("/new", htmx.NewHxControllerHandler(workloadController.New))
		// workloads.Post("/new", htmx.NewHxControllerHandler(workloadController.Store))
		// workloads.Get("/:id", htmx.NewHxControllerHandler(workloadController.Show))
		// workloads.Delete("/:id", htmx.NewHxControllerHandler(workloadController.Destroy))

		// workloadLens := workloads.Group("/:id/lens/:lens")
		// workloadLens.Get("/list", htmx.NewHxControllerHandler(workloadLensController.List))

		site := app.Group("/site")
		siteSettings := site.Group("/settings")
		siteSettings.Get("/index", htmx.NewHxControllerHandler(controllers.NewSettingsIndexController(a.db)))

		err := app.Listen(a.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
