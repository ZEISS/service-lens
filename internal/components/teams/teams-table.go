package teams

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
)

// TeamsTableProps ...
type TeamsTableProps struct {
	Teams  []*models.Team
	Offset int
	Limit  int
	Total  int
}

// TeamsTableProps ...
func TeamsTable(props TeamsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "teams-tables",
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
									URL:    "/teams",
								},
							),

							tables.Select(
								tables.SelectProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									Limits: tables.DefaultLimits,
									URL:    "/teams",
								},
							),
							tables.Next(
								tables.PaginationProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									URL:    "/teams",
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
						htmx.Href("/environments/new"),
						buttons.Outline(
							buttons.ButtonProps{
								ClassNames: htmx.ClassNames{
									"btn-sm": true,
								},
							},
							htmx.Text("Create Team"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Team]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.Team) htmx.Node {
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
					Cell: func(p tables.TableProps, row *models.Team) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: "/teams/" + row.ID.String(),
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
					Cell: func(p tables.TableProps, row *models.Team) htmx.Node {
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
			props.Teams,
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
