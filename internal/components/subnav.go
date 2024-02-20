package components

import htmx "github.com/zeiss/fiber-htmx"

// SubNavProps ...
type SubNavProps struct{}

// SubNav ...
func SubNav(p SubNavProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"navbar":         true,
			"bg-base-100":    true,
			"w-full":         true,
			"border-neutral": true,
			"border-b":       true,
			"border-t":       true,
			"px-6":           true,
		},
		htmx.Div(
			htmx.ClassNames{
				"navbar-start": true,
			},
			htmx.Group(children...),
			htmx.A(
				htmx.ClassNames{
					"btn":       true,
					"btn-ghost": true,
					"lg:hidden": true,
				},
				htmx.Text("Menu"),
			),
		),
	)
}
