package environments

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// EnvironmentNewController ...
type EnvironmentNewController struct {
	db ports.Repository

	htmx.DefaultController
}

// NewEnvironmentNewController ...
func NewEnvironmentNewController(db ports.Repository) *EnvironmentNewController {
	return &EnvironmentNewController{
		db: db,
	}
}

// EnvironmentNewControllerQuery ...
type EnvironmentNewControllerQuery struct {
	Name        string `json:"name" xml:"name" form:"name"`
	Description string `json:"description" xml:"description" form:"description"`
}

// NewDefaultEnvironmentNewControllerQuery ...
func NewDefaultEnvironmentNewControllerQuery() *EnvironmentNewControllerQuery {
	return &EnvironmentNewControllerQuery{}
}

// Prepare ...
func (p *EnvironmentNewController) Prepare() error {
	if err := p.BindValues(utils.User(p.db), utils.Team(p.db)); err != nil {
		return err
	}

	return nil
}

// Post ...
func (p *EnvironmentNewController) Post() error {
	team := p.Values(utils.ValuesKeyTeam).(*authz.Team)

	query := NewDefaultEnvironmentNewControllerQuery()
	if err := p.BindBody(query); err != nil {
		return err
	}

	Environment := &models.Environment{
		Name:        query.Name,
		Description: query.Description,
		Team:        *team,
	}

	err := p.db.NewEnvironment(p.Context(), Environment)
	if err != nil {
		return err
	}

	p.Hx().Redirect(fmt.Sprintf("/teams/%s/environments/%s", team.Slug, Environment.ID))

	return nil
}

// Error ...
func (p *EnvironmentNewController) Error(err error) error {
	fmt.Println(err)
	return nil
}

// New ...
func (p *EnvironmentNewController) Get() error {
	return p.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: p.Values(utils.ValuesKeyUser).(*authz.User),
					Team: p.Values(utils.ValuesKeyTeam).(*authz.Team),
				},
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
										htmx.Text("A unique identifier for the environment."),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name: "name",
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
											htmx.Text("A brief description of the environment that workloads run in."),
										),
									),
									forms.TextareaBordered(
										forms.TextareaProps{
											Name: "description",
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
								htmx.Text("Tags - Optional"),
							),
						),
					),
					buttons.OutlinePrimary(
						buttons.ButtonProps{},
						htmx.Attribute("type", "submit"),
						htmx.Text("Create Environment"),
					),
				),
			),
		),
	)
}
