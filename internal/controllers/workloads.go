package controllers

import (
	"fmt"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	links "github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/loading"
)

// Workloads ...
type Workloads struct {
	db ports.Repository
}

// NewWorkloadsController ...
func NewWorkloadsController(db ports.Repository) *Workloads {
	return &Workloads{db}
}

// Search ...
func (w *Workloads) Search(hx *htmx.Htmx) error {
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

// Store ...
func (w *Workloads) Store(hx *htmx.Htmx) error {
	profileId, err := uuid.Parse(hx.Ctx().FormValue("profile"))
	if err != nil {
		return err
	}

	lensId, err := uuid.Parse(hx.Ctx().FormValue("lens"))
	if err != nil {
		return err
	}

	workload := &models.Workload{
		Name:        hx.Ctx().FormValue("name"),
		Description: hx.Ctx().FormValue("description"),
		ProfileID:   profileId,
		Lenses:      []*models.Lens{{ID: lensId}},
	}

	err = w.db.StoreWorkload(hx.Ctx().Context(), workload)
	if err != nil {
		return err
	}

	hx.Redirect("/workloads/" + workload.ID.String())

	return nil
}

// List ...
func (w *Workloads) List(c *fiber.Ctx) (htmx.Node, error) {
	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", 10)

	profiles, err := w.db.ListWorkloads(c.Context(), &models.Pagination{Limit: limit, Offset: offset})
	if err != nil {
		return nil, err
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
			htmx.Td(htmx.Text(profile.Description)),
		)
	}

	return components.Page(
		components.PageProps{}.WithContext(c),
		components.SubNav(
			components.SubNavProps{},
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
		// components.SubNav(
		// 	components.SubNavProps{},

		// 	htmx.A(
		// 		htmx.ClassNames{
		// 			"btn": true,
		// 		},
		// 		htmx.Attribute("href", "/workloads/new"),
		// 		htmx.Text("Create Workload"),
		// 	),
		// ),
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
					htmx.ClassNames{"table": true},
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
	), nil
}

// New ...
func (w *Workloads) New(c *fiber.Ctx) (htmx.Node, error) {
	profiles, err := w.db.ListProfiles(c.Context(), &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return nil, err
	}

	profilesItems := make([]htmx.Node, len(profiles))
	for i, profile := range profiles {
		profilesItems[i] = htmx.Option(
			htmx.Attribute("value", profile.ID.String()),
			htmx.Text(profile.Name),
		)
	}

	lenses, err := w.db.ListLenses(c.Context(), &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return nil, err
	}

	lensesItems := make([]htmx.Node, len(lenses))
	for i, lens := range lenses {
		lensesItems[i] = htmx.Option(
			htmx.Attribute("value", lens.ID.String()),
			htmx.Text(lens.Name),
		)
	}

	return components.Page(
		components.PageProps{}.WithContext(c),
		htmx.FormElement(
			htmx.HxPost("/workloads/new"),
			htmx.Label(
				htmx.ClassNames{
					"form-control": true,
					"w-full":       true,
					"max-w-lg":     true,
					"mb-4":         true,
				},
				htmx.Div(
					htmx.ClassNames{
						"label": true,
					},
					htmx.Span(
						htmx.ClassNames{
							"label-text": true,
						},
					),
				),
				htmx.Input(
					htmx.Attribute("type", "text"),
					htmx.Attribute("name", "name"),
					htmx.Attribute("placeholder", "Name ..."),
					htmx.ClassNames{
						"input":          true,
						"input-bordered": true,
						"w-full":         true,
						"max-w-lg":       true,
					},
				),
			),
			htmx.Label(
				htmx.ClassNames{
					"form-control": true,
					"w-full":       true,
					"max-w-lg":     true,
				},
				htmx.Div(
					htmx.ClassNames{
						"label":   true,
						"sr-only": true,
					},
				),
				htmx.Input(
					htmx.Attribute("type", "text"),
					htmx.Attribute("name", "description"),
					htmx.Attribute("placeholder", "Description ..."),
					htmx.ClassNames{
						"input":          true,
						"input-bordered": true,
						"w-full":         true,
						"max-w-lg":       true,
					},
				),
			),
			htmx.Select(
				htmx.ClassNames{
					"select":   true,
					"max-w-xs": true,
					"block":    true,
				},
				htmx.Attribute("name", "profile"),
				htmx.Group(profilesItems...),
			),
			htmx.Select(
				htmx.ClassNames{
					"select":   true,
					"max-w-xs": true,
					"block":    true,
				},
				htmx.Attribute("name", "lens"),
				htmx.Group(lensesItems...),
			),
			htmx.Button(
				htmx.ClassNames{
					"btn":         true,
					"btn-default": true,
					"my-4":        true,
				},
				htmx.Attribute("type", "submit"),
				htmx.Text("Create Workload"),
			),
		),
	), nil
}

// Show ...
func (w *Workloads) Show(c *fiber.Ctx) (htmx.Node, error) {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return nil, err
	}

	workload, err := w.db.ShowWorkload(c.Context(), id)
	if err != nil {
		return nil, err
	}

	return components.Page(
		components.PageProps{}.WithContext(c),
		components.SubNav(
			components.SubNavProps{},
			htmx.Button(
				htmx.ClassNames{
					"btn":       true,
					"btn-ghost": true,
				},
				// htmx.HxTrigger("click"),
				// htmx.Target("#htmx-modal"),
				htmx.HxOn("click", "daisy_modal.showModal()"),
				htmx.Text("Delete"),
			),
			htmx.Dialog(
				htmx.ClassNames{
					"modal": true,
				},
				htmx.ID("daisy_modal"),
				htmx.Div(
					htmx.ClassNames{
						"modal-box": true,
					},
					htmx.H3(
						htmx.ClassNames{
							"font-bold": true,
							"text-lg":   true,
						},
						htmx.ID("htmx-modal"),
						htmx.Text("Confirm to delete the workload"),
					),
					htmx.Div(
						htmx.ClassNames{
							"modal-action": true,
						},
						htmx.FormElement(
							htmx.HxDelete("/workloads/"+workload.ID.String()),
							htmx.Method("dialog"),
							htmx.Button(
								htmx.ClassNames{
									"btn": true,
								},
								htmx.Attribute("type", "submit"),
								htmx.Text("Confirm"),
							),
						),
					),
				),
			),
		),
		htmx.Div(
			htmx.ClassNames{"p-4": true},
			htmx.H1(
				htmx.Text(workload.Name),
			),
			htmx.P(
				htmx.Text(workload.Description),
			),
		),
	), nil
}

// Destroy ...
func (w *Workloads) Destroy(hx *htmx.Htmx) error {
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
