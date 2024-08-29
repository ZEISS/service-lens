package workloads

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/models"
)

// WorkloadTagsCardProps ...
type WorkloadTagsCardProps struct {
	ClassNames htmx.ClassNames
	Workload   models.Workload
}

// WorkloadTagsCard ...
func WorkloadTagsCard(props WorkloadTagsCardProps) htmx.Node {
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
				htmx.Text("Tags"),
			),
			htmx.Div(
				htmx.ID("tags"),
				htmx.Group(
					htmx.ForEach(props.Workload.Tags, func(tag models.Tag, idx int) htmx.Node {
						return WorkloadTag(
							WorkloadTagProps{
								WorkloadID: props.Workload.ID,
								Tag:        tag,
							},
						)
					},
					)...,
				),
			),
			AddTagModal(
				AddTagModalProps{
					WorkloadID: props.Workload.ID,
				},
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{
						Type: "button",
					},
					htmx.OnClick("add_tag_modal.showModal()"),
					htmx.Text("Add Tag"),
				),
			),
		),
	)
}
