package workloads

import (
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	links "github.com/zeiss/fiber-htmx/components/links"
)

// WorkloadSearchController ...
type WorkloadSearchController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewWorkloadSearchController ...
func NewWorkloadSearchController(db ports.Repository) *WorkloadSearchController {
	return &WorkloadSearchController{db, htmx.UnimplementedController{}}
}

// Post ...
func (w *WorkloadSearchController) Post() error {
	hx := w.Hx()

	q := hx.Ctx().FormValue("q")

	pagination := &models.Pagination{
		Limit:  10,
		Offset: 0,
		Search: q,
	}

	workloads, err := w.db.ListWorkloads(hx.Ctx().Context(), pagination)
	if err != nil {
		return err
	}

	workloadsItems := make([]htmx.Node, len(workloads))
	for i, workload := range workloads {
		workloadsItems[i] = htmx.Tr(
			htmx.Th(
				htmx.Label(
					htmx.Input(
						htmx.ClassNames{
							"checkbox": true,
						},
						htmx.Attribute("type", "checkbox"),
						htmx.Attribute("name", "profile"),
						htmx.Attribute("value", workload.ID.String()),
					),
				),
			),
			htmx.Th(htmx.Text(workload.ID.String())),
			htmx.Td(
				links.Link(
					links.LinkProps{
						ClassNames: htmx.ClassNames{
							"link": false,
						},
						Href: fmt.Sprintf("/workloads/%s", workload.ID.String()),
					},
					htmx.Text(workload.Name),
				),
			),
			htmx.Td(htmx.Text(workload.Description)),
		)
	}

	return hx.RenderComp(
		htmx.TBody(
			htmx.ID("data-table"),
			htmx.Group(workloadsItems...),
		),
	)
}
