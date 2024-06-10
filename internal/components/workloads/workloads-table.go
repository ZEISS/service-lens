package workloads

import (
	"fmt"

	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"

	htmx "github.com/zeiss/fiber-htmx"
)

const (
	profileShowURL     = "/profiles/%s"
	environmentShowURL = "/environments/%s"
	deleteWorkloadURL  = "/workloads/%s"
)

// WorkloadsTableProps ...
type WorkloadsTableProps struct {
	Workloads []*models.Workload
	Offset    int
	Limit     int
	Total     int
}

// WorkloadsTable ...
func WorkloadsTable(props WorkloadsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "workloads-tables",
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
									URL:    "/workloads",
								},
							),

							tables.Select(
								tables.SelectProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									Limits: tables.DefaultLimits,
									URL:    "/workloads",
								},
							),
							tables.Next(
								tables.PaginationProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									URL:    "/workloads",
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
						htmx.Href("/workloads/new"),
						buttons.Outline(
							buttons.ButtonProps{
								ClassNames: htmx.ClassNames{
									"btn-sm": true,
								},
							},
							htmx.Text("Create Workload"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Workload]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
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
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: "/workloads/" + row.ID.String(),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					ID:          "profile",
					AccessorKey: "profile",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Profile"))
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf(profileShowURL, row.Profile.ID),
								},
								htmx.Text(row.Profile.Name),
							),
						)
					},
				},
				{
					ID:          "environment",
					AccessorKey: "environment",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Environment"))
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf(environmentShowURL, row.Environment.ID),
								},
								htmx.Text(row.Environment.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							dropdowns.Dropdown(
								dropdowns.DropdownProps{},
								dropdowns.DropdownButton(
									dropdowns.DropdownButtonProps{},
									icons.BoltOutline(
										icons.IconProps{},
									),
								),
								dropdowns.DropdownMenuItems(
									dropdowns.DropdownMenuItemsProps{},
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
										buttons.Error(
											buttons.ButtonProps{},
											htmx.HxDelete(fmt.Sprintf(deleteWorkloadURL, row.ID)),
											htmx.HxConfirm("Are you sure you want to delete this workload?"),
											htmx.Text("Delete"),
										),
									),
								),
							),
						)
					},
				},
			},
			props.Workloads,
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
