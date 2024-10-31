package components

import htmx "github.com/zeiss/fiber-htmx"

// MarkdownToolbar ...
func MarkdownToolbar(children ...htmx.Node) htmx.Node {
	return htmx.Element("markdown-toolbar", children...)
}

// MarkdownBold ...
func MarkdownBold(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-bold", children...)
}

// MarkdownHeader ...
func MarkdownHeader(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-header", children...)
}

// MarkdownItalic ...
func MarkdownItalic(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-italic", children...)
}

// MarkdownQuote ...
func MarkdownQuote(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-quote", children...)
}

// MarkdownCode ...
func MarkdownCode(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-code", children...)
}

// MarkdownLink ...
func MarkdownLink(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-link", children...)
}

// MarkdownImage ...
func MarkdownImage(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-image", children...)
}

// MarkdownUnorderedList ...
func MarkdownUnorderedList(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-unordered-list", children...)
}

// MarkdownOrderedList ...
func MarkdownOrderedList(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-ordered-list", children...)
}

// MarkdownTaskList ...
func MarkdownTaskList(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-task-list", children...)
}

// MarkdownMention ...
func MarkdownMention(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-mention", children...)
}

// MarkdownRef ...
func MarkdownRef(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-ref", children...)
}

// MarkdownStrikethrough ...
func MarkdownStrikethrough(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-strikethrough", children...)
}
