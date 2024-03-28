package teams

import (
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"
)

// TeamDashboardController ...
type TeamDashboardController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewTeamDashboardController ...
func NewTeamDashboardController(db ports.Repository) *TeamDashboardController {
	return &TeamDashboardController{db, htmx.UnimplementedController{}}
}

// Get ...
func (t *TeamDashboardController) Get() error {
	team := t.Hx().Values(resolvers.ValuesKeyTeam).(*authz.Team)

	return t.Hx().RenderComp(
		components.Page(
			t.Hx(),
			components.PageProps{},
			components.Layout(
				t.Hx(),
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
					htmx.Div(
						htmx.Text(team.Name),
					),
				),
			),
		),
	)
}
