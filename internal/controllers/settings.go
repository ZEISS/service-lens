package controllers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// Settings ...
type Settings struct {
	db ports.Repository
}

// NewSettingsController ...
func NewSettingsController(db ports.Repository) *Settings {
	return &Settings{db}
}

// List ...
func (a *Settings) List(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		components.PageProps{}.WithContext(c),
		components.SubNav(
			components.SubNavProps{},
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
	), nil
}
