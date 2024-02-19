package components

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

// PageProps is the properties for the Page component.
type PageProps struct {
	Title    string
	Path     string
	Children []htmx.Node

	ctx *fiber.Ctx
}

// WithContext returns a new PageProps with the given context.
func (p PageProps) WithContext(ctx *fiber.Ctx) PageProps {
	p.ctx = ctx

	return p
}

// Context ...
func (p PageProps) Context() *fiber.Ctx {
	return p.ctx
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
			Body: []htmx.Node{
				Layout(
					LayoutProps{}.WithContext(p.Context()),
					htmx.Group(children...),
				),
			},
		}.WithContext(p.Context()))
}
