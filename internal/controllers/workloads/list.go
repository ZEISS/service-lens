package workloads

import (
	"fmt"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	links "github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/loading"
)

// WorkloadListController ...
type WorkloadListController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewWorkloadListController ...
func NewWorkloadListController(db ports.Repository) *WorkloadListController {
	return &WorkloadListController{db, htmx.UnimplementedController{}}
}

// Get ...
func (w *WorkloadListController) Get() error {
	hx := w.Hx

	offset := hx.Context().QueryInt("offset", 0)
	limit := hx.Context().QueryInt("limit", 10)

	profiles, err := w.db.ListWorkloads(hx.Context().Context(), &models.Pagination{Limit: limit, Offset: offset})
	if err != nil {
		return err
	}

	profilesItems := make([]htmx.Node, len(profiles))
	for i, profile := range profiles {
		profilesItems[i] = htmx.Tr(
			htmx.Th(
				htmx.Label(
					htmx.Input(
						htmx.ClassNames{
							"checkbox": true,
						},
						htmx.Attribute("type", "checkbox"),
						htmx.Attribute("name", "profile"),
						htmx.Attribute("value", profile.ID.String()),
					),
				),
			),
			htmx.Th(htmx.Text(profile.ID.String())),
			htmx.Td(
				links.Link(
					links.LinkProps{
						Href: fmt.Sprintf("/workloads/%s", profile.ID.String()),
					},
					htmx.Text(profile.Name),
				),
			),
			htmx.Td(
				htmx.Text(
					profile.Description,
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
									Href:  "/workloads/list",
									Title: "Workloads",
								},
							),
						),
					),
					components.SubNavActions(
						components.SubNavActionsProps{},
						links.Link(
							links.LinkProps{
								Href: "/workloads/new",
								ClassNames: htmx.ClassNames{
									"btn":         true,
									"btn-outline": true,
									"btn-xs":      true,
									"link-hover":  true,
								},
							},
							htmx.Text("Create Workload"),
						),
					),
				),
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
								htmx.Group(profilesItems...),
							),
						),
						htmx.Div(
							htmx.FormElement(
								htmx.ClassNames{},
								htmx.Select(
									htmx.HxTrigger("change"),
									htmx.HxTarget("html"),
									htmx.HxSwap("outerHTML"),
									htmx.HxGet(fmt.Sprintf("/workloads/list?limit=%d&offset=%d", limit, offset)),
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
