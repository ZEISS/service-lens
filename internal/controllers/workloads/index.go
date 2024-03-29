package workloads

import (
	"fmt"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	links "github.com/zeiss/fiber-htmx/components/links"
)

// WorkloadIndexController ...
type WorkloadIndexController struct {
	db       ports.Repository
	workload *models.Workload

	htmx.UnimplementedController
}

// NewWorkloadIndexController ...
func NewWorkloadIndexController(db ports.Repository) *WorkloadIndexController {
	return &WorkloadIndexController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadIndexController) Prepare() error {
	hx := w.Hx()

	id, err := uuid.Parse(hx.Context().Params("id"))
	if err != nil {
		return err
	}

	workload, err := w.db.IndexWorkload(hx.Context().Context(), id)
	if err != nil {
		return err
	}
	w.workload = workload

	return nil
}

// Get ...
func (w *WorkloadIndexController) Get() error {
	hx := w.Hx()

	lenses := make([]htmx.Node, len(w.workload.Lenses))
	for i, lens := range w.workload.Lenses {
		lenses[i] = htmx.Tr(
			htmx.Th(htmx.Text(lens.ID.String())),
			htmx.Td(
				links.Link(
					links.LinkProps{
						Href: fmt.Sprintf("%s/lenses/%s", w.workload.ID, lens.ID.String()),
					},
					htmx.Text(lens.Name),
				),
			),
		)
	}

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
							htmx.Text(w.workload.Name),
						),
						htmx.P(
							htmx.Text(w.workload.Description),
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
									w.workload.CreatedAt.Format("2006-01-02 15:04:05"),
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
									w.workload.UpdatedAt.Format("2006-01-02 15:04:05"),
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
							htmx.TBody(
								htmx.Group(lenses...),
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
							htmx.Text(
								w.workload.Profile.Name,
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
							htmx.Text("Description"),
						),
						htmx.H3(
							htmx.Text(
								w.workload.Profile.Description,
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
								w.workload.Profile.UpdatedAt.Format("2006-01-02 15:04:05"),
							),
						),
					),
				),
			),
		),
	)
}

// Delete ...
func (w *WorkloadIndexController) Delete() error {
	hx := w.Hx()

	id, err := uuid.Parse(hx.Ctx().Params("id"))
	if err != nil {
		return err
	}

	err = w.db.DestroyWorkload(hx.Ctx().Context(), id)
	if err != nil {
		return err
	}

	hx.Redirect("/workloads/list")

	return nil
}
