package components

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// TableProps ...
type TableProps[R tables.Row] struct {
	// ID is the unique identifier for the table.
	ID string
	// Rows is the data to be displayed in the table.
	Rows   []R
	URL    string
	Search string
	Offset int
	Limit  int
	Total  int
}

// Table ...
func Table[R tables.Row](props TableProps[R], columns ...tables.ColumnDef[R]) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "designs-tables",
				Pagination: tables.TablePagination(
					tables.TablePaginationProps{},
					tables.Pagination(
						tables.PaginationProps{},
						tables.Prev(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								URL:    props.URL,
							},
						),
						tables.Select(
							tables.SelectProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								Limits: tables.DefaultLimits,
								URL:    props.URL,
							},
						),
						tables.Next(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								URL:    props.URL,
							},
						),
					),
				),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps{
						ClassNames: htmx.ClassNames{
							"flex":            true,
							"items-center":    true,
							"justify-between": true,
						},
					},
					htmx.Div(
						htmx.ClassNames{
							"inline-flex":  true,
							"items-center": true,
							"gap-3":        true,
						},
						tables.Search(
							tables.SearchProps{
								Name:        "search",
								Placeholder: "Search ...",
								URL:         props.URL,
							},
						),
					),
					// dropdowns.Dropdown(
					// 	dropdowns.DropdownProps{
					// 		ClassNames: htmx.ClassNames{
					// 			"dropdown-end": true,
					// 		},
					// 	},
					// 	dropdowns.DropdownButton(
					// 		dropdowns.DropdownButtonProps{},
					// 		htmx.Text("Create Design"),
					// 	),
					// 	dropdowns.DropdownMenuItems(
					// 		dropdowns.DropdownMenuItemsProps{},
					// 		dropdowns.DropdownMenuItem(
					// 			dropdowns.DropdownMenuItemProps{},
					// 			htmx.Group(
					// 				htmx.ForEach(props.Templates, func(template *models.Template, idx int) htmx.Node {
					// 					return htmx.A(
					// 						htmx.Href(fmt.Sprintf(utils.CreateDesignUrlFormat, template.ID)),
					// 						htmx.Text(template.Name),
					// 					)
					// 				})...,
					// 			),
					// 			htmx.A(
					// 				htmx.Href(fmt.Sprintf(utils.CreateDesignUrlFormat, "_blank")),
					// 				htmx.Text("Blank Template"),
					// 			),
					// 		),
					// 	),
					// ),
				),
			},
			columns,
			props.Rows,
		),
	)
}
