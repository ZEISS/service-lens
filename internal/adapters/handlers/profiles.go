package handlers

import (
	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

type profilesHandler struct {
	pc ports.Profiles
}

// NewProfilesHandler returns a new ProfilesHandler.
func NewProfilesHandler(pc ports.Profiles) *profilesHandler {
	return &profilesHandler{pc}
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
					htmx.H1(htmx.Text(profile.Name)),
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

		hx.Redirect("/")

		return nil
	}
}

// New is the handler for the new page.
func (p *profilesHandler) New() fiber.Handler {
	return htmx.NewCompHandler(
		components.Page(
			components.PageProps{
				Children: []htmx.Node{
					htmx.FormElement(
						htmx.HxPost("/profiles"),
						htmx.LabElement(
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
						htmx.LabElement(
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
				},
			},
		),
	)
}
