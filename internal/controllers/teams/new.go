package teams

import (
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// TeamsNewController ...
type TeamsNewController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewTeamsNewController ...
func NewTeamsNewController(db ports.Repository) *TeamsNewController {
	return &TeamsNewController{db, htmx.UnimplementedController{}}
}

// Get ...
func (a *TeamsNewController) Get() error {
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
		),
	)
}

// Post ...
func (a *TeamsNewController) Post() error {
	hx := a.Hx

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
