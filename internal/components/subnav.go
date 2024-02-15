package components

import htmx "github.com/zeiss/fiber-htmx"

// SubNavProps ...
type SubNavProps struct{}

// SubNav ...
func SubNav(p SubNavProps) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{"navbar": true, "bg-base-100": true},
		htmx.Div(
			htmx.ClassNames{
				"flex":     true,
				"flex-row": true,
			},
			htmx.Div(),
			htmx.H1(htmx.Text("Profiles")),
			htmx.Div(
				htmx.Button(
					htmx.ClassNames{
						"btn": true,
					},
					htmx.Text("Refresh"),
				),
				htmx.A(
					htmx.ClassNames{
						"btn": true,
					},
					htmx.Attribute("href", "/profiles/new"),
					htmx.Text("Create Profile"),
				),
			),
		),
	)
}
