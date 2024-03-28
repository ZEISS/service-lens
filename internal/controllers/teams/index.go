package teams

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// TeamIndexController ...
type TeamIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewTeamsIndexController ...
func NewTeamsIndexController(db ports.Repository) *TeamIndexController {
	return &TeamIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (a *TeamIndexController) Get() error {
	hx := a.Hx()

	id, err := uuid.Parse(a.Hx().Context().Params("id"))
	if err != nil {
		return err
	}

	team, err := a.db.GetTeamByID(hx.Context().Context(), id)
	if err != nil {
		return err
	}

	return hx.RenderComp(
		components.Page(
			a.Hx(),
			components.PageProps{},
			components.SubNav(
				components.SubNavProps{},
				components.SubNavBreadcrumb(
					components.SubNavBreadcrumbProps{},
					breadcrumbs.Breadcrumbs(
						breadcrumbs.BreadcrumbsProps{},
						breadcrumbs.Breadcrumb(
							breadcrumbs.BreadcrumbProps{
								Href:  "/",
								Title: "Home",
							},
						),
						breadcrumbs.Breadcrumb(
							breadcrumbs.BreadcrumbProps{
								Href:  "/teams/list",
								Title: "Teams",
							},
						),
						breadcrumbs.Breadcrumb(
							breadcrumbs.BreadcrumbProps{
								Href:  "/teams/" + team.ID.String(),
								Title: team.Name,
							},
						),
					),
				),
			),
			components.Wrap(
				components.WrapProps{},
				htmx.Div(
					htmx.H1(
						htmx.Text(team.Name),
					),
					htmx.P(
						htmx.Text(utils.PtrStr(team.Description)),
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
								team.CreatedAt.Format("2006-01-02 15:04:05"),
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
								team.UpdatedAt.Format("2006-01-02 15:04:05"),
							),
						),
					),
				),
			),
		),
	)
}
