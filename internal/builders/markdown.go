package builders

import (
	"fmt"
	"text/template"

	"github.com/yuin/goldmark/ast"
	ext "github.com/yuin/goldmark/extension"
	gast "github.com/yuin/goldmark/extension/ast"
	render "github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"github.com/zeiss/pkg/conv"
)

// MarkdownBuilder ...
type MarkdownBuilder struct {
	html.Config

	taskURL   string
	taskItems int
}

// WithTaskURL ...
func WithTaskURL(url string) MarkdownBuilderOption {
	return func(builder *MarkdownBuilder) {
		builder.taskURL = url
	}
}

// WithHTMLOptions ...
func WithHTMLOptions(opts ...html.Option) MarkdownBuilderOption {
	return func(builder *MarkdownBuilder) {
		for _, opt := range opts {
			opt.SetHTMLOption(&builder.Config)
		}
	}
}

// MarkdownBuilderOption ...
type MarkdownBuilderOption func(*MarkdownBuilder)

// NewMarkdownBuilder ...
func NewMarkdownBuilder(opts ...MarkdownBuilderOption) render.NodeRenderer {
	builder := &MarkdownBuilder{
		Config: html.NewConfig(),
	}

	for _, opt := range opts {
		opt(builder)
	}

	return builder
}

func (builder MarkdownBuilder) RegisterFuncs(registerer render.NodeRendererFuncRegisterer) {
	registerer.Register(ast.KindHeading, builder.Heading)
	registerer.Register(ast.KindList, builder.List)
	registerer.Register(ast.KindListItem, builder.ListItem)
	registerer.Register(ast.KindParagraph, builder.Paragraph)
	registerer.Register(ast.KindLink, builder.Link)
	registerer.Register(gast.KindTable, builder.Table)
	registerer.Register(gast.KindTaskCheckBox, builder.TaskCheckBox)
	// registerer.Register(ast.KindListItem, r.renderListItem)
}

// HeadingClasses ...
var HeadingClasses = map[int]string{
	0: "text-5xl font-bold py-4",
	1: "text-4xl font-bold py-4",
	2: "text-3xl font-bold py-4",
	3: "text-2xl font-bold py-4",
	4: "text-xl font-bold py-4",
	5: "text-lg font-bold py-4",
	6: "text-base font-bold py-4",
}

// TaskCheckBox ...
func (r *MarkdownBuilder) TaskCheckBox(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*gast.TaskCheckBox)

	r.taskItems++

	if n.IsChecked {
		_, _ = w.WriteString(`<input checked="" disabled="" type="checkbox" class="checkbox"`)
	} else {
		_, _ = w.WriteString(`<input name="task" value="` + conv.String(r.taskItems) + `" hx-trigger="click" hx-swap="outerHTML" hx-post="` + r.taskURL + `" type="checkbox" class="checkbox"`)
	}
	if r.XHTML {
		_, _ = w.WriteString(" /> ")
	} else {
		_, _ = w.WriteString("> ")
	}

	return ast.WalkContinue, nil
}

// Table ...
func (r *MarkdownBuilder) Table(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString(`<table class="table table-pin-cols my-4"`)
		if node.Attributes() != nil {
			html.RenderAttributes(w, node, ext.TableAttributeFilter)
		}
		_, _ = w.WriteString(">\n")
	} else {
		_, _ = w.WriteString("</table>\n")
	}

	return ast.WalkContinue, nil
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
		_, _ = w.WriteString(" class=\"" + template.HTMLEscapeString(classes) + "\"")

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

// ListItem ...
func (r *MarkdownBuilder) ListItem(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if node.Attributes() != nil {
			_, _ = w.WriteString(`<li`)
			html.RenderAttributes(w, node, html.ListItemAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString(`<li>`)
		}
		fc := node.FirstChild()

		if fc != nil {
			if _, ok := fc.(*ast.TextBlock); !ok {
				_ = w.WriteByte('\n')
			}
		}
	} else {
		_, _ = w.WriteString("</li>\n")
	}

	return ast.WalkContinue, nil
}

// Heading ...
func (r *MarkdownBuilder) Heading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)

	if entering {
		_, _ = w.WriteString("<h")
		_ = w.WriteByte("0123456"[n.Level])
		_, _ = w.WriteString(" class=\"" + template.HTMLEscapeString(HeadingClasses[n.Level]) + "\"")

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
