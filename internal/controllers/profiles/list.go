package profiles

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// ProfileListController ...
type ProfileListController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewProfileListController ...
func NewProfileListController(db ports.Repository) *ProfileListController {
	return &ProfileListController{db, htmx.UnimplementedController{}}
}

// Get ...
func (p *ProfileListController) Get() error {
	hx := p.Hx

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
