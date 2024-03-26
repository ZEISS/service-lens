package lenses

import (
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

// LensIndexController ...
type LensIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewLensIndexController ...
func NewLensIndexController(db ports.Repository) *LensIndexController {
	return &LensIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (l *LensIndexController) Get(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		l.Hx,
		components.PageProps{},
		components.Layout(
			l.Hx,
			components.LayoutProps{},
			components.SubNav(
				components.SubNavProps{},
				components.SubNavBreadcrumb(
					components.SubNavBreadcrumbProps{},
					breadcrumbs.Breadcrumb(
						breadcrumbs.BreadcrumbProps{
							Href:  "/",
							Title: "Home",
						},
					),
					breadcrumbs.Breadcrumb(
						breadcrumbs.BreadcrumbProps{
							Href:  "/workloads/list",
							Title: "Workloads",
						},
					),
					breadcrumbs.Breadcrumb(
						breadcrumbs.BreadcrumbProps{
							Href:  "/lenses/list",
							Title: "Lenses",
						},
					),
				),
				components.SubNavActions(
					components.SubNavActionsProps{},
					links.Link(
						links.LinkProps{
							Href: "/lenses/new",
							ClassNames: htmx.ClassNames{
								"btn": true,
							},
						},
						htmx.Text("Create Lens"),
					),
				),
			),
			components.Wrap(
				components.WrapProps{},
				htmx.Div(
					htmx.ClassNames{
						"overflow-x-auto": true,
					},
					htmx.Table(
						htmx.ClassNames{"table": true},
						htmx.THead(
							htmx.Tr(
								htmx.Th(
									htmx.Label(
										htmx.Input(
											htmx.ClassNames{
												"checkbox": true,
											},
											htmx.Attribute("type", "checkbox"),
											htmx.Attribute("name", "all"),
										),
									),
								),
								htmx.Th(htmx.Text("ID")),
								htmx.Th(htmx.Text("Name")),
								htmx.Th(htmx.Text("Description")),
							),
						),
						htmx.TBody(
							htmx.ID("data-table"),
							// htmx.Group(profilesItems...),
						),
					),
				),
			),
		),
	), nil
}
