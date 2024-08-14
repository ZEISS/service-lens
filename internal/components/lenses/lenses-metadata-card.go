package lenses

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/service-lens/internal/models"
)

// LensMetadataCardProps ...
type LensMetadataCardProps struct {
	Lens models.Lens
}

// LensMetadataCard ...
func LensMetadataCard(props LensMetadataCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.ClassNames{
				"m-2": true,
			},
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text("Metadata"),
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
					htmx.Text("ID"),
				),
				htmx.H3(htmx.Text(conv.String(props.Lens.ID))),
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
							props.Lens.CreatedAt.Format("2006-01-02 15:04:05"),
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
							props.Lens.UpdatedAt.Format("2006-01-02 15:04:05"),
						),
					),
				),
			),
		),
	)
}
