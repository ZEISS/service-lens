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
	Rows []R
	// URL is the URL to send the request to.
	URL string
	// Search is the search string.
	Search string
	// Offset is the offset of the table.
	Offset int
	// Limit is the limit of the table.
	Limit int
	// Total is the total number of rows in the table.
	Total int
	// Toolbar is the toolbar of the table.
	Toolbar htmx.Node
}

// Table ...
func Table[R tables.Row](props TableProps[R], columns ...tables.ColumnDef[R]) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: props.ID,
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
					props.Toolbar,
				),
			},
			columns,
			props.Rows,
		),
	)
}
