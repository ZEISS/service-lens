package components

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// PageProps is the properties for the Page component.
type PageProps struct {
	Title    string
	Path     string
	Children []htmx.Node
	Head     []htmx.Node
}

// Page is a whole document to output.
func Page(props PageProps, children ...htmx.Node) htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    props.Title,
			Language: "en",
			Attributes: []htmx.Node{
				htmx.DataAttribute("theme", "light"),
			},
			Head: append([]htmx.Node{
				htmx.Link(
					htmx.Attribute("href", "https://unpkg.com/tailwindcss@2.0.1/dist/tailwind.min.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
					htmx.CrossOrigin("anonymous"),
				),
				htmx.Link(
					htmx.Attribute("href", "https://unpkg.com/daisyui@4.12.13/dist/full.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
					htmx.CrossOrigin("anonymous"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/htmx.org@2.0.2"),
					htmx.CrossOrigin("anonymous"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/alpinejs@1.1.2/dist/alpine.js"),
					htmx.CrossOrigin("anonymous"),
					htmx.Attribute("defer", ""),
				),
			}, props.Head...),
		},
		htmx.Body(
			// htmx.HxBoost(true),
			htmx.Group(children...),
		),
	)
}
