package components

import htmx "github.com/zeiss/fiber-htmx"

// NavbarProps is the properties for the Navbar component.
type NavbarProps struct {
	Children []htmx.Node
}

// Navbar is a whole document to output.
func Navbar(p NavbarProps) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{"navbar": true, "bg-base-100": true},
		htmx.Div(
			htmx.ClassNames{"navbar-start": true},
			htmx.Div(
				htmx.ClassNames{"dropdown": true},
				htmx.Div(
					htmx.ClassNames{"btn": true, "btn-ghost": true, "lg:hidden": true},
					htmx.SVG(
						htmx.ClassNames{"h-5": true, "w-5": true},
						htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
						htmx.Attribute("fill", "none"),
						htmx.Attribute("viewBox", "0 0 24 24"),
						htmx.Attribute("stroke", "currentColor"),
						htmx.Path(
							htmx.Attribute("stroke-linecap", "round"),
							htmx.Attribute("stroke-linejoin", "round"),
							htmx.Attribute("stroke-width", "2"),
							htmx.Attribute("d", "M4 6h16M4 12h8m-8 6h16"),
						),
					),
				),
				htmx.Ul(
					htmx.Attribute("tabindex", "0"),
					htmx.ClassNames{
						"menu":             true,
						"menu-sm":          true,
						"dropdown-content": true,
						"mt-3":             true,
						"z-[1]":            true,
						"p-2":              true,
						"shadow":           true,
						"bg-base-100":      true,
						"rounded-box":      true,
						"w-52":             true,
					},
					htmx.Li(htmx.A(htmx.Text("Item 1"))),
					htmx.Li(
						htmx.A(htmx.Text("Parent")),
						htmx.Ul(
							htmx.ClassNames{"p-2": true},
							htmx.Li(htmx.A(htmx.Text("Submenu 1"))),
							htmx.Li(htmx.A(htmx.Text("Submenu 2"))),
						),
					),
					htmx.Li(htmx.A(htmx.Text("Item 3"))),
				),
			),
			htmx.A(
				htmx.ClassNames{
					"btn":       true,
					"btn-ghost": true,
					"text-xl":   true,
				},
				htmx.Text("Service Lens"),
			),
		),
		htmx.Div(
			htmx.ClassNames{"navbar-center": true, "hidden": true, "lg:flex": true},
			htmx.Ul(
				htmx.ClassNames{"menu": true, "menu-horizontal": true, "px-1": true},
				htmx.Li(htmx.A(htmx.Text("Item 1"))),
				htmx.Li(
					htmx.Details(
						htmx.Summary(htmx.Text("Parent")),
						htmx.Ul(
							htmx.ClassNames{"p-2": true},
							htmx.Li(htmx.A(htmx.Text("Submenu 1"))),
							htmx.Li(htmx.A(htmx.Text("Submenu 2"))),
						),
					),
				),
				htmx.Li(
					htmx.A(
						htmx.ClassNames{
							"active": true,
						},
						htmx.Attribute("href", "/profiles"),
						htmx.Text("Profiles"),
					),
				),
				htmx.Li(
					htmx.A(
						htmx.ClassNames{},
						htmx.Attribute("href", "/lenses"),
						htmx.Text("Lenses"),
					),
				),
			),
		),
		htmx.Div(
			htmx.ClassNames{"navbar-end": true},
			htmx.A(
				htmx.ClassNames{
					"btn":       true,
					"btn-ghost": true,
				},
				htmx.Text("Add Profile"),
				htmx.Attribute("href", "/profiles/new"),
			),
			UserNav(UserNavProps{}),
		),
	)
}
