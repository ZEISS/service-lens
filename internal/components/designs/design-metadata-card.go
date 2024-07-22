package designs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/service-lens/internal/models"
)

// DesignMetadataCardProps ...
type DesignMetadataCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
}

// DesignMetadataCard ...
func DesignMetadataCard(props DesignMetadataCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
				props.ClassNames,
			),
		},
		cards.Body(
			cards.BodyProps{},
			htmx.Div(
				htmx.Div(
					htmx.ClassNames{
						"flex":     true,
						"flex-col": true,
						"py-2":     true,
					},
					htmx.H4(
						htmx.ClassNames{
							"text-gray-500": true,
						},
						htmx.Text("Authored by"),
					),
					htmx.H3(htmx.Text(props.Design.Author.Name)),
				),
				htmx.Div(
					htmx.ClassNames{
						"flex":     true,
						"flex-col": true,
						"py-2":     true,
					},
					htmx.H4(
						htmx.ClassNames{
							"text-gray-500": true,
						},
						htmx.Text("Created at"),
					),
					htmx.H3(
						htmx.Text(
							props.Design.CreatedAt.Format("2006-01-02 15:04:05"),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"flex":     true,
						"flex-col": true,
						"py-2":     true,
					},
					htmx.H4(
						htmx.ClassNames{
							"text-gray-500": true,
						},
						htmx.Text("Updated at"),
					),
					htmx.H3(
						htmx.Text(
							props.Design.UpdatedAt.Format("2006-01-02 15:04:05"),
						),
					),
				),
			),
		),
	)
}
