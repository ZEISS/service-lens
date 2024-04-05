package components

import htmx "github.com/zeiss/fiber-htmx"

// CardDataBlockProps ...
type CardDataBlockProps struct {
	Title string
	Data  string
}

// CardDataBlock ...
func CardDataBlock(props *CardDataBlockProps) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"flex":     true,
			"flex-col": true,
			"py-2":     true,
		},
		htmx.H4(
			htmx.ClassNames{
				"text-gray-500": true,
			},
			htmx.Text(props.Title),
		),
		htmx.H3(
			htmx.Text(props.Data),
		),
	)
}
