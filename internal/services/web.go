package services

import (
	"context"
	"net/http"
	"os"

	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/controllers"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
	"github.com/zeiss/service-lens/static"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	authz "github.com/zeiss/fiber-authz"
	goth "github.com/zeiss/fiber-goth"
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-goth/providers"
	"github.com/zeiss/fiber-goth/providers/github"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/filters"
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

		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		app.Use("/static", filesystem.New(filesystem.Config{
			Root: http.FS(static.Assets),
		}))

		app.Use(goth.NewProtectMiddleware(gothConfig))
		app.Use(authz.SetAuthzHandler(authz.NewNoopObjectResolver(), authz.NewNoopActionResolver(), authz.NewGothAuthzPrincipalResolver()))

		app.Get("/", htmx.NewHxControllerHandler(controllers.NewDashboardController(a.db)))
		app.Get("/login", htmx.NewHxControllerHandler(controllers.NewLoginIndexController(a.db)))
		app.Get("/login/:provider", goth.NewBeginAuthHandler(gothConfig))
		app.Get("/auth/:provider/callback", goth.NewCompleteAuthHandler(gothConfig))
		app.Get("/logout", goth.NewLogoutHandler(gothConfig))
		app.Get("/settings", htmx.NewHxControllerHandler(controllers.NewSettingsIndexController(a.db)))

		me := app.Group("/me")
		me.Get("/", htmx.NewHxControllerHandler(controllers.NewMeIndexController(a.db)))

		teams := app.Group("/teams")
		teams.Get("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))
		teams.Post("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))

		// Team ...
		team := teams.Group("/:team")
		team.Get(
			"/index",
			htmx.NewHxControllerHandler(
				controllers.NewTeamDashboardController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)

		// Profiles ...
		profiles := team.Group("/profiles")
		profiles.Get(
			"/list",
			htmx.NewHxControllerHandler(
				controllers.NewProfileListController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		profiles.Get(
			"/new",
			htmx.NewHxControllerHandler(
				controllers.NewProfileNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		profiles.Post(
			"/new",
			htmx.NewHxControllerHandler(
				controllers.NewProfileNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		profiles.Get(
			"/:id",
			htmx.NewHxControllerHandler(
				controllers.NewProfileIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		profiles.Delete(
			"/:id",
			htmx.NewHxControllerHandler(
				controllers.NewProfileIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionDelete, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		profiles.Get(
			"/:id/edit",
			htmx.NewHxControllerHandler(
				controllers.NewProfileEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		profiles.Post(
			"/:id/edit",
			htmx.NewHxControllerHandler(
				controllers.NewProfileEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)

		// Environments ...
		environments := team.Group("/environments")
		environments.Get(
			"/list",
			htmx.NewHxControllerHandler(
				controllers.NewEnvironmentListController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		environments.Get(
			"/new",
			htmx.NewHxControllerHandler(
				controllers.NewEnvironmentNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		environments.Post(
			"/new",
			htmx.NewHxControllerHandler(
				controllers.NewEnvironmentNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		environments.Get(
			"/:id",
			htmx.NewHxControllerHandler(
				controllers.NewEnvironmentIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		environments.Delete(
			"/:id",
			htmx.NewHxControllerHandler(controllers.NewEnvironmentIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionDelete, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		environments.Get(
			"/:id/edit",
			htmx.NewHxControllerHandler(
				controllers.NewEnvironmentEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		environments.Post(
			"/:id/edit",
			htmx.NewHxControllerHandler(controllers.NewEnvironmentEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)

		// Lenses ...
		lenses := team.Group("/lenses")
		lenses.Get(
			"/list",
			htmx.NewHxControllerHandler(controllers.NewLensListController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		lenses.Get(
			"/new",
			htmx.NewHxControllerHandler(
				controllers.NewLensNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		lenses.Post(
			"/new",
			htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		lenses.Get(
			"/:id/index",
			htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		lenses.Delete(
			"/:id",
			htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionDelete, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		lenses.Get(
			"/:id/edit",
			htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		lenses.Post(
			"/:id/edit",
			htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)

		// Workloads ...
		workloads := team.Group("/workloads")
		workloads.Get(
			"/list",
			htmx.NewHxControllerHandler(controllers.NewWorkloadListController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Get(
			"/new",
			htmx.NewHxControllerHandler(controllers.NewWorkloadNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Post(
			"/new",
			htmx.NewHxControllerHandler(controllers.NewWorkloadNewController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionCreate, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Get(
			"/:id",
			htmx.NewHxControllerHandler(controllers.NewWorkloadIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Delete(
			"/:id",
			htmx.NewHxControllerHandler(controllers.NewWorkloadIndexController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionDelete, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Get(
			"/:id/lenses/:lens/edit",
			htmx.NewHxControllerHandler(controllers.NewWorkloadLensEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Get(
			"/:id/lenses/:lens/edit/:question",
			htmx.NewHxControllerHandler(controllers.NewWorkloadLensEditController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Post(
			"/:id/lenses/:lens/edit/:question",
			htmx.NewHxControllerHandler(controllers.NewWorkloadLensQuestionUpdateController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionEdit, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)
		workloads.Get(
			"/:id/lenses/:lens/pillars/:pillar",
			htmx.NewHxControllerHandler(controllers.NewWorkloadPillarController(a.db),
				htmx.Config{
					Filters: []htmx.FilterFunc{
						filters.NewAuthzParamFilter(utils.PermissionView, "team", a.adapter.(authz.AuthzChecker)),
					},
				},
			),
		)

		// teams.Get(
		// 	"/:team/index",
		// 	authz.NewTBACHandler(
		// 		htmx.NewHxControllerHandler(
		// 			controllers.NewTeamDashboardController(a.db),
		// 			utils.Resolvers(
		// 				resolvers.User(a.db),
		// 				resolvers.Team(a.db),
		// 			),
		// 		),
		// 		utils.PermissionView, "team",
		// 		authzConfig,
		// 	),
		// )

		// team := app.Group("/:team")
		// team.Get(
		// 	"/",
		// 	authz.NewTBACHandler(
		// 		htmx.NewHxControllerHandler(controllers.NewTeamDashboardController(a.db), config),
		// 		authz.Read, "team",
		// 		authzConfig),
		// )

		// profiles := team.Group("/profiles")
		// profiles.Get("/list", htmx.NewHxControllerHandler(controllers.NewProfileListController(a.db), config))
		// profiles.Get("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), config))
		// profiles.Post("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), config))
		// profiles.Get("/:id", htmx.NewHxControllerHandler(controllers.NewProfileIndexController(a.db), config))
		// profiles.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewProfileIndexController(a.db), config))
		// profiles.Get("/:id/edit", htmx.NewHxControllerHandler(controllers.NewProfileEditController(a.db), config))
		// profiles.Post("/:id/edit", htmx.NewHxControllerHandler(controllers.NewProfileEditController(a.db), config))

		// lenses := team.Group("/lenses")
		// lenses.Get("/list", htmx.NewHxControllerHandler(controllers.NewLensListController(a.db), config))
		// lenses.Get("/new", htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db), config))
		// lenses.Post("/new", htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db), config))
		// lenses.Get("/:id", htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db), config))
		// lenses.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db), config))
		// lenses.Get("/:id/edit", htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db), config))
		// lenses.Post("/:id/edit", htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db), config))

		site := app.Group("/site")
		siteSettings := site.Group("/settings")
		siteSettings.Get("/", htmx.NewHxControllerHandler(controllers.NewSettingsIndexController(a.db)))

		siteTeams := site.Group("/teams")
		siteTeams.Get("/", htmx.NewHxControllerHandler(controllers.NewTeamListController(a.db)))
		siteTeams.Get("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))
		siteTeams.Post("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))
		siteTeams.Get("/:id", htmx.NewHxControllerHandler(controllers.NewTeamIndexController(a.db)))
		siteTeams.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewTeamIndexController(a.db)))

		err := app.Listen(a.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
