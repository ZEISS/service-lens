package lenses

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
)

// LensesTableProps ...
type LensesTableProps struct {
	Lenses []*models.Lens
	Offset int
	Limit  int
	Total  int
}

// LensesTable ...
func LensesTable(props LensesTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "lenses-tables",
				Pagination: tables.TablePagination(
					tables.TablePaginationProps{
						Pagination: tables.Pagination(
							tables.PaginationProps{
								Offset: props.Offset,
								Limit:  props.Limit,
								Total:  props.Total,
							},
							tables.Prev(
								tables.PaginationProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									URL:    "/lenses",
								},
							),

							tables.Select(
								tables.SelectProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									Limits: tables.DefaultLimits,
									URL:    "/lenses",
								},
							),
							tables.Next(
								tables.PaginationProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									URL:    "/lenses",
								},
							),
						),
					},
				),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps{
						ClassNames: htmx.ClassNames{
							"flex":            true,
							"items-center":    true,
							"justify-between": true,
							"px-5":            true,
							"pt-5":            true,
						},
					},
					htmx.Div(
						htmx.ClassNames{
							"inline-flex":  true,
							"items-center": true,
							"gap-3":        true,
						},
						forms.TextInputBordered(
							forms.TextInputProps{
								ClassNames: htmx.ClassNames{
									"input-sm": true,
								},
								Placeholder: "Search ...",
							},
						),
					),
					htmx.A(
						htmx.Href("/lenses/new"),
						buttons.Outline(
							buttons.ButtonProps{
								ClassNames: htmx.ClassNames{
									"btn-sm": true,
								},
							},
							htmx.Text("Create Lens"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Lens]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.Lens) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps, row *models.Lens) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: "/lenses/" + row.ID.String(),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.Lens) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-square": true,
									},
								},
							),
						)
					},
				},
			},
			props.Lenses,
			// Pagination: ProfileListTablePaginationComponent(
			// 	ProfileListTablePaginationProps{
			// 		Limit:  props.Limit,
			// 		Offset: props.Offset,
			// 		Total:  props.Total,
			// 		Target: "profiles-tables",
			// 		Team:   props.Team,
			// 	},
			// ),
		),
	)
}
