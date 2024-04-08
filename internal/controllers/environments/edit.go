package environments

import (
	"fmt"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"
)

// EnvironmentEditControllerParams ...
type EnvironmentEditControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultEnvironmentEditControllerParams ...
func NewDefaultEnvironmentEditControllerParams() *EnvironmentEditControllerParams {
	return &EnvironmentEditControllerParams{}
}

// EnvironmentEditController ...
type EnvironmentEditController struct {
	db          ports.Repository
	params      *EnvironmentEditControllerParams
	Environment *models.Environment
	team        *authz.Team

	htmx.UnimplementedController
}

// NewEnvironmentEditController ...
func NewEnvironmentEditController(db ports.Repository) *EnvironmentEditController {
	return &EnvironmentEditController{
		db: db,
	}
}

// Prepare ...
func (p *EnvironmentEditController) Prepare() error {
	p.params = NewDefaultEnvironmentEditControllerParams()
	if err := p.Hx().Ctx().ParamsParser(p.params); err != nil {
		return err
	}

	Environment, err := p.db.GetEnvironment(p.Hx().Ctx().Context(), p.params.ID)
	if err != nil {
		return err
	}
	p.Environment = Environment
	p.team = p.Hx().Values(resolvers.ValuesKeyTeam).(*authz.Team)

	return nil
}

// Post ...
func (p *EnvironmentEditController) Post() error {
	query := NewDefaultEnvironmentNewControllerQuery()
	if err := p.Hx().Ctx().BodyParser(query); err != nil {
		return err
	}

	p.Environment.Description = query.Description

	err := p.db.UpdateEnvironment(p.Hx().Ctx().Context(), p.Environment)
	if err != nil {
		return err
	}

	p.Hx().Redirect(fmt.Sprintf("/%s/environments/%s", p.team.Slug, p.Environment.ID))

	return nil
}

// New ...
func (p *EnvironmentEditController) Get() error {
	return p.Hx().RenderComp(
		components.Page(
			p.Hx(),
			components.PageProps{},
			components.Layout(
				p.Hx(),
				components.LayoutProps{},
				htmx.FormElement(
					htmx.HxPost(""),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"w-full": true,
								"my-4":   true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Properties"),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{
										"py-4": true,
									},
								},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"-my-4": true,
											},
										},
										htmx.Text("Name"),
									),
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("A unique identifier for the workload."),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:     "name",
										Value:    p.Environment.Name,
										Disabled: true,
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("The name must be from 3 to 100 characters. At least 3 characters must be non-whitespace."),
									),
								),
								forms.FormControl(
									forms.FormControlProps{
										ClassNames: htmx.ClassNames{
											"py-4": true,
										},
									},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"-my-4": true,
												},
											},
											htmx.Text("Description"),
										),
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
												},
											},
											htmx.Text("A brief description of the workload to document its scope and intended purpose."),
										),
									),
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:  "description",
											Value: p.Environment.Description,
										},
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
												},
											},
											htmx.Text("The description must be from 3 to 1024 characters."),
										),
									),
								),
							),
						),
					),
					buttons.OutlinePrimary(
						buttons.ButtonProps{},
						htmx.Attribute("type", "submit"),
						htmx.Text("Update Environment"),
					),
				),
			),
		),
	)
}
