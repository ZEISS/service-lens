package dashboard

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// DashboardIndexController ...
type DashboardIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewDashboardIndexController ...
func NewDashboardController(db ports.Repository) *DashboardIndexController {
	return &DashboardIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (d *DashboardIndexController) Get() error {
	return d.Hx.RenderComp(
		components.Page(
			d.Hx,
			components.PageProps{},
			components.Layout(
				d.Hx,
				components.LayoutProps{},
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
						),
					),
				),
				components.Wrap(
					components.WrapProps{},
				),
			),
		),
	)
}
