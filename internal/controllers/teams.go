package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
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
				htmx.HxPost("/teams/new"),
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
					htmx.Text("Create Team"),
				),
			),
		),
	), nil
}

// Store ...
func (a *Teams) Store(hx *htmx.Htmx) error {
	team := &authz.Team{
		Name:        hx.Ctx().FormValue("name"),
		Slug:        hx.Ctx().FormValue("slug"),
		Description: utils.StrPtr(hx.Ctx().FormValue("description")),
	}

	team, err := a.db.AddTeam(hx.Ctx().Context(), team)
	if err != nil {
		return err
	}

	hx.Redirect("/teams/" + team.ID.String())

	return nil
}

// Show ...
func (a *Teams) Show(c *fiber.Ctx) (htmx.Node, error) {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return nil, err
	}

	team, err := a.db.GetTeamByID(c.Context(), id)
	if err != nil {
		return nil, err
	}

	return components.Page(
		components.PageProps{}.WithContext(c),
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
					breadcrumbs.Breadcrumb(
						breadcrumbs.BreadcrumbProps{
							Href:  "/teams/" + team.ID.String(),
							Title: team.Name,
						},
					),
				),
			),
		),
		components.Wrap(
			components.WrapProps{},
			htmx.Div(
				htmx.H1(
					htmx.Text(team.Name),
				),
				htmx.P(
					htmx.Text(utils.PtrStr(team.Description)),
				),
				htmx.Div(
					htmx.ClassNames{
						"flex":     true,
						"flex-col": true,
						"py-2":     true,
					},
					htmx.H4(
						htmx.ClassNames{
							"text-gray-500": true,
						},
						htmx.Text("Created at"),
					),
					htmx.H3(
						htmx.Text(
							team.CreatedAt.Format("2006-01-02 15:04:05"),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"flex":     true,
						"flex-col": true,
						"py-2":     true,
					},
					htmx.H4(
						htmx.ClassNames{
							"text-gray-500": true,
						},
						htmx.Text("Updated at"),
					),
					htmx.H3(
						htmx.Text(
							team.UpdatedAt.Format("2006-01-02 15:04:05"),
						),
					),
				),
			),
		),
	), nil
}
