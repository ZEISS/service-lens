package teams

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"
	"github.com/zeiss/service-lens/internal/utils"
)

// TeamNewController ...
type TeamNewController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewTeamNewController ...
func NewTeamNewController(db ports.Repository) *TeamNewController {
	return &TeamNewController{
		db: db,
	}
}

// TeamNewControllerQuery ...
type TeamNewControllerQuery struct {
	Name        string `json:"name" xml:"name" form:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" xml:"description" form:"description" validate:"omitempty,min=3,max=1024"`
	Slug        string `json:"slug" xml:"slug" form:"slug" validate:"required,min=3,max=100,lowercase"`
}

// NewDefaultTeamNewControllerQuery ...
func NewDefaultTeamNewControllerQuery() *TeamNewControllerQuery {
	return &TeamNewControllerQuery{}
}

// Post ...
func (p *TeamNewController) Post() error {
	hx := p.Hx()

	query := NewDefaultTeamNewControllerQuery()
	if err := hx.Ctx().BodyParser(query); err != nil {
		return err
	}

	user := hx.Values(resolvers.ValuesKeyUser).(*authz.User)

	team := &authz.Team{
		Name:        query.Name,
		Description: utils.StrPtr(query.Description),
		Slug:        query.Slug,
		Users:       &[]authz.User{*user},
	}

	validator := validator.New()
	if err := validator.Struct(query); err != nil {
		return err
	}

	team, err := p.db.CreateTeam(hx.Ctx().Context(), team, user)
	if err != nil {
		return err
	}

	hx.Redirect(fmt.Sprintf("/site/teams/%s", team.ID))

	return nil
}

// Get ...
func (p *TeamNewController) Get() error {
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
										htmx.Text("The display name of the team."),
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
											htmx.Text("Slug"),
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
											htmx.Text("A unique identifier for the team."),
										),
									),
									forms.TextInputBordered(
										forms.TextInputProps{
											Name: "slug",
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
											htmx.Text("The slug must be from 3 to 100 characters. At least 3 characters must be non-whitespace. The slug must be lowercase and contain only letters, numbers, and hyphens. The slug must be unique within the site"),
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
												htmx.Text("A brief description of the team to document its scope and intended purpose."),
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
					),
					buttons.OutlinePrimary(
						buttons.ButtonProps{},
						htmx.Attribute("type", "submit"),
						htmx.Text("Create Team"),
					),
				),
			),
		),
	)
}
