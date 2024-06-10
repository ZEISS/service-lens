package workloads

import (
	"fmt"

	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadLensController ...
type WorkloadLensController struct {
	db   ports.Repository
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
	// if err := w.BindValues(utils.User(w.db), utils.Team(w.db)); err != nil {
	// 	return err
	// }

	// lensID, err := uuid.Parse(w.Ctx().Params("lens"))
	// if err != nil {
	// 	return err
	// }

	// lens, err := w.db.GetLensByID(w.Context(), lensID)
	// if err != nil {
	// 	return err
	// }
	// w.lens = lens

	return nil
}

// Get ...
func (w *WorkloadLensController) Get() error {
	pillars := make([]htmx.Node, len(w.lens.Pillars))
	for _, pillar := range w.lens.Pillars {
		tr := htmx.Tr(
			htmx.Td(
				links.Link(
					links.LinkProps{
						Href: fmt.Sprintf("%s/pillars/%d", w.lens.ID, pillar.ID),
					},
					htmx.Text(pillar.Name),
				),
			),
		)

		pillars = append(pillars, tr)
	}

	return w.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: w.Path(),
				},
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
						htmx.Div(
							htmx.ClassNames{},
							links.Button(
								links.LinkProps{
									Href: fmt.Sprintf("%s/edit", w.lens.ID),
								},
								htmx.Text("Review"),
							),
						),
					),
				),
				components.Wrap(
					components.WrapProps{
						ClassNames: htmx.ClassNames{
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
									htmx.Th(htmx.Text("Pillar")),
								),
							),
							htmx.TBody(
								htmx.Group(pillars...),
							),
						),
					),
				),
			),
		),
	)
}
