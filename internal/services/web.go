package services

import (
	"context"
	"net/http"
	"os"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/controllers"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"
	"github.com/zeiss/service-lens/internal/utils"
	"github.com/zeiss/service-lens/static"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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

		config := htmx.Config{
			Resolvers: []htmx.ResolveFunc{
				resolvers.Team(a.db),
				resolvers.User(a.db),
			},
		}

		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		app.Use("/static", filesystem.New(filesystem.Config{
			Root: http.FS(static.Assets),
		}))

		app.Use(goth.NewProtectMiddleware(gothConfig))
		app.Use(authz.SetAuthzHandler(authz.NewNoopObjectResolver(), authz.NewNoopActionResolver(), authz.NewGothAuthzPrincipalResolver()))

		app.Get("/", htmx.NewHxControllerHandler(controllers.NewDashboardController(a.db), htmx.Config{
			Resolvers: []htmx.ResolveFunc{
				resolvers.User(a.db),
			},
		}))
		app.Get("/login", htmx.NewHxControllerHandler(controllers.NewLoginIndexController(a.db)))
		app.Get("/login/:provider", goth.NewBeginAuthHandler(gothConfig))
		app.Get("/auth/:provider/callback", goth.NewCompleteAuthHandler(gothConfig))
		app.Get("/logout", goth.NewLogoutHandler(gothConfig))
		app.Get("/settings", htmx.NewHxControllerHandler(controllers.NewSettingsIndexController(a.db), htmx.Config{
			Resolvers: []htmx.ResolveFunc{
				resolvers.User(a.db),
			},
		}))

		me := app.Group("/me")
		me.Get("/", htmx.NewHxControllerHandler(controllers.NewMeIndexController(a.db), htmx.Config{
			Resolvers: []htmx.ResolveFunc{
				resolvers.User(a.db),
			},
		}))

		team := app.Group("/:team")
		team.Get("/", htmx.NewHxControllerHandler(controllers.NewTeamDashboardController(a.db), config))

		environments := team.Group("/environments")
		environments.Get("/list", htmx.NewHxControllerHandler(controllers.NewEnvironmentListController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))
		environments.Get("/new", htmx.NewHxControllerHandler(controllers.NewEnvironmentNewController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))
		environments.Post("/new", htmx.NewHxControllerHandler(controllers.NewEnvironmentNewController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))
		environments.Get("/:id", htmx.NewHxControllerHandler(controllers.NewEnvironmentIndexController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))
		environments.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewEnvironmentIndexController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))
		environments.Get("/:id/edit", htmx.NewHxControllerHandler(controllers.NewEnvironmentEditController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))
		environments.Post("/:id/edit", htmx.NewHxControllerHandler(controllers.NewEnvironmentEditController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))

		profiles := team.Group("/profiles")
		profiles.Get("/list", htmx.NewHxControllerHandler(controllers.NewProfileListController(a.db), config))
		profiles.Get("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), config))
		profiles.Post("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), config))
		profiles.Get("/:id", htmx.NewHxControllerHandler(controllers.NewProfileIndexController(a.db), config))
		profiles.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewProfileIndexController(a.db), config))
		profiles.Get("/:id/edit", htmx.NewHxControllerHandler(controllers.NewProfileEditController(a.db), config))
		profiles.Post("/:id/edit", htmx.NewHxControllerHandler(controllers.NewProfileEditController(a.db), config))

		lenses := team.Group("/lenses")
		lenses.Get("/list", htmx.NewHxControllerHandler(controllers.NewLensListController(a.db), config))
		lenses.Get("/new", htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db), config))
		lenses.Post("/new", htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db), config))
		lenses.Get("/:id", htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db), config))
		lenses.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db), config))
		lenses.Get("/:id/edit", htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db), config))
		lenses.Post("/:id/edit", htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db), config))

		workloads := team.Group("/workloads")
		workloads.Get("/", htmx.NewHxControllerHandler(controllers.NewWorkloadListController(a.db), utils.Resolvers(resolvers.Team(a.db), resolvers.User(a.db))))
		workloads.Get("/new", htmx.NewHxControllerHandler(controllers.NewWorkloadNewController(a.db), config))
		workloads.Post("/new", htmx.NewHxControllerHandler(controllers.NewWorkloadNewController(a.db), config))
		workloads.Get("/:id", htmx.NewHxControllerHandler(controllers.NewWorkloadIndexController(a.db), config))
		workloads.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewWorkloadIndexController(a.db), config))
		workloads.Get("/:id/lenses/:lens/edit", htmx.NewHxControllerHandler(controllers.NewWorkloadLensEditController(a.db), config))
		workloads.Get("/:id/lenses/:lens/edit/:question", htmx.NewHxControllerHandler(controllers.NewWorkloadLensEditController(a.db), config))
		workloads.Post("/:id/lenses/:lens/edit/:question", htmx.NewHxControllerHandler(controllers.NewWorkloadLensQuestionUpdateController(a.db), config))
		workloads.Get("/:id/lenses/:lens/pillars/:pillar", htmx.NewHxControllerHandler(controllers.NewWorkloadPillarController(a.db), config))

		site := app.Group("/site")
		siteSettings := site.Group("/settings")
		siteSettings.Get("/", htmx.NewHxControllerHandler(controllers.NewSettingsIndexController(a.db), utils.Resolvers(resolvers.User(a.db))))

		siteTeams := site.Group("/teams")
		siteTeams.Get("/", htmx.NewHxControllerHandler(controllers.NewTeamListController(a.db), utils.Resolvers(resolvers.User(a.db))))
		siteTeams.Get("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db), utils.Resolvers(resolvers.User(a.db))))
		siteTeams.Post("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db), utils.Resolvers(resolvers.User(a.db))))
		siteTeams.Get("/:id", htmx.NewHxControllerHandler(controllers.NewTeamIndexController(a.db), utils.Resolvers(resolvers.User(a.db))))
		siteTeams.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewTeamIndexController(a.db), utils.Resolvers(resolvers.User(a.db))))

		err := app.Listen(a.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
