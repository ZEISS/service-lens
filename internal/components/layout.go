package components

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	Children []htmx.Node

	htmx.Ctx
}

// WrapProps ...
type WrapProps struct {
	ClassName map[string]bool
}

// Wrap ...
func Wrap(p WrapProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"p-6": true,
		}.Merge(p.ClassName),
		htmx.Group(children...),
	)
}

// Layout is a whole document to output.
func Layout(p LayoutProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"drawer": true,
		},
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
			Navbar(
				NavbarProps{
					Ctx: p.Ctx,
				},
			),
			htmx.Div(
				htmx.ClassNames{},
				htmx.Group(children...),
			),
		),
		htmx.Div(
			htmx.ClassNames{
				"drawer-side": true,
			},
			htmx.Label(
				htmx.Attribute(
					"for",
					"app-drawer",
				),
				htmx.Attribute(
					"aria-label",
					"close sidebar"),
				htmx.ClassNames{
					"drawer-overlay": true,
				},
			),
			htmx.Ul(
				htmx.ClassNames{
					"menu":        true,
					"p-4":         true,
					"w-80":        true,
					"min-h-full":  true,
					"bg-base-200": true,
				},
				htmx.Li(
					htmx.A(
						htmx.Text(
							"Sidebar Item 1",
						),
					),
				),
				htmx.Li(
					htmx.A(
						htmx.Text(
							"Sidebar Item 2",
						),
					),
				),
			),
		),
	)
}