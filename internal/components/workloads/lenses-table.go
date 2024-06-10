package workloads

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
)

const (
	deleteProfileURL = "/profiles/%s"
	workloadLensURL  = "/workloads/%s/lenses/%s"
)

// LensesTableProps ...
type LensesTableProps struct {
	Workload *models.Workload
	Offset   int
	Limit    int
	Total    int
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
						// Search(
						// 	SearchProps{
						// 		URL:    "/lenses",
						// 		Limit:  props.Limit,
						// 		Offset: props.Offset,
						// 	},
						// ),
						forms.TextInputBordered(
							forms.TextInputProps{
								ClassNames: htmx.ClassNames{
									"input-sm": true,
								},
								Placeholder: "Search ...",
							},
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
									Href: fmt.Sprintf(workloadLensURL, props.Workload.ID, row.ID),
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
							dropdowns.Dropdown(
								dropdowns.DropdownProps{},
								dropdowns.DropdownButton(
									dropdowns.DropdownButtonProps{},
									icons.BoltOutline(
										icons.IconProps{},
									),
								),
								// dropdowns.DropdownMenuItems(
								// 	dropdowns.DropdownMenuItemsProps{},
								// 	dropdowns.DropdownMenuItem(
								// 		dropdowns.DropdownMenuItemProps{},
								// 		buttons.Error(
								// 			buttons.ButtonProps{},
								// 			htmx.HxDelete(fmt.Sprintf(deleteProfileURL, props.Workload.ID, row.ID)),
								// 			htmx.HxConfirm("Are you sure you want to delete this lens?"),
								// 			htmx.Text("Delete"),
								// 		),
								// 	),
								// ),
							),
						)
					},
				},
			},
			props.Workload.Lenses,
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

// SearchProps ...
type SearchProps struct {
	ClassNames htmx.ClassNames
	URL        string
	Limit      int
	Offset     int
	Search     string
	Sort       string
}

// Search ...
func Search(p SearchProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Action(fmt.Sprintf("%s?limit=%d&offset=%d&search=%s", p.URL, p.Limit, p.Offset, p.Search)),
		forms.FormControl(
			forms.FormControlProps{
				ClassNames: htmx.ClassNames{
					"py-4": true,
				},
			},

			forms.TextInputBordered(
				forms.TextInputProps{
					Name:  "search",
					Value: p.Search,
				},
				htmx.Type("search"),
				htmx.HxGet(fmt.Sprintf("%s?limit=%d&offset=%d&search=%s", p.URL, p.Limit, p.Offset, p.Search)),
				htmx.HxTrigger("search, keyup delay:200ms changed"),
			),
		),
	)
}
