package designs

import (
	"fmt"
	"time"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// DesignsTableProps ...
type DesignsTableProps struct {
	Designs   []*models.Design
	Templates []*models.Template
	URL       string
	Search    string
	Offset    int
	Limit     int
	Total     int
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
					NewDesignModal(),
					buttons.Button(
						buttons.ButtonProps{},
						htmx.OnClick("new_design_modal.showModal()"),
						htmx.Text("Add Design"),
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
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-sm": true,
									},
								},
								htmx.HxDelete(fmt.Sprintf(utils.DeleteDesignUrlFormat, row.ID)),
								htmx.HxConfirm("Are you sure you want to delete this design?"),
								htmx.HxTarget("closest tr"),
								htmx.HxSwap("outerHTML swap:1s"),
								icons.TrashOutline(
									icons.IconProps{
										ClassNames: htmx.ClassNames{
											"w-6 h-6": false,
											"w-4":     true,
											"h-4":     true,
										},
									},
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
