package controllers

import (
	"github.com/zeiss/service-lens/internal/controllers/dashboard"
	"github.com/zeiss/service-lens/internal/controllers/lenses"
	"github.com/zeiss/service-lens/internal/controllers/login"
	"github.com/zeiss/service-lens/internal/controllers/me"
	"github.com/zeiss/service-lens/internal/controllers/profiles"
	"github.com/zeiss/service-lens/internal/controllers/settings"
	"github.com/zeiss/service-lens/internal/controllers/teams"
	"github.com/zeiss/service-lens/internal/controllers/workloads"
	"github.com/zeiss/service-lens/internal/ports"
)

// NewDashboardIndexController ...
func NewDashboardController(db ports.Repository) *dashboard.DashboardIndexController {
	return dashboard.NewDashboardController(db)
}

// NewLoginIndexController ...
func NewLoginIndexController(db ports.Repository) *login.LoginIndexController {
	return login.NewLoginIndexController(db)
}

// NewLensNewController ...
func NewLensNewController(db ports.Repository) *lenses.LensNewController {
	return lenses.NewLensNewController(db)
}

// NewLensIndexController ...
func NewLensIndexController(db ports.Repository) *lenses.LensIndexController {
	return lenses.NewLensIndexController(db)
}

// NewLensListController ...
func NewLensListController(db ports.Repository) *lenses.LensListController {
	return lenses.NewLensListController(db)
}

// NewWorkloadIndexController ...
func NewWorkloadIndexController(db ports.Repository) *workloads.WorkloadIndexController {
	return workloads.NewWorkloadIndexController(db)
}

// NewMeIndexController ...
func NewMeIndexController(db ports.Repository) *me.MeIndexController {
	return me.NewMeIndexController(db)
}

// NewSettingsIndexController ...
func NewSettingsIndexController(db ports.Repository) *settings.SettingsIndexController {
	return settings.NewSettingsIndexController(db)
}

// NewTeamIndexController ...
func NewTeamIndexController(db ports.Repository) *teams.TeamIndexController {
	return teams.NewTeamsIndexController(db)
}

// NewTeamsNewController ...
func NewTeamsNewController(db ports.Repository) *teams.TeamsNewController {
	return teams.NewTeamsNewController(db)
}

// NewTeamDashboardController ...
func NewTeamDashboardController(db ports.Repository) *teams.TeamDashboardController {
	return teams.NewTeamDashboardController(db)
}

// NewProfileListController ...
func NewProfileListController(db ports.Repository) *profiles.ProfileListController {
	return profiles.NewProfileListController(db)
}

// NewProfileNewController ...
func NewProfileNewController(db ports.Repository) *profiles.ProfileNewController {
	return profiles.NewProfileNewController(db)
}

// NewProfileIndexController ...
func NewProfileIndexController(db ports.Repository) *profiles.ProfileIndexController {
	return profiles.NewProfileIndexController(db)
}

// NewWorkloadListController ...
func NewWorkloadListController(db ports.Repository) *workloads.WorkloadListController {
	return workloads.NewWorkloadListController(db)
}

// NewWorkloadNewController ...
func NewWorkloadNewController(db ports.Repository) *workloads.WorkloadNewController {
	return workloads.NewWorkloadsNewController(db)
}

// NewWorkloadLensController ...
func NewWorkloadLensController(db ports.Repository) *workloads.WorkloadLensController {
	return workloads.NewWorkloadLensController(db)
}

// NewWorkloadPillarController ...
func NewWorkloadPillarController(db ports.Repository) *workloads.WorkloadPillarController {
	return workloads.NewWorkloadPillarController(db)
}
