package workloads

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	htmx "github.com/zeiss/fiber-htmx"
	links "github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/loading"
)

// WorkloadListController ...
type WorkloadListController struct {
	db        ports.Repository
	workloads []*models.Workload
	team      *authz.Team
	limit     int
	offset    int

	htmx.UnimplementedController
}

// NewWorkloadListController ...
func NewWorkloadListController(db ports.Repository) *WorkloadListController {
	return &WorkloadListController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadListController) Prepare() error {
	hx := w.Hx()

	team := hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)
	w.team = team

	w.offset = hx.Context().QueryInt("offset", 0)
	w.limit = hx.Context().QueryInt("limit", 10)

	workloads, err := w.db.ListWorkloads(hx.Context().Context(), team.Slug, &models.Pagination{Limit: w.limit, Offset: w.offset})
	if err != nil {
		return err
	}

	w.workloads = workloads

	return nil
}

// Get ...
func (w *WorkloadListController) Get() error {
	hx := w.Hx()

	workloadItems := make([]htmx.Node, len(w.workloads))
	for i, workload := range w.workloads {
		workloadItems[i] = htmx.Tr(
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
						Href: fmt.Sprintf("/%s/workloads/%s", w.team.Slug, workload.ID.String()),
					},
					htmx.Text(workload.Name),
				),
			),
			htmx.Td(
				htmx.Text(
					workload.Description,
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
						htmx.ClassNames{"overflow-x-auto": true},
						htmx.Div(
							htmx.ClassNames{
								"flex":            true,
								"justify-between": true,
								"items-center":    true,
							},
							htmx.Input(
								htmx.ClassNames{
									"input":          true,
									"input-bordered": true,
								},
								htmx.Attribute("type", "search"),
								htmx.Attribute("placeholder", "Search ..."),
								htmx.Attribute("name", "q"),
								htmx.HxPost("/workloads/search"),
								htmx.HxTarget("#data-table"),
								htmx.HxSwap("outerHTML"),
								htmx.HxIndicator(".htmx-indicator"),
								htmx.HxTrigger("keyup changed delay:500ms, search"),
							),
							htmx.Div(
								loading.Spinner(loading.SpinnerProps{
									ClassNames: htmx.ClassNames{
										"htmx-indicator": true,
									},
								}),
							),
						),
						htmx.Table(
							htmx.ClassNames{
								"table": true,
							},
							htmx.THead(
								htmx.Tr(
									htmx.Th(
										htmx.Label(
											htmx.Input(
												htmx.ClassNames{
													"checkbox": true,
												},
												htmx.Attribute("type", "checkbox"),
												htmx.Attribute("name", "all"),
											),
										),
									),
									htmx.Th(htmx.Text("ID")),
									htmx.Th(htmx.Text("Name")),
									htmx.Th(htmx.Text("Description")),
								),
							),
							htmx.TBody(
								htmx.ID("data-table"),
								htmx.Group(workloadItems...),
							),
						),
						htmx.Div(
							htmx.FormElement(
								htmx.ClassNames{},
								htmx.Select(
									htmx.HxTrigger("change"),
									htmx.HxTarget("html"),
									htmx.HxSwap("outerHTML"),
									htmx.HxGet(fmt.Sprintf("/workloads/list?limit=%d&offset=%d", w.limit, w.offset)),
									htmx.ClassNames{
										"select":   true,
										"max-w-xs": true,
									},
									htmx.Option(
										htmx.Text("10"),
										htmx.Attribute("value", "10"),
									),
									htmx.Option(
										htmx.Text("20"),
										htmx.Attribute("value", "20"),
									),
									htmx.Option(
										htmx.Text("30"),
										htmx.Attribute("value", "30"),
									),
								),
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":   true,
								"w-full": true,
							},
						),
					),
				),
			),
		),
	)
}
