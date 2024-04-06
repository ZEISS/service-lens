package teams

import (
	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// TeamIndexControllerParams ...
type TeamIndexControllerParams struct {
	ID uuid.UUID `json:"id" xml:"id" form:"id" validate:"required,uuid"`
}

// NewDefaultTeamIndexControllerParams ...
func NewDefaultTeamIndexControllerParams() *TeamIndexControllerParams {
	return &TeamIndexControllerParams{}
}

// TeamIndexController ...
type TeamIndexController struct {
	db     ports.Repository
	team   *authz.Team
	params *TeamIndexControllerParams

	htmx.UnimplementedController
}

// NewTeamIndexController ...
func NewTeamIndexController(db ports.Repository) *TeamIndexController {
	return &TeamIndexController{
		db: db,
	}
}

// Prepare ...
func (l *TeamIndexController) Prepare() error {
	l.params = NewDefaultTeamIndexControllerParams()
	if err := l.Hx().Ctx().ParamsParser(l.params); err != nil {
		return err
	}

	team, err := l.db.GetTeamByID(l.Hx().Ctx().Context(), l.params.ID)
	if err != nil {
		return err
	}
	l.team = team

	return nil
}

// Get ...
func (l *TeamIndexController) Get() error {
	return l.Hx().RenderComp(
		components.Page(
			l.Hx(),
			components.PageProps{},
			components.Layout(
				l.Hx(),
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
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "ID",
										Data:  l.team.ID.String(),
									},
								),
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "Name",
										Data:  l.team.Name,
									},
								),
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "Slug",
										Data:  l.team.Slug,
									},
								),
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "Description",
										Data:  utils.PtrStr(l.team.Description),
									},
								),
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "Created at",
										Data:  l.team.CreatedAt.Format("2006-01-02 15:04:05"),
									},
								),
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "Updated at",
										Data:  l.team.UpdatedAt.Format("2006-01-02 15:04:05"),
									},
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.OutlinePrimary(
									buttons.ButtonProps{},
									htmx.HxDelete(""),
									htmx.HxConfirm("Are you sure you want to delete this team?"),
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
func (l *TeamIndexController) Delete() error {
	err := l.db.DeleteTeam(l.Hx().Ctx().Context(), l.params.ID)
	if err != nil {
		return err
	}

	l.Hx().Redirect("/site/teams")

	return nil
}
