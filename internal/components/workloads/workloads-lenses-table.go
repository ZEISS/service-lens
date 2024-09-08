package workloads

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// LensesTableProps ...
type LensesTableProps struct {
	Workload models.Workload
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
									Href: fmt.Sprintf(utils.WorkloadLensUrlFormat, props.Workload.ID, row.ID),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					ID:          "updated_at",
					AccessorKey: "updated_at",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Last Updated"))
					},
					Cell: func(p tables.TableProps, row *models.Lens) htmx.Node {
						return htmx.Td(htmx.Text(row.UpdatedAt.Format("2006-01-02 15:04:05")))
					},
				},
			},
			props.Workload.Lenses,
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
