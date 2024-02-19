package controllers

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Workloads ...
type Workloads struct {
	db ports.Repository
}

// NewWorkloadsController ...
func NewWorkloadsController(db ports.Repository) *Workloads {
	return &Workloads{db}
}

// Store ...
func (w *Workloads) Store(hx *htmx.Htmx) error {
	profileId, err := uuid.Parse(hx.Ctx().FormValue("profile"))
	if err != nil {
		return err
	}

	workload := &models.Workload{
		Name:        hx.Ctx().FormValue("name"),
		Description: hx.Ctx().FormValue("description"),
		ProfileID:   profileId,
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
	profiles, err := w.db.ListWorkloads(c.Context(), &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return nil, err
	}

	items := make([]htmx.Node, len(profiles))
	for i, profile := range profiles {
		items[i] = htmx.Tr(
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
			htmx.Td(htmx.Text(profile.Name)),
			htmx.Td(htmx.Text(profile.Description)),
		)
	}

	return components.Page(
		components.PageProps{}.WithContext(c),
		components.SubNav(
			components.SubNavProps{},
			htmx.A(
				htmx.ClassNames{
					"btn": true,
				},
				htmx.Attribute("href", "/workloads/new"),
				htmx.Text("Create Workload"),
			),
		),
		htmx.Div(
			htmx.ClassNames{"overflow-x-auto": true},
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
					items...,
				),
			),
			htmx.Div(
				htmx.ClassNames{},
				htmx.Select(
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
	), nil
}

// New ...
func (w *Workloads) New(c *fiber.Ctx) (htmx.Node, error) {
	profiles, err := w.db.ListProfiles(c.Context(), &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return nil, err
	}

	items := make([]htmx.Node, len(profiles))
	for i, profile := range profiles {
		items[i] = htmx.Option(
			htmx.Attribute("value", profile.ID.String()),
			htmx.Text(profile.Name),
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
				htmx.Group(items...),
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
