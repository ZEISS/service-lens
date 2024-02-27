package components

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// PageProps is the properties for the Page component.
type PageProps struct {
	Title    string
	Path     string
	Children []htmx.Node

	htmx.Ctx
}

// Page is a whole document to output.
func Page(p PageProps, children ...htmx.Node) htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    p.Title,
			Language: "en",
			Head: []htmx.Node{
				htmx.Link(htmx.Attribute("href", "https://cdn.jsdelivr.net/npm/daisyui@4.7.0/dist/full.min.css"), htmx.Attribute("rel", "stylesheet"), htmx.Attribute("type", "text/css")),
				htmx.Script(htmx.Attribute("src", "https://unpkg.com/htmx.org@1.9.10"), htmx.Attribute("type", "application/javascript")),
				htmx.Script(htmx.Attribute("src", "https://cdn.tailwindcss.com"), htmx.Attribute("type", "application/javascript")),
			},
			Body: children,
		})
}
