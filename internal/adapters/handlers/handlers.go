package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/zeiss/service-lens/internal/controllers/dashboard"
	"github.com/zeiss/service-lens/internal/controllers/designs"
	"github.com/zeiss/service-lens/internal/controllers/designs/comments"
	designs_edit_body "github.com/zeiss/service-lens/internal/controllers/designs/edit/body"
	design_edit_title "github.com/zeiss/service-lens/internal/controllers/designs/edit/title"
	"github.com/zeiss/service-lens/internal/controllers/environments"
	"github.com/zeiss/service-lens/internal/controllers/lenses"
	"github.com/zeiss/service-lens/internal/controllers/login"
	"github.com/zeiss/service-lens/internal/controllers/me"
	"github.com/zeiss/service-lens/internal/controllers/previews"
	"github.com/zeiss/service-lens/internal/controllers/profiles"
	"github.com/zeiss/service-lens/internal/controllers/settings"
	"github.com/zeiss/service-lens/internal/controllers/tags"
	"github.com/zeiss/service-lens/internal/controllers/templates"
	"github.com/zeiss/service-lens/internal/controllers/workflows"
	"github.com/zeiss/service-lens/internal/controllers/workloads"
	"github.com/zeiss/service-lens/internal/controllers/workloads/partials"
	"github.com/zeiss/service-lens/internal/controllers/workloads/questions"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

var _ ports.Handlers = (*handlers)(nil)

type handlers struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
}

// New ...
func New(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *handlers {
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

// ListDesigns ...
func (a *handlers) ListDesigns() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return designs.NewListDesignsController(a.store)
	})
}

// NewDesign ...
func (a *handlers) NewDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return designs.NewDesignController(a.store)
	})
}

// ShowDesign ...
func (a *handlers) ShowDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return designs.NewShowDesignController(a.store)
	})
}

// CreateDesignComment ...
func (a *handlers) CreateDesignComment() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return comments.NewCreateDesignCommentController(a.store)
	})
}

// CreateDesign ...
func (a *handlers) CreateDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return designs.NewCreateDesignController(a.store)
	})
}

// UpdateDesign ...
func (a *handlers) UpdateDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return nil
	})
}

// DeleteDesign ...
func (a *handlers) DeleteDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return designs.NewDesignDeleteController(a.store)
	})
}

// EditBodyDesign ...
func (a *handlers) EditBodyDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return designs_edit_body.NewEditController(a.store)
	})
}

// EditTitleDesign ...
func (a *handlers) EditTitleDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return design_edit_title.NewEditController(a.store)
	})
}

// UpdateTitleDesign ...
func (a *handlers) UpdateTitleDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return design_edit_title.NewUpdateController(a.store)
	})
}

// UpdateBodyDesign ...
func (a *handlers) UpdateBodyDesign() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return designs_edit_body.NewUpdateController(a.store)
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

// ShowLensQuestion ...
func (a *handlers) ShowLensQuestion() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workloads.NewWorkloadLensEditQuestionController(a.store)
	})
}

// UpdateWorkloadAnswer ...
func (a *handlers) UpdateWorkloadAnswer() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return questions.NewWorkloadUpdateAnswerController(a.store)
	})
}

// ListTags ...
func (a *handlers) ListTags() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return tags.NewTagsListController(a.store)
	})
}

// CreateTag ...
func (a *handlers) CreateTag() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return tags.NewTagController(a.store)
	})
}

// DeleteTag ...
func (a *handlers) DeleteTag() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return tags.NewTagDeleteController(a.store)
	})
}

// ListWorkflows ...
func (a *handlers) ListWorkflows() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return workflows.NewListWorkflowsController(a.store)
	})
}

// Preview ...
func (a *handlers) Preview() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return previews.NewPreviewController(a.store)
	})
}

// ListTemplates ...
func (a *handlers) ListTemplates() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return templates.NewListTemplatesController(a.store)
	})
}
