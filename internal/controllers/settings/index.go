package settings

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// SettingsIndexController ...
type SettingsIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewSettingsIndexController ...
func NewSettingsIndexController(db ports.Repository) *SettingsIndexController {
	return &SettingsIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (a *SettingsIndexController) Get() error {
	return a.Hx.RenderComp(

		components.Page(
			a.Hx,
			components.PageProps{},
			components.SubNav(
				components.SubNavProps{},
				components.SubNavBreadcrumb(
					components.SubNavBreadcrumbProps{},
					breadcrumbs.Breadcrumbs(
						breadcrumbs.BreadcrumbsProps{},
						breadcrumbs.Breadcrumb(
							breadcrumbs.BreadcrumbProps{
								Href:  "/",
								Title: "Home",
							},
						),
						breadcrumbs.Breadcrumb(
							breadcrumbs.BreadcrumbProps{
								Href:  "/settings/list",
								Title: "Settings",
							},
						),
					),
				),
				components.SubNavActions(
					components.SubNavActionsProps{},
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-xs": true,
							},
						},
						htmx.Text("Create Workload"),
					),
				),
			),
		),
	)
}
