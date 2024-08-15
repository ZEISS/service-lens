package profiles

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/service-lens/internal/models"
)

// ProfilesMetadataCardProps ...
type ProfilesMetadataCardProps struct {
	ClassNames htmx.ClassNames
	Profile    models.Profile
}

// ProfilesMetadataCard ...
func ProfilesMetadataCard(props ProfilesMetadataCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
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
						htmx.Text("ID"),
					),
					htmx.H3(
						htmx.Text(
							conv.String(props.Profile.ID),
						),
					),
				),
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
							htmx.Text("Created at"),
						),
						htmx.H3(
							htmx.Text(
								props.Profile.CreatedAt.Format("2006-01-02 15:04:05"),
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
								props.Profile.UpdatedAt.Format("2006-01-02 15:04:05"),
							),
						),
					),
				),
			),
		),
	)
}
