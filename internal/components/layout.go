package components

import htmx "github.com/zeiss/fiber-htmx"

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	Children []htmx.Node
}

// Layout is a whole document to output.
func Layout(p LayoutProps) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{"drawer": true},
		htmx.Input(
			htmx.Attribute("id", "app-drawer"),
			htmx.Attribute("type", "checkbox"),
			htmx.Attribute("class", "drawer-toggle"),
		),
		htmx.Div(
			htmx.ClassNames{
				"drawer-content": true,
				"flex":           true,
				"flex-col":       true,
			},
			htmx.Div(
				htmx.ClassNames{
					"w-full":      true,
					"navbar":      true,
					"bg-base-300": true,
				},
				htmx.Div(
					htmx.ClassNames{
						"flex-none": true,
						"lg:hidden": true,
					},
					htmx.LabElement(
						htmx.ClassNames{
							"btn":        true,
							"btn-square": true,
							"btn-ghost":  true,
						},
						htmx.Attribute("for", "app-drawer"),
						htmx.Attribute("aria-label", "open sidebar"),
						htmx.SVG(
							htmx.ClassNames{
								"inline-block":   true,
								"w-6":            true,
								"h-6":            true,
								"stroke-current": true,
							},
							htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
							htmx.Attribute("fill", "none"),
							htmx.Attribute("viewBox", "0 0 24 24"),
							htmx.Path(
								htmx.Attribute("stroke-linecap", "round"),
								htmx.Attribute("stroke-linejoin", "round"),
								htmx.Attribute("stroke-width", "2"),
								htmx.Attribute("d", "M4 6h16M4 12h16M4 18h16"),
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{"flex-1": true, "px-2": true, "mx-2": true},
					htmx.Text("Navbar Title"),
				),
				htmx.Div(
					htmx.ClassNames{"flex-none": true, "hidden": true, "lg:block": true},
					htmx.Div(
						htmx.ClassNames{
							"flex":            true,
							"justify-between": true,
							"items-center":    true,
						},
						htmx.Ul(
							htmx.ClassNames{
								"menu":            true,
								"menu-horizontal": true,
							},
							htmx.Button(
								htmx.ClassNames{
									"btn":        true,
									"btn-square": true,
									"btn-ghost":  true,
								},
								htmx.SVG(
									htmx.ClassNames{"inline-block": true, "w-5": true, "h-5": true, "stroke-current": true},
									htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
									htmx.Attribute("fill", "none"),
									htmx.Attribute("viewBox", "0 0 24 24"),
									htmx.Path(
										htmx.Attribute("stroke-linecap", "round"),
										htmx.Attribute("stroke-linejoin", "round"),
										htmx.Attribute("stroke-width", "2"),
										htmx.Attribute("d", "M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"),
									),
								),
							),
							htmx.Li(htmx.A(htmx.Text("Navbar Item 1"))),
							htmx.Li(htmx.A(htmx.Text("Navbar Item 2"))),
						),
						UserNav(UserNavProps{}),
					),
				),
			),
			htmx.Div(p.Children...),
		),
		htmx.Div(htmx.ClassNames{"drawer-side": true},
			htmx.LabElement(htmx.Attribute("for", "app-drawer"), htmx.Attribute("aria-label", "close sidebar"), htmx.ClassNames{"drawer-overlay": true}),
			htmx.Ul(htmx.ClassNames{"menu": true, "p-4": true, "w-80": true, "min-h-full": true, "bg-base-200": true},
				htmx.Li(htmx.A(htmx.Text("Sidebar Item 1"))),
				htmx.Li(htmx.A(htmx.Text("Sidebar Item 2"))),
			),
		),
	)
}
