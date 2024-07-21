package designs

import (
	"fmt"
	"time"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignsTableProps ...
type DesignsTableProps struct {
	Designs []*models.Design
	Offset  int
	Limit   int
	Total   int
}

// DesignsTable ...
func DesignsTable(props DesignsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "designs-tables",
				Pagination: tables.TablePagination(
					tables.TablePaginationProps{},
					tables.Pagination(
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
								URL:    "/designs",
							},
						),

						tables.Select(
							tables.SelectProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								Limits: tables.DefaultLimits,
								URL:    "/designs",
							},
						),
						tables.Next(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								URL:    "/designs",
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
						htmx.Href("designs/new"),
						buttons.Outline(
							buttons.ButtonProps{
								ClassNames: htmx.ClassNames{
									"btn-sm": true,
								},
							},
							htmx.Text("Create Design"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Design]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.Design) htmx.Node {
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
					Cell: func(p tables.TableProps, row *models.Design) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf(utils.ShowDesigUrlFormat, row.ID),
								},
								htmx.Text(row.Title),
							),
						)
					},
				},
				{
					ID:          "created_at",
					AccessorKey: "created_at",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Created At"))
					},
					Cell: func(p tables.TableProps, row *models.Design) htmx.Node {
						return htmx.Td(htmx.Text(row.CreatedAt.Format(time.RFC822)))
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.Design) htmx.Node {
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
											buttons.ButtonProps{
												ClassNames: htmx.ClassNames{
													"btn-sm": true,
												},
											},
											htmx.HxDelete(fmt.Sprintf("//%s", row.ID)),
											htmx.HxConfirm("Are you sure you want to delete this design?"),
											htmx.Text("Delete"),
										),
									),
								),
							),
						)
					},
				},
			},
			props.Designs,
		),
	)
}
