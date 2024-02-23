package controllers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// Teams ...
type Teams struct {
	db ports.Repository
}

// NewTeamsController ...
func NewTeamsController(db ports.Repository) *Teams {
	return &Teams{db}
}

// New ...
func (a *Teams) New(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
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
							Href:  "/teams/list",
							Title: "Teams",
						},
					),
				),
			),
		),
		components.Wrap(
			components.WrapProps{},
			htmx.FormElement(
				htmx.HxPost("/teams"),
				htmx.Label(
					htmx.ClassNames{
						"form-control": true,
						"w-full":       true,
						"max-w-lg":     true,
					},
					htmx.Div(
						htmx.ClassNames{
							"label": true,
						},
						htmx.Span(
							htmx.ClassNames{
								"label-text": true,
							},
						),
					),
					htmx.Input(
						htmx.Attribute("type", "text"),
						htmx.Attribute("name", "name"),
						htmx.Attribute("placeholder", "Name ..."),
						htmx.ClassNames{
							"input":          true,
							"input-bordered": true,
							"w-full":         true,
							"max-w-lg":       true,
						},
					),
				),
				htmx.Label(
					htmx.ClassNames{
						"form-control": true,
						"w-full":       true,
						"max-w-lg":     true,
						"mb-4":         true,
					},
					htmx.Div(
						htmx.ClassNames{
							"label": true,
						},
						htmx.Span(
							htmx.ClassNames{
								"label-text": true,
							},
						),
					),
					htmx.Input(
						htmx.Attribute("type", "text"),
						htmx.Attribute("slug", "slug"),
						htmx.Attribute("placeholder", "Slug ..."),
						htmx.ClassNames{
							"input":          true,
							"input-bordered": true,
							"w-full":         true,
							"max-w-lg":       true,
						},
					),
				),
				htmx.Label(
					htmx.ClassNames{
						"form-control": true,
						"w-full":       true,
						"max-w-lg":     true,
					},
					htmx.Div(
						htmx.ClassNames{
							"label":   true,
							"sr-only": true,
						},
					),
					htmx.Input(
						htmx.Attribute("type", "text"),
						htmx.Attribute("name", "description"),
						htmx.Attribute("placeholder", "Description ..."),
						htmx.ClassNames{
							"input":          true,
							"input-bordered": true,
							"w-full":         true,
							"max-w-lg":       true,
						},
					),
				),
				buttons.Outline(
					buttons.ButtonProps{
						ClassNames: htmx.ClassNames{
							"my-4": true,
						},
						Type: "submit",
					},
					htmx.Text("Create Profile"),
				),
			),
		),
	), nil
}
