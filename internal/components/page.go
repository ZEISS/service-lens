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
					htmx.Attribute("href", "https://unpkg.com/daisyui@4.12.13/dist/full.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
					htmx.CrossOrigin("anonymous"),
				),
				htmx.Link(
					htmx.Rel("stylesheet"),
					htmx.Href("https://unpkg.com/fiber-htmx@1.3.31/dist/out.css"),
					htmx.Attribute("type", "text/css"),
					htmx.CrossOrigin("anonymous"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/htmx.org@2.0.2"),
					htmx.CrossOrigin("anonymous"),
				),
				htmx.Script(
					htmx.Src("https://unpkg.com/fiber-htmx@1.3.31"),
					htmx.CrossOrigin("anonymous"),
					htmx.Defer(),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/alpinejs@3.14.3/dist/cdn.min.js"),
					htmx.CrossOrigin("anonymous"),
					htmx.Defer(),
				),
			}, props.Head...),
		},
		htmx.Body(
			htmx.Group(children...),
		),
	)
}
