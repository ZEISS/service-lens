package workloads

import (
	"github.com/zeiss/service-lens/internal/models"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
)

// WorkloadProfileCardProps ...
type WorkloadProfileCardProps struct {
	Workload models.Workload
}

// WorkloadProfileCard ...
func WorkloadProfileCard(props WorkloadProfileCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.ClassNames{
				tailwind.M2: true,
			},
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text("Profile"),
			),
			htmx.Div(
				htmx.ClassNames{
					"flex":     true,
					"flex-col": true,
					"py-2":     true,
				},
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
						htmx.Text("Name"),
					),
					htmx.H3(
						htmx.Text(
							props.Workload.Profile.Name,
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
						htmx.Text("Description"),
					),
					htmx.H3(
						htmx.Text(
							props.Workload.Profile.Description,
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
						htmx.Text("ID"),
					),
					htmx.H3(
						htmx.Text(
							conv.String(props.Workload.Profile.ID),
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
						htmx.Text("Created at"),
					),
					htmx.H3(
						htmx.Text(
							props.Workload.Profile.CreatedAt.Format("2006-01-02 15:04:05"),
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
							props.Workload.Profile.UpdatedAt.Format("2006-01-02 15:04:05"),
						),
					),
				),
			),
		),
	)
}
