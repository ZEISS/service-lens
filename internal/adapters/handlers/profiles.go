package handlers

import (
	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

type profilesHandler struct {
	pc ports.Profiles
	db ports.Repository
}

// NewProfilesHandler returns a new ProfilesHandler.
func NewProfilesHandler(pc ports.Profiles, db ports.Repository) *profilesHandler {
	return &profilesHandler{pc, db}
}

// Index is the handler for the index page.
func (p *profilesHandler) Index() fiber.Handler {
	return htmx.NewCompHandler(
		components.Page(
			components.PageProps{
				Children: []htmx.Node{
					components.Table(components.TableProps{}),
				},
			},
		),
	)
}

// GetProfile is the handler for the get profile page.
func (p *profilesHandler) GetProfile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return err
		}

		profile, err := p.pc.FetchProfile(c.Context(), id)
		if err != nil {
			return err
		}

		page := components.Page(
			components.PageProps{
				Children: []htmx.Node{
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
								htmx.Attribute("value", profile.Name),
								htmx.Attribute("readonly", "true"),
								htmx.Attribute("disabled", "true"),
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
								htmx.Attribute("value", profile.Description),
								htmx.Attribute("readonly", "true"),
								htmx.Attribute("disabled", "true"),
								htmx.ClassNames{
									"input":          true,
									"input-bordered": true,
									"w-full":         true,
									"max-w-lg":       true,
								},
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"divider": true,
							},
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
								htmx.Text("Last updated"),
							),
							htmx.H3(
								htmx.Text(profile.UpdatedAt.Format("2006-01-02 15:04:05")),
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
								htmx.Text("Created at"),
							),
							htmx.H3(
								htmx.Text(
									profile.CreatedAt.Format("2006-01-02 15:04:05"),
								),
							),
						),
					),
				},
			},
		)

		c.Set("Content-Type", "text/html")

		return page.Render(c)
	}
}

// NewProfile is the handler for the new profile page.
func (p *profilesHandler) NewProfile() htmx.HtmxHandlerFunc {
	return func(hx *htmx.Htmx) error {
		profile := &models.Profile{
			Name:        hx.Ctx().FormValue("name"),
			Description: hx.Ctx().FormValue("description"),
		}

		err := p.pc.NewProfile(hx.Ctx().Context(), profile)
		if err != nil {
			return err
		}

		hx.Redirect("/profiles/" + profile.ID.String())

		return nil
	}
}

// New is the handler for the new page.
func (p *profilesHandler) New() fiber.Handler {
	return htmx.NewCompHandler(
		components.Page(
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

// List is the handler for the list page.
func (p *profilesHandler) List(c *fiber.Ctx) (htmx.Node, error) {
	profiles, err := p.pc.ListProfiles(c.Context(), &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return nil, err
	}

	items := make([]htmx.Node, len(profiles))
	for i, profile := range profiles {
		items[i] = htmx.Tr(
			htmx.Th(
				htmx.Label(
					htmx.Input(
						htmx.ClassNames{
							"checkbox": true,
						},
						htmx.Attribute("type", "checkbox"),
						htmx.Attribute("name", "profile"),
						htmx.Attribute("value", profile.ID.String()),
					),
				),
			),
			htmx.Th(htmx.Text(profile.ID.String())),
			htmx.Td(htmx.Text(profile.Name)),
			htmx.Td(htmx.Text(profile.Description)),
		)
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
							Href:  "/profiles/list",
							Title: "Profiles",
						},
					),
				),
			),
			components.SubNavActions(
				components.SubNavActionsProps{},
				links.Link(
					links.LinkProps{
						Href: "/profiles/new",
						ClassNames: htmx.ClassNames{
							"btn":         true,
							"btn-outline": true,
							"btn-xs":      true,
							"link-hover":  true,
						},
					},
					htmx.Text("Create Profile"),
				),
			),
		),
		htmx.Div(
			htmx.ClassNames{"overflow-x-auto": true},
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
					items...,
				),
			),
			htmx.Div(
				htmx.ClassNames{},
				htmx.Select(
					htmx.ClassNames{
						"select":   true,
						"max-w-xs": true,
					},
					htmx.Option(
						htmx.Text("10"),
						htmx.Attribute("value", "10"),
					),
					htmx.Option(
						htmx.Text("20"),
						htmx.Attribute("value", "20"),
					),
					htmx.Option(
						htmx.Text("30"),
						htmx.Attribute("value", "30"),
					),
				),
			),
		),
	), nil
}
