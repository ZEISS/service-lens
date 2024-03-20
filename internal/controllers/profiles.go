package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// Profiles ...
type Profiles struct {
	db ports.Repository
}

// NewProfilesController ...
func NewProfilesController(db ports.Repository) *Profiles {
	return &Profiles{db}
}

// Store ...
func (p *Profiles) Store(hx *htmx.Htmx) error {
	profile := &models.Profile{
		Name:        hx.Ctx().FormValue("name"),
		Description: hx.Ctx().FormValue("description"),
	}

	err := p.db.NewProfile(hx.Ctx().Context(), profile)
	if err != nil {
		return err
	}

	hx.Redirect("/profiles/" + profile.ID.String())

	return nil
}

// New ...
func (p *Profiles) New(c *fiber.Ctx) (htmx.Node, error) {
	ctx := htmx.FromContext(c)

	return components.Page(
		ctx,
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
							Href:  "/profiles/list",
							Title: "Profiles",
						},
					),
				),
			),
		),
		components.Wrap(
			components.WrapProps{},
			htmx.FormElement(
				htmx.HxPost("/profiles"),
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
							htmx.Text("What is your name?"),
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
				htmx.Button(
					htmx.ClassNames{
						"btn":         true,
						"btn-default": true,
						"my-4":        true,
					},
					htmx.Attribute("type", "submit"),
					htmx.Text("Create Profile"),
				),
			),
		),
	), nil
}

// Show ...
func (p *Profiles) Show(c *fiber.Ctx) (htmx.Node, error) {
	ctx := htmx.FromContext(c)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return nil, err
	}

	profile, err := p.db.FetchProfile(c.Context(), id)
	if err != nil {
		return nil, err
	}

	return components.Page(
		ctx,
		components.PageProps{
			Children: []htmx.Node{
				htmx.FormElement(
					htmx.HxPost("/profiles"),
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
								htmx.Text("What is your name?"),
							),
						),
						htmx.Input(
							htmx.Attribute("type", "text"),
							htmx.Attribute("name", "name"),
							htmx.Attribute("placeholder", "Name ..."),
							htmx.Attribute("value", profile.Name),
							htmx.Attribute("readonly", "true"),
							htmx.Attribute("disabled", "true"),
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
							htmx.Attribute("value", profile.Description),
							htmx.Attribute("readonly", "true"),
							htmx.Attribute("disabled", "true"),
							htmx.ClassNames{
								"input":          true,
								"input-bordered": true,
								"w-full":         true,
								"max-w-lg":       true,
							},
						),
					),
					htmx.Div(
						htmx.ClassNames{
							"divider": true,
						},
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
							htmx.Text("Last updated"),
						),
						htmx.H3(
							htmx.Text(profile.UpdatedAt.Format("2006-01-02 15:04:05")),
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
								profile.CreatedAt.Format("2006-01-02 15:04:05"),
							),
						),
					),
				),
			},
		},
	), nil
}

// List ...
func (p *Profiles) List(hx *htmx.Htmx) error {
	limit, offset := tables.PaginationPropsFromContext(hx.Ctx())

	team, err := utils.TeamFromContext(hx.Ctx())
	if err != nil {
		return err
	}

	profiles, err := p.db.ListProfiles(hx.Ctx().Context(), team, &models.Pagination{Limit: limit, Offset: offset})
	if err != nil {
		return err
	}

	table := tables.Table[*models.Profile](
		tables.TableProps[*models.Profile]{
			Columns: []tables.ColumnDef[*models.Profile]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps[*models.Profile]) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps[*models.Profile], row *models.Profile) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps[*models.Profile]) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps[*models.Profile], row *models.Profile) htmx.Node {
						return htmx.Td(
							htmx.Text(row.Name),
						)
					},
				},
				{
					ID:          "description",
					AccessorKey: "description",
					Header: func(p tables.TableProps[*models.Profile]) htmx.Node {
						return htmx.Th(htmx.Text("Description"))
					},
					Cell: func(p tables.TableProps[*models.Profile], row *models.Profile) htmx.Node {
						return htmx.Td(
							htmx.Text(row.Description),
						)
					},
				},
			},
			Rows: tables.NewRows(profiles),
		},
		htmx.ID("data-table"),
	)

	if hx.IsHxRequest() {
		hx.ReplaceURL(fmt.Sprintf("%s?limit=%d,offset=%d", hx.Ctx().Path(), limit, offset))

		return hx.RenderComp(table)
	}

	ctx := htmx.FromContext(hx.Ctx())

	return hx.RenderComp(components.Page(
		ctx,
		components.PageProps{},
		components.Layout(
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
								Href:  "/profiles/list",
								Title: "Profiles",
							},
						),
					),
				),
				components.SubNavActions(
					components.SubNavActionsProps{},
					links.Link(
						links.LinkProps{
							Href: "/profiles/new",
							ClassNames: htmx.ClassNames{
								"btn":         true,
								"btn-outline": true,
								"btn-xs":      true,
								"link-hover":  true,
							},
						},
						htmx.Text("Create Profile"),
					),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"overflow-x-auto": true,
				},
				table,
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					tables.Pagination(
						tables.PaginationProps{
							Limit:  limit,
							Offset: offset,
						},
						tables.Prev(
							tables.PaginationProps{
								URL:    "/api/data",
								Offset: offset,
								Limit:  limit,
							},
						),
						tables.Next(
							tables.PaginationProps{
								URL: "/api/data",
							},
						),
					),
				),
			),
		),
	),
	)
}
