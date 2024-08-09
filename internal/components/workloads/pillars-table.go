package workloads

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
)

// LensPillarTableProps ...
type LensPillarTableProps struct {
	Lens   *models.Lens
	Offset int
	Limit  int
	Total  int
}

// LensPillarTable ...
func LensPillarTable(props LensPillarTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "lenses-pillar-table",
			},
			[]tables.ColumnDef[*models.Pillar]{
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps, row *models.Pillar) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									// Href: fmt.Sprintf(workloadLensURL, props.Workload.ID, row.ID),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
			},
			props.Lens.GetPillars(),
		),
	)
}
