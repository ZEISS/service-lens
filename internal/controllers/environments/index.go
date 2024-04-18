package environments

import (
	"fmt"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// EnvironmentIndexControllerParams ...
type EnvironmentIndexControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultEnvironmentIndexControllerParams ...
func NewDefaultEnvironmentIndexControllerParams() *EnvironmentIndexControllerParams {
	return &EnvironmentIndexControllerParams{}
}

// EnvironmentIndexController ...
type EnvironmentIndexController struct {
	db          ports.Repository
	environment *models.Environment
	params      *EnvironmentIndexControllerParams

	htmx.UnimplementedController
}

// NewEnvironmentIndexController ...
func NewEnvironmentIndexController(db ports.Repository) *EnvironmentIndexController {
	return &EnvironmentIndexController{
		db: db,
	}
}

// Prepare ...
func (p *EnvironmentIndexController) Prepare() error {
	p.params = NewDefaultEnvironmentIndexControllerParams()
	if err := p.BindParams(p.params); err != nil {
		return err
	}

	environment, err := p.db.GetEnvironment(p.Context(), p.params.ID)
	if err != nil {
		return err
	}
	p.environment = environment

	if err := p.BindValues(utils.User(p.db), utils.Team(p.db)); err != nil {
		return err
	}

	return nil
}

// Get ...
func (p *EnvironmentIndexController) Get() error {
	return p.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: p.Values(utils.ValuesKeyUser).(*authz.User),
					Team: p.Values(utils.ValuesKeyTeam).(*authz.Team),
				},
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
										htmx.Text(p.environment.ID.String()),
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
										htmx.Text(p.environment.Name),
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
										htmx.Text(p.environment.Description),
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
											p.environment.CreatedAt.Format("2006-01-02 15:04:05"),
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
											p.environment.UpdatedAt.Format("2006-01-02 15:04:05"),
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
									htmx.HxConfirm("Are you sure you want to delete this Environment?"),
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
func (p *EnvironmentIndexController) Delete() error {
	err := p.db.DeleteEnvironment(p.Context(), p.params.ID)
	if err != nil {
		return err
	}

	p.Hx().Redirect("list")

	return nil
}
