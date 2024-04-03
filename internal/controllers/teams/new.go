package teams

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// TeamNewControllerParams ...
type TeamNewControllerParams struct {
	ID uuid.UUID `json:"id" xml:"id" form:"id"`
}

// NewDefaultTeamNewControllerParams ...
func NewDefaultTeamNewControllerParams() *TeamNewControllerParams {
	return &TeamNewControllerParams{}
}

// TeamNewControllerQuery ...
type TeamNewControllerQuery struct{}

// NewDefaultTeamNewControllerQuery ...
func NewDefaultTeamNewControllerQuery() *TeamNewControllerQuery {
	return &TeamNewControllerQuery{}
}

// TeamNewControllerBody ...
type TeamNewControllerBody struct {
	Name        string `json:"name" xml:"name" form:"name"`
	Slug        string `json:"slug" xml:"slug" form:"slug"`
	Description string `json:"description" xml:"description" form:"description"`
}

// NewDefaultTeamNewControllerBody ...
func NewDefaultTeamNewControllerBody() *TeamNewControllerBody {
	return &TeamNewControllerBody{}
}

// TeamNewController ...
type TeamNewController struct {
	db     ports.Repository
	params *TeamNewControllerParams
	query  *TeamNewControllerQuery
	body   *TeamNewControllerBody

	htmx.UnimplementedController
}

// Prepare ...
func (a *TeamNewController) Prepare() error {
	params := NewDefaultTeamNewControllerParams()
	if err := a.Hx().Context().ParamsParser(params); err != nil {
		return err
	}
	a.params = params

	query := NewDefaultTeamNewControllerQuery()
	if err := a.Hx().Context().QueryParser(query); err != nil {
		return err
	}
	a.query = query

	body := NewDefaultTeamNewControllerBody()
	if err := a.Hx().Context().BodyParser(body); err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
		return err
	}
	a.body = body

	return nil
}

// NewTeamNewController ...
func NewTeamNewController(db ports.Repository) *TeamNewController {
	return &TeamNewController{
		db: db,
	}
}

// Get ...
func (a *TeamNewController) Get() error {
	return a.Hx().RenderComp(
		components.Page(
			a.Hx(),
			components.PageProps{},
			components.Layout(
				a.Hx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.FormElement(
						htmx.HxPost(""),
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
								htmx.Span(
									htmx.ClassNames{
										"label-text": true,
									},
								),
							),
							forms.TextInputBordered(
								forms.TextInputProps{
									ClassNames: htmx.ClassNames{
										"w-full":   true,
										"max-w-xs": false,
									},
									Name:        "name",
									Placeholder: "Name ...",
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
								htmx.Span(
									htmx.ClassNames{
										"label-text": true,
									},
								),
							),
							forms.TextInputBordered(
								forms.TextInputProps{
									ClassNames: htmx.ClassNames{
										"w-full":   true,
										"max-w-lg": true,
									},
									Name:        "slug",
									Placeholder: "Slug ...",
								},
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
									htmx.Span(
										htmx.ClassNames{
											"label-text": true,
										},
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										ClassNames: htmx.ClassNames{
											"w-full":   true,
											"max-w-xs": false,
										},
										Name:        "description",
										Placeholder: "Description ...",
									},
								),
							),
						),
						buttons.OutlinePrimary(
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
		),
	)
}

// Post ...
func (a *TeamNewController) Post() error {
	team := &authz.Team{
		Name:        a.body.Name,
		Slug:        a.body.Slug,
		Description: utils.StrPtr(a.body.Description),
	}

	team, err := a.db.AddTeam(a.Hx().Ctx().Context(), team)
	if err != nil {
		return err
	}

	a.Hx().Redirect(fmt.Sprintf("/site/teams/%s", team.ID))

	return nil
}
