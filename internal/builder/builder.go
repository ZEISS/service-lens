package builder

import (
	"fmt"
	"text/template"

	"github.com/yuin/goldmark/ast"
	render "github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// MarkdownBuilder ...
type MarkdownBuilder struct {
	html.Config
}

// NewMarkdownBuilder ...
func NewMarkdownBuilder(opts ...html.Option) render.NodeRenderer {
	builder := &MarkdownBuilder{
		Config: html.NewConfig(),
	}

	for _, opt := range opts {
		opt.SetHTMLOption(&builder.Config)
	}

	return builder
}

func (builder MarkdownBuilder) RegisterFuncs(registerer render.NodeRendererFuncRegisterer) {
	registerer.Register(ast.KindHeading, builder.Heading)
	registerer.Register(ast.KindList, builder.List)
	registerer.Register(ast.KindParagraph, builder.Paragraph)
	registerer.Register(ast.KindLink, builder.Link)
	// registerer.Register(ast.KindListItem, r.renderListItem)
}

// HeadingClasses ...
var HeadingClasses = map[int]string{
	0: "text-5xl font-bold",
	1: "text-4xl font-bold",
	2: "text-3xl font-bold",
	3: "text-2xl font-bold",
	4: "text-xl font-bold",
	5: "text-lg font-bold",
	6: "text-base font-bold",
}

// Link ...
func (r *MarkdownBuilder) Link(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	node := n.(*ast.Link)

	if entering {
		_, _ = w.WriteString("<a class=\"link link-primary\" href=\"")
		if r.Unsafe || !html.IsDangerousURL(node.Destination) {
			_, _ = w.Write(util.EscapeHTML(util.URLEscape(node.Destination, true)))
		}
		_ = w.WriteByte('"')
		if node.Title != nil {
			_, _ = w.WriteString(` title="`)
			r.Writer.Write(w, node.Title)
			_ = w.WriteByte('"')
		}
		if n.Attributes() != nil {
			html.RenderAttributes(w, n, html.LinkAttributeFilter)
		}
		_ = w.WriteByte('>')
	} else {
		_, _ = w.WriteString("</a>")
	}

	return ast.WalkContinue, nil
}

// Paragraph ...
func (r *MarkdownBuilder) Paragraph(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<p class=\"py-2 leading-relaxed\"")

			html.RenderAttributes(w, n, html.ParagraphAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<p class=\"py-2 leading-relaxed\">")
		}
	} else {
		_, _ = w.WriteString("</p>\n")
	}
	return ast.WalkContinue, nil
}

// List ...
func (r *MarkdownBuilder) List(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.List)
	tag := "ul"
	classes := "list-disc list-inside"

	if n.IsOrdered() {
		tag = "ol"
		classes = "list-decimal list-inside"
	}

	if entering {
		_ = w.WriteByte('<')
		_, _ = w.WriteString(tag)
		_, _ = w.Write([]byte(" " + "class" + `="` + template.HTMLEscapeString(classes) + `"`))

		if n.IsOrdered() && n.Start != 1 {
			_, _ = fmt.Fprintf(w, " start=\"%d\"", n.Start)
		}

		if n.Attributes() != nil {
			html.RenderAttributes(w, n, html.ListAttributeFilter)
		}

		_, _ = w.WriteString(">\n")
	} else {
		_, _ = w.WriteString("</")
		_, _ = w.WriteString(tag)
		_, _ = w.WriteString(">\n")
	}

	return ast.WalkContinue, nil
}

// Heading ...
func (r *MarkdownBuilder) Heading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)

	if entering {
		_, _ = w.WriteString("<h")
		_ = w.WriteByte("0123456"[n.Level])
		_, _ = w.Write([]byte(" " + "class" + `="` + template.HTMLEscapeString(HeadingClasses[n.Level]) + `"`))

		if n.Attributes() != nil {
			html.RenderAttributes(w, node, html.HeadingAttributeFilter)
		}

		_ = w.WriteByte('>')

	} else {
		_, _ = w.WriteString("</h")
		_ = w.WriteByte("0123456"[n.Level])
		_, _ = w.WriteString(">\n")
	}

	return ast.WalkContinue, nil
}
