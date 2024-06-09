package handlers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/controllers/dashboard"
	"github.com/zeiss/service-lens/internal/controllers/environments"
	"github.com/zeiss/service-lens/internal/controllers/login"
	"github.com/zeiss/service-lens/internal/controllers/me"
	"github.com/zeiss/service-lens/internal/controllers/profiles"
	"github.com/zeiss/service-lens/internal/ports"
)

var _ ports.Handlers = (*handlers)(nil)

type handlers struct {
	store ports.Datastore
}

// New ...
func New(store ports.Datastore) *handlers {
	return &handlers{store}
}

// Login ...
func (a *handlers) Login() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return login.NewIndexLoginController()
	})
}

// Dashboard ...
func (a *handlers) Dashboard() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return dashboard.NewShowDashboardController(a.store)
	})
}

// Settings ...
// func (a *handlers) Settings() fiber.Handler {
// 	return htmx.NewHxControllerHandler(func() htmx.Controller {
// 		return settings.NewSettingsIndexController(a.store)
// 	})
// }

// Me ...
func (a *handlers) Me() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return me.NewMeController(a.store)
	})
}

// ListProfiles ...
func (a *handlers) ListProfiles() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return profiles.NewProfilesListController(a.store)
	})
}

// NewProfile ...
func (a *handlers) NewProfile() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return profiles.NewProfileController(a.store)
	})
}

// GetProfile ...
func (a *handlers) GetProfile() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return profiles.NewProfileController(a.store)
	})
}

// EditProfile ...
func (a *handlers) EditProfile() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return profiles.NewProfileEditController(a.store)
	})
}

// CreateProfile ...
func (a *handlers) CreateProfile() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return profiles.NewCreateProfileController(a.store)
	})
}

// ListEnvironments ...
func (a *handlers) ListEnvironments() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return environments.NewEnvironmentListController(a.store)
	})
}

// NewEnvironment ...
func (a *handlers) NewEnvironment() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return environments.NewEnvironmentController(a.store)
	})
}

// ShowEnvironment ...
func (a *handlers) ShowEnvironment() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return environments.NewEnvironmentShowController(a.store)
	})
}

// EditEnvironment ...
func (a *handlers) EditEnvironment() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return environments.NewEnvironmentEditController(a.store)
	})
}

// UpdateEnvironment ...
func (a *handlers) UpdateEnvironment() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return environments.NewEnvironmentUpdateController(a.store)
	})
}

// DeleteEnvironment ...
func (a *handlers) DeleteEnvironment() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return environments.NewEnvironmentDeleteController(a.store)
	})
}

// CreateEnvironment ...
func (a *handlers) CreateEnvironment() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return environments.NewCreateEnvironmentController(a.store)
	})
}

// // Team ...

// 		teams := app.Group("/teams")
// 		teams.Get("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))
// 		teams.Post("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))

// 		// Team ...
// 		team := teams.Group("/:team")
// 		team.Get("/index", htmx.NewHxControllerHandler(controllers.NewTeamDashboardController(a.db)))

// 		// Profiles ...
// 		profiles := team.Group("/profiles")
// 		profiles.Get(
// 			"/list",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewProfileListController(a.db),
// 			),
// 		)
// 		profiles.Get(
// 			"/new",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewProfileNewController(a.db),
// 			),
// 		)
// 		profiles.Post(
// 			"/new",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewProfileNewController(a.db),
// 			),
// 		)
// 		profiles.Get(
// 			"/:id",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewProfileIndexController(a.db),
// 			),
// 		)
// 		profiles.Delete(
// 			"/:id",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewProfileIndexController(a.db),
// 			),
// 		)
// 		profiles.Get(
// 			"/:id/edit",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewProfileEditController(a.db),
// 			),
// 		)
// 		profiles.Post(
// 			"/:id/edit",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewProfileEditController(a.db),
// 			),
// 		)

// 		// Environments ...
// 		environments := team.Group("/environments")
// 		environments.Get(
// 			"/list",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewEnvironmentListController(a.db),
// 			),
// 		)
// 		environments.Get(
// 			"/new",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewEnvironmentNewController(a.db),
// 			),
// 		)
// 		environments.Post(
// 			"/new",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewEnvironmentNewController(a.db),
// 			),
// 		)
// 		environments.Get(
// 			"/:id",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewEnvironmentIndexController(a.db),
// 			),
// 		)
// 		environments.Delete(
// 			"/:id",
// 			htmx.NewHxControllerHandler(controllers.NewEnvironmentIndexController(a.db)),
// 		)
// 		environments.Get(
// 			"/:id/edit",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewEnvironmentEditController(a.db),
// 			),
// 		)
// 		environments.Post(
// 			"/:id/edit",
// 			htmx.NewHxControllerHandler(controllers.NewEnvironmentEditController(a.db)),
// 		)

// 		// Lenses ...
// 		lenses := team.Group("/lenses")
// 		lenses.Get(
// 			"/list",
// 			htmx.NewHxControllerHandler(controllers.NewLensListController(a.db)),
// 		)
// 		lenses.Get(
// 			"/new",
// 			htmx.NewHxControllerHandler(
// 				controllers.NewLensNewController(a.db),
// 			),
// 		)
// 		lenses.Post(
// 			"/new",
// 			htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db)),
// 		)
// 		lenses.Get(
// 			"/:id/index",
// 			htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db)),
// 		)
// 		lenses.Delete(
// 			"/:id",
// 			htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db)),
// 		)
// 		lenses.Get(
// 			"/:id/edit",
// 			htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db)),
// 		)
// 		lenses.Post(
// 			"/:id/edit",
// 			htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db)),
// 		)

// 		// Workloads ...
// 		workloads := team.Group("/workloads")
// 		workloads.Get(
// 			"/list",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadListController(a.db)),
// 		)
// 		workloads.Get(
// 			"/new",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadNewController(a.db)),
// 		)
// 		workloads.Post(
// 			"/new",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadNewController(a.db)),
// 		)
// 		workloads.Get(
// 			"/:id",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadIndexController(a.db)),
// 		)
// 		workloads.Delete(
// 			"/:id",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadIndexController(a.db)),
// 		)
// 		workloads.Get(
// 			"/:id/lenses/:lens/edit",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadLensEditController(a.db)),
// 		)
// 		workloads.Get(
// 			"/:id/lenses/:lens/edit/:question",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadLensEditController(a.db)),
// 		)
// 		workloads.Post(
// 			"/:id/lenses/:lens/edit/:question",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadLensQuestionUpdateController(a.db)),
// 		)
// 		workloads.Get(
// 			"/:id/lenses/:lens/pillars/:pillar",
// 			htmx.NewHxControllerHandler(controllers.NewWorkloadPillarController(a.db)),
// 		)

// 		// teams.Get(
// 		// 	"/:team/index",
// 		// 	authz.NewTBACHandler(
// 		// 		htmx.NewHxControllerHandler(
// 		// 			controllers.NewTeamDashboardController(a.db),
// 		// 			utils.Resolvers(
// 		// 				resolvers.User(a.db),
// 		// 				resolvers.Team(a.db),
// 		// 			),
// 		// 		),
// 		// 		utils.PermissionView, "team",
// 		// 		authzConfig,
// 		// 	),
// 		// )

// 		// team := app.Group("/:team")
// 		// team.Get(
// 		// 	"/",
// 		// 	authz.NewTBACHandler(
// 		// 		htmx.NewHxControllerHandler(controllers.NewTeamDashboardController(a.db), config),
// 		// 		authz.Read, "team",
// 		// 		authzConfig),
// 		// )

// 		// profiles := team.Group("/profiles")
// 		// profiles.Get("/list", htmx.NewHxControllerHandler(controllers.NewProfileListController(a.db), config))
// 		// profiles.Get("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), config))
// 		// profiles.Post("/new", htmx.NewHxControllerHandler(controllers.NewProfileNewController(a.db), config))
// 		// profiles.Get("/:id", htmx.NewHxControllerHandler(controllers.NewProfileIndexController(a.db), config))
// 		// profiles.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewProfileIndexController(a.db), config))
// 		// profiles.Get("/:id/edit", htmx.NewHxControllerHandler(controllers.NewProfileEditController(a.db), config))
// 		// profiles.Post("/:id/edit", htmx.NewHxControllerHandler(controllers.NewProfileEditController(a.db), config))

// 		// lenses := team.Group("/lenses")
// 		// lenses.Get("/list", htmx.NewHxControllerHandler(controllers.NewLensListController(a.db), config))
// 		// lenses.Get("/new", htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db), config))
// 		// lenses.Post("/new", htmx.NewHxControllerHandler(controllers.NewLensNewController(a.db), config))
// 		// lenses.Get("/:id", htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db), config))
// 		// lenses.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewLensIndexController(a.db), config))
// 		// lenses.Get("/:id/edit", htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db), config))
// 		// lenses.Post("/:id/edit", htmx.NewHxControllerHandler(controllers.NewLensEditController(a.db), config))

// 		site := app.Group("/site")
// 		siteSettings := site.Group("/settings")
// 		siteSettings.Get("/", htmx.NewHxControllerHandler(controllers.NewSettingsIndexController(a.db)))

// 		siteTeams := site.Group("/teams")
// 		siteTeams.Get("/", htmx.NewHxControllerHandler(controllers.NewTeamListController(a.db)))
// 		siteTeams.Get("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))
// 		siteTeams.Post("/new", htmx.NewHxControllerHandler(controllers.NewTeamNewController(a.db)))
// 		siteTeams.Get("/:id", htmx.NewHxControllerHandler(controllers.NewTeamIndexController(a.db)))
// 		siteTeams.Delete("/:id", htmx.NewHxControllerHandler(controllers.NewTeamIndexController(a.db)))
