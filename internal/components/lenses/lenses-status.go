package lenses

import htmx "github.com/zeiss/fiber-htmx"

// LensesStatusProps ...
type LensesStatusProps struct {
	// ClassNames ...
	ClassNames htmx.ClassNames
	// IsDraft ...
	IsDraft bool
}

// LensesStatus ...
func LensesStatus(props LensesStatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ID("status"),
		htmx.Merge(
			htmx.ClassNames{
				"flex":     true,
				"flex-col": true,
				"py-2":     true,
			},
			props.ClassNames,
		),
		htmx.H4(
			htmx.ClassNames{
				"text-gray-500": true,
			},
			htmx.Text("Status"),
		),
		htmx.H3(
			htmx.IfElse(props.IsDraft, htmx.Text("Draft"), htmx.Text("Published")),
		),
		htmx.Group(children...),
	)
}
