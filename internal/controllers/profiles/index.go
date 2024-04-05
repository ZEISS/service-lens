package profiles

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// ProfileIndexControllerParams ...
type ProfileIndexControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultProfileIndexControllerParams ...
func NewDefaultProfileIndexControllerParams() *ProfileIndexControllerParams {
	return &ProfileIndexControllerParams{}
}

// ProfileIndexController ...
type ProfileIndexController struct {
	db      ports.Repository
	profile *models.Profile
	params  *ProfileIndexControllerParams

	htmx.UnimplementedController
}

// NewProfileIndexController ...
func NewProfileIndexController(db ports.Repository) *ProfileIndexController {
	return &ProfileIndexController{
		db: db,
	}
}

// Prepare ...
func (p *ProfileIndexController) Prepare() error {
	p.params = NewDefaultProfileIndexControllerParams()
	if err := p.Hx().Ctx().ParamsParser(p.params); err != nil {
		return err
	}

	profile, err := p.db.GetProfileByID(p.Hx().Ctx().Context(), p.params.ID)
	if err != nil {
		return err
	}
	p.profile = profile

	return nil
}

// Get ...
func (p *ProfileIndexController) Get() error {
	return p.Hx().RenderComp(
		components.Page(
			p.Hx(),
			components.PageProps{},
			components.Layout(
				p.Hx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Overview"),
							),
							htmx.Div(
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
										htmx.Text("ID"),
									),
									htmx.H3(
										htmx.Text(p.profile.ID.String()),
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
										htmx.Text("Name"),
									),
									htmx.H3(
										htmx.Text(p.profile.Name),
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
										htmx.Text("Description"),
									),
									htmx.H3(
										htmx.Text(p.profile.Description),
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
											p.profile.CreatedAt.Format("2006-01-02 15:04:05"),
										),
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
										htmx.Text("Updated at"),
									),
									htmx.H3(
										htmx.Text(
											p.profile.UpdatedAt.Format("2006-01-02 15:04:05"),
										),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								links.Button(
									links.LinkProps{
										ClassNames: htmx.ClassNames{
											"btn-outline": true,
											"btn-primary": true,
										},
										Href: fmt.Sprintf("%s/edit", p.params.ID),
									},
									htmx.Text("Edit"),
								),
								buttons.OutlinePrimary(
									buttons.ButtonProps{},
									htmx.HxDelete(""),
									htmx.HxConfirm("Are you sure you want to delete this profile?"),
									htmx.Text("Delete"),
								),
							),
						),
					),
				),
			),
		),
	)
}

// Delete ...
func (p *ProfileIndexController) Delete() error {
	err := p.db.DestroyProfile(p.Hx().Ctx().Context(), p.params.ID)
	if err != nil {
		return err
	}

	p.Hx().Redirect("list")

	return nil
}
