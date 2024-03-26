package profiles

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// ProfileNewController ...
type ProfileNewController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewProfileNewController ...
func NewProfileNewController(db ports.Repository) *ProfileNewController {
	return &ProfileNewController{db, htmx.UnimplementedController{}}
}

// Post ...
func (p *ProfileNewController) Post() error {
	hx := p.Hx

	profile := &models.Profile{
		Name:        hx.Ctx().FormValue("name"),
		Description: hx.Ctx().FormValue("description"),
	}

	err := p.db.NewProfile(hx.Ctx().Context(), profile)
	if err != nil {
		return err
	}

	hx.Redirect("/profiles/" + profile.ID.String())

	return nil
}

// New ...
func (p *ProfileNewController) Get() error {
	return p.Hx.RenderComp(
		components.Page(
			p.Hx,
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
								Href:  "/profiles/list",
								Title: "Profiles",
							},
						),
					),
				),
			),
			components.Wrap(
				components.WrapProps{},
				htmx.FormElement(
					htmx.HxPost("/profiles"),
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
								htmx.Text("What is your name?"),
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
					htmx.Button(
						htmx.ClassNames{
							"btn":         true,
							"btn-default": true,
							"my-4":        true,
						},
						htmx.Attribute("type", "submit"),
						htmx.Text("Create Profile"),
					),
				),
			),
		),
	)
}
