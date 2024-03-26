package profiles

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// ProfileIndexController ...
type ProfileIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewProfileIndexController ...
func NewProfileIndexController(db ports.Repository) *ProfileIndexController {
	return &ProfileIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (p *ProfileIndexController) Get() error {
	hx := p.Hx

	id, err := uuid.Parse(hx.Context().Params("id"))
	if err != nil {
		return err
	}

	profile, err := p.db.FetchProfile(hx.Context().Context(), id)
	if err != nil {
		return err
	}

	return hx.RenderComp(
		components.Page(
			hx,
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
		),
	)
}
