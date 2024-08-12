package ports

import "github.com/gofiber/fiber/v2"

// Handlers ...
type Handlers interface {
	// Login ...
	Login() fiber.Handler
	// Dashboard ...
	Dashboard() fiber.Handler
	// Me ...
	Me() fiber.Handler
	// ListDesigns ...
	ListDesigns() fiber.Handler
	// NewDesign ...
	NewDesign() fiber.Handler
	// ShowDesign ...
	ShowDesign() fiber.Handler
	// CreateDesign ...
	CreateDesign() fiber.Handler
	// CreateDesignComment ...
	CreateDesignComment() fiber.Handler
	// UpdateDesign ...
	UpdateDesign() fiber.Handler
	// DeleteDesign ...
	DeleteDesign() fiber.Handler
	// ListProfiles ...
	ListProfiles() fiber.Handler
	// EditBodyDesign ...
	EditBodyDesign() fiber.Handler
	// UpdateBodyDesign ...
	UpdateBodyDesign() fiber.Handler
	// EditTitleDesign ...
	EditTitleDesign() fiber.Handler
	// UpdateTitleDesign ...
	UpdateTitleDesign() fiber.Handler
	// NewProfile ...
	NewProfile() fiber.Handler
	// ShowProfile ...
	ShowProfile() fiber.Handler
	// EditProfile ...
	EditProfile() fiber.Handler
	// CreateProfile ...
	CreateProfile() fiber.Handler
	// DeleteProfile ...
	DeleteProfile() fiber.Handler
	// ListEnvironments ...
	ListEnvironments() fiber.Handler
	// NewEnvironment ...
	NewEnvironment() fiber.Handler
	// ShowEnvironment ...
	ShowEnvironment() fiber.Handler
	// EditEnvironment ...
	EditEnvironment() fiber.Handler
	// UpdateEnvironment ...
	UpdateEnvironment() fiber.Handler
	// DeleteEnvironment ...
	DeleteEnvironment() fiber.Handler
	// CreateEnvironment ...
	CreateEnvironment() fiber.Handler
	// ListLenses ...
	ListLenses() fiber.Handler
	// NewLens ...
	NewLens() fiber.Handler
	// ShowLens ...
	ShowLens() fiber.Handler
	// EditLens ...
	EditLens() fiber.Handler
	// UpdateLens ...
	UpdateLens() fiber.Handler
	// DeleteLens ...
	DeleteLens() fiber.Handler
	// CreateLens ...
	CreateLens() fiber.Handler
	// ShowSettings ...
	ShowSettings() fiber.Handler
	// CreateWorkload ...
	CreateWorkload() fiber.Handler
	// ListWorkloads ...
	ListWorkloads() fiber.Handler
	// ShowWorkload ...
	ShowWorkload() fiber.Handler
	// EditWorkload ...
	EditWorkload() fiber.Handler
	// NewWorkload ...
	NewWorkload() fiber.Handler
	// UpdateWorkload ...
	// UpdateWorkload() fiber.Handler
	// DeleteWorkload ...
	DeleteWorkload() fiber.Handler
	// ListEnvironmentsPartial ...
	ListEnvironmentsPartial() fiber.Handler
	// ListProfilesPartial ...
	ListProfilesPartial() fiber.Handler
	// ShowWorkloadLens ...
	ShowWorkloadLens() fiber.Handler
	// EditWorkloadLens ...
	EditWorkloadLens() fiber.Handler
	// ShowLensQuestion ...
	ShowLensQuestion() fiber.Handler
	// UpdateWorkloadAnswer ...
	UpdateWorkloadAnswer() fiber.Handler
	// DeleteTag ...
	DeleteTag() fiber.Handler
	// ListWorkflows ...
	ListWorkflows() fiber.Handler
	// Preview ...
	Preview() fiber.Handler
	// ListTemplates ...
	ListTemplates() fiber.Handler
}
