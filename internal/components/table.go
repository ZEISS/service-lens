package components

import htmx "github.com/zeiss/fiber-htmx"

// TableProps ...
type TableProps struct{}

// Table ...
func Table(p TableProps) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{"overflow-x-auto": true},
		htmx.Table(
			htmx.ClassNames{"table": true},
			htmx.THead(
				htmx.Tr(
					htmx.Th(),
					htmx.Th(htmx.Text("Name")),
					htmx.Th(htmx.Text("Job")),
					htmx.Th(htmx.Text("Favorite Color")),
				),
			),
			htmx.TBody(
				htmx.Tr(
					htmx.Th(htmx.Text("1")),
					htmx.Td(htmx.Text("Cy Ganderton")),
					htmx.Td(htmx.Text("Quality Control Specialist")),
					htmx.Td(htmx.Text("Blue")),
				),
				htmx.Tr(
					htmx.Th(htmx.Text("1")),
					htmx.Td(htmx.Text("Cy Ganderton")),
					htmx.Td(htmx.Text("Quality Control Specialist")),
					htmx.Td(htmx.Text("Blue")),
				),
				htmx.Tr(
					htmx.Th(htmx.Text("1")),
					htmx.Td(htmx.Text("Cy Ganderton")),
					htmx.Td(htmx.Text("Quality Control Specialist")),
					htmx.Td(htmx.Text("Blue")),
				),
			),
		),
	)
}
