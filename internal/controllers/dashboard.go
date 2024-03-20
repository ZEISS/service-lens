package controllers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// Dashboard ...
type Dashboard struct {
	db ports.Repository
}

// NewDashboardController ...
func NewDashboardController(db ports.Repository) *Dashboard {
	return &Dashboard{db}
}

// Show ...
func (d *Dashboard) Index(c *fiber.Ctx) (htmx.Node, error) {
	ctx := htmx.FromContext(c)

	return components.Page(
		ctx,
		components.PageProps{},
		components.Layout(
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
	), nil
}
