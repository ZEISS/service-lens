package templates

import (
	"fmt"
	"time"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// TemplatesTableProps ...
type TemplatesTableProps struct {
	URL       string
	Templates []*models.Template
	Offset    int
	Limit     int
	Total     int
}

// TemplatesTable ...
func TemplatesTable(props TemplatesTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "templates-tables",
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
					htmx.A(
						htmx.Href(utils.CreateTemplateUrlFormat),
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Text("Create Template"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Template]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.Template) htmx.Node {
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
					Cell: func(p tables.TableProps, row *models.Template) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf(utils.ShowTemplateUrlFormat, row.ID.String()),
								},
								htmx.Text(row.Name),
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
					Cell: func(p tables.TableProps, row *models.Template) htmx.Node {
						return htmx.Td(htmx.Text(row.CreatedAt.Format(time.RFC822)))
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.Template) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-sm": true,
									},
								},
								htmx.HxDelete(fmt.Sprintf(utils.DeleteTemplateUrlFormat, row.ID)),
								htmx.HxConfirm("Are you sure you want to delete this template?"),
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
			props.Templates,
		),
	)
}
