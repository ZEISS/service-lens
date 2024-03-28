package workloads

import (
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadLensController ...
type WorkloadLensController struct {
	db   ports.Repository
	team *authz.Team
	lens *models.Lens

	htmx.UnimplementedController
}

// NewWorkloadLensController ...
func NewWorkloadLensController(db ports.Repository) *WorkloadLensController {
	return &WorkloadLensController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadLensController) Prepare() error {
	hx := w.Hx()

	team := hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)
	w.team = team

	lensID, err := uuid.Parse(hx.Context().Params("lens"))
	if err != nil {
		return err
	}

	lens, err := w.db.GetLensByID(hx.Context().Context(), team.Slug, lensID)
	if err != nil {
		return err
	}
	w.lens = lens

	return nil
}

// Get ...
func (w *WorkloadLensController) Get() error {
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
							htmx.Text(w.lens.Name),
						),
						htmx.P(
							htmx.Text(w.lens.Description),
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
									w.lens.CreatedAt.Format("2006-01-02 15:04:05"),
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
									w.lens.UpdatedAt.Format("2006-01-02 15:04:05"),
								),
							),
						),
					),
				),
				components.Wrap(
					components.WrapProps{
						ClassName: htmx.ClassNames{
							"border-neutral": true,
							"border-t":       true,
							"px-6":           true,
						},
					},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						htmx.Table(
							htmx.ClassNames{
								"table": true,
							},
							htmx.THead(
								htmx.Tr(
									htmx.Th(htmx.Text("ID")),
									htmx.Th(htmx.Text("Lens")),
								),
							),
							// htmx.TBody(
							// 	htmx.Group(lenses...),
							// ),
						),
					),
				),
			),
		),
	)
}
