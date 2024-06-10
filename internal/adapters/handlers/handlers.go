package handlers

import (
	"github.com/gofiber/fiber/v2"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/controllers/dashboard"
	"github.com/zeiss/service-lens/internal/controllers/environments"
	"github.com/zeiss/service-lens/internal/controllers/lenses"
	"github.com/zeiss/service-lens/internal/controllers/login"
	"github.com/zeiss/service-lens/internal/controllers/me"
	"github.com/zeiss/service-lens/internal/controllers/profiles"
	"github.com/zeiss/service-lens/internal/controllers/settings"
	"github.com/zeiss/service-lens/internal/controllers/workloads"
	"github.com/zeiss/service-lens/internal/controllers/workloads/partials"
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

// ShowProfile ...
func (a *handlers) ShowProfile() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return profiles.NewProfileShowController(a.store)
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

// DeleteProfile ...
func (a *handlers) DeleteProfile() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return profiles.NewProfileDeleteController(a.store)
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

// ListLenses ...
func (a *handlers) ListLenses() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return lenses.NewLensListController(a.store)
	})
}

// NewLens ...
func (a *handlers) NewLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return lenses.NewLensController(a.store)
	})
}

// ShowLens ...
func (a *handlers) ShowLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return lenses.NewLensShowController(a.store)
	})
}

// EditLens ...
func (a *handlers) EditLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return lenses.NewLensEditController(a.store)
	})
}

// UpdateLens ...
func (a *handlers) UpdateLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return lenses.NewLensEditController(a.store)
	})
}

// DeleteLens ...
func (a *handlers) DeleteLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return lenses.NewLensDeleteController(a.store)
	})
}

// CreateLens ...
func (a *handlers) CreateLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return lenses.NewCreateLensController(a.store)
	})
}

// ShowSettings ...
func (a *handlers) ShowSettings() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return settings.NewSettingsShowController(a.store)
	})
}

// NewWorkload ...
func (a *handlers) NewWorkload() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadController(a.store)
	})
}

// CreateWorkload ...
func (a *handlers) CreateWorkload() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewCreateWorkloadController(a.store)
	})
}

// ListWorkloads ...
func (a *handlers) ListWorkloads() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadListController(a.store)
	})
}

// ShowWorkload ...
func (a *handlers) ShowWorkload() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadShowController(a.store)
	})
}

// EditWorkload ...
func (a *handlers) EditWorkload() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadEditController(a.store)
	})
}

// DeleteWorkload ...
func (a *handlers) DeleteWorkload() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadDeleteController(a.store)
	})
}

// ListEnvironmentsPartial ...
func (a *handlers) ListEnvironmentsPartial() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return partials.NewEnvironmentPartialListController(a.store)
	})
}

// ListProfilesPartial ...
func (a *handlers) ListProfilesPartial() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return partials.NewProfilePartialListController(a.store)
	})
}

// ShowWorkloadLens ...
func (a *handlers) ShowWorkloadLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadLensController(a.store)
	})
}

// EditWorkloadLens ...
func (a *handlers) EditWorkloadLens() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadLensEditController(a.store)
	})
}
