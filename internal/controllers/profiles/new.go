package profiles

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

// ProfileNewController ...
type ProfileNewController struct {
	db  ports.Repository
	ctx htmx.Ctx

	htmx.DefaultController
}

// NewProfileNewController ...
func NewProfileNewController(db ports.Repository) *ProfileNewController {
	return &ProfileNewController{
		db: db,
	}
}

// ProfileNewControllerQuery ...
type ProfileNewControllerQuery struct {
	Name        string `json:"name" xml:"name" form:"name"`
	Description string `json:"description" xml:"description" form:"description"`
}

// NewDefaultProfileNewControllerQuery ...
func NewDefaultProfileNewControllerQuery() *ProfileNewControllerQuery {
	return &ProfileNewControllerQuery{}
}

// Prepare ...
func (p *ProfileNewController) Prepare() error {
	ctx, err := htmx.NewDefaultContext(p.Hx().Ctx(), utils.Team(p.Hx().Ctx(), p.db), utils.User(p.Hx().Ctx(), p.db))
	if err != nil {
		return err
	}
	p.ctx = ctx

	return nil
}

// Post ...
func (p *ProfileNewController) Post() error {
	hx := p.Hx()

	team := htmx.Locals[*authz.Team](p.ctx, utils.ValuesKeyTeam)

	query := NewDefaultProfileNewControllerQuery()
	if err := hx.Ctx().BodyParser(query); err != nil {
		return err
	}

	profile := &models.Profile{
		Name:        query.Name,
		Description: query.Description,
		Team:        *team,
	}

	err := p.db.NewProfile(hx.Ctx().Context(), profile)
	if err != nil {
		return err
	}

	hx.Redirect(fmt.Sprintf("/teams/%s/profiles/%s", team.Slug, profile.ID))

	return nil
}

// New ...
func (p *ProfileNewController) Get() error {
	return p.Hx().RenderComp(
		components.Page(
			p.ctx,
			components.PageProps{},
			components.Layout(
				p.ctx,
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
											htmx.Text("A brief description of the workload to document its scope and intended purpose."),
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
								htmx.Text("Evaluation"),
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
						htmx.Text("Create Profile"),
					),
				),
			),
		),
	)
}
