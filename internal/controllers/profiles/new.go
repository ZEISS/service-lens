package profiles

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"
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

	team := hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)

	profile := &models.Profile{
		Name:        hx.Ctx().FormValue("name"),
		Description: hx.Ctx().FormValue("description"),
		Team:        *team,
	}

	err := p.db.NewProfile(hx.Ctx().Context(), profile)
	if err != nil {
		return err
	}

	hx.Redirect(fmt.Sprintf("/%s/profiles/%s", team.Slug, profile.ID))

	return nil
}

// New ...
func (p *ProfileNewController) Get() error {
	team := p.Hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)

	return p.Hx.RenderComp(
		components.Page(
			p.Hx,
			components.PageProps{},
			components.Layout(
				p.Hx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.FormElement(
						htmx.HxPost(fmt.Sprintf("/%s/profiles/new", team.Slug)),
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
									htmx.Text("Name"),
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
		),
	)
}
