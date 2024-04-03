package teams

import (
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
)

// TeamIndexControllerParams ...
type TeamIndexControllerParams struct {
	ID uuid.UUID `json:"id" xml:"id" form:"id"`
}

// NewDefaultTeamIndexControllerParams ...
func NewDefaultTeamIndexControllerParams() *TeamIndexControllerParams {
	return &TeamIndexControllerParams{}
}

// TeamIndexControllerQuery ...
type TeamIndexControllerQuery struct{}

// NewDefaultTeamIndexControllerQuery ...
func NewDefaultTeamIndexControllerQuery() *TeamIndexControllerQuery {
	return &TeamIndexControllerQuery{}
}

// TeamIndexController ...
type TeamIndexController struct {
	db     ports.Repository
	params *TeamIndexControllerParams
	query  *TeamIndexControllerQuery
	team   *authz.Team

	htmx.UnimplementedController
}

// NewTeamIndexController ...
func NewTeamIndexController(db ports.Repository) *TeamIndexController {
	return &TeamIndexController{
		db: db,
	}
}

// Prepare ...
func (w *TeamIndexController) Prepare() error {
	hx := w.Hx()

	params := NewDefaultTeamIndexControllerParams()
	if err := hx.Context().ParamsParser(params); err != nil {
		return err
	}
	w.params = params

	query := NewDefaultTeamIndexControllerQuery()
	if err := hx.Context().QueryParser(query); err != nil {
		return err
	}
	w.query = query

	team, err := w.db.GetTeamByID(hx.Context().Context(), w.params.ID)
	if err != nil {
		return err
	}
	w.team = team

	return nil
}

// Get ...
func (w *TeamIndexController) Get() error {
	hx := w.Hx()

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.H1(
							htmx.Text(w.team.Name),
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
								htmx.Text(
									utils.PtrStr(w.team.Description),
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
								htmx.Text("Created at"),
							),
							htmx.H3(
								htmx.Text(
									w.team.CreatedAt.Format("2006-01-02 15:04:05"),
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
									w.team.UpdatedAt.Format("2006-01-02 15:04:05"),
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
func (w *TeamIndexController) Delete() error {
	return nil
}
