package tags

import (
	"fmt"
	"time"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/badges"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// TagsTableProps ...
type TagsTableProps struct {
	Tags   []*models.Tag
	Offset int
	Limit  int
	Total  int
}

// TagsTable ...
func TagsTable(props TagsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "tags-tables",
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
					NewTagModal(),
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-sm": true,
							},
						},
						htmx.OnClick("new_tag_modal.showModal()"),
						htmx.Text("Create Tag"),
					),
					// htmx.A(
					// 	htmx.Href("tags/new"),
					// 	buttons.Outline(
					// 		buttons.ButtonProps{
					// 			ClassNames: htmx.ClassNames{
					// 				"btn-sm": true,
					// 			},
					// 		},
					// 		htmx.Text("Create Tags"),
					// 	),
					// ),
				),
			},
			[]tables.ColumnDef[*models.Tag]{
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps, row *models.Tag) htmx.Node {
						return htmx.Td(
							badges.Primary(
								badges.BadgeProps{},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					ID:          "value",
					AccessorKey: "value",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Value"))
					},
					Cell: func(p tables.TableProps, row *models.Tag) htmx.Node {
						return htmx.Td(htmx.Text(row.Value))
					},
				},
				{
					ID:          "created_at",
					AccessorKey: "created_at",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Created At"))
					},
					Cell: func(p tables.TableProps, row *models.Tag) htmx.Node {
						return htmx.Td(htmx.Text(row.CreatedAt.Format(time.RFC822)))
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.Tag) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-sm": true,
									},
								},
								htmx.HxDelete(fmt.Sprintf(utils.DeleteTagUrlFormat, row.ID)),
								htmx.HxConfirm("Are you sure you want to delete this tag?"),
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
			props.Tags,
		),
	)
}
