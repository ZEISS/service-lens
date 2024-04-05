package profiles

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

// ProfileEditControllerParams ...
type ProfileEditControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultProfileEditControllerParams ...
func NewDefaultProfileEditControllerParams() *ProfileEditControllerParams {
	return &ProfileEditControllerParams{}
}

// ProfileEditController ...
type ProfileEditController struct {
	db      ports.Repository
	params  *ProfileEditControllerParams
	profile *models.Profile
	team    *authz.Team

	htmx.UnimplementedController
}

// NewProfileEditController ...
func NewProfileEditController(db ports.Repository) *ProfileEditController {
	return &ProfileEditController{
		db: db,
	}
}

// Prepare ...
func (p *ProfileEditController) Prepare() error {
	p.params = NewDefaultProfileEditControllerParams()
	if err := p.Hx().Ctx().ParamsParser(p.params); err != nil {
		return err
	}

	profile, err := p.db.GetProfileByID(p.Hx().Ctx().Context(), p.params.ID)
	if err != nil {
		return err
	}
	p.profile = profile
	p.team = p.Hx().Values(resolvers.ValuesKeyTeam).(*authz.Team)

	return nil
}

// Post ...
func (p *ProfileEditController) Post() error {
	query := NewDefaultProfileNewControllerQuery()
	if err := p.Hx().Ctx().BodyParser(query); err != nil {
		return err
	}

	p.profile.Description = query.Description

	err := p.db.UpdateProfile(p.Hx().Ctx().Context(), p.profile)
	if err != nil {
		return err
	}

	p.Hx().Redirect(fmt.Sprintf("/%s/profiles/%s", p.team.Slug, p.profile.ID))

	return nil
}

// New ...
func (p *ProfileEditController) Get() error {
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
										Value:    p.profile.Name,
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
											Value: p.profile.Description,
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
						htmx.Text("Upload Profile"),
					),
				),
			),
		),
	)
}