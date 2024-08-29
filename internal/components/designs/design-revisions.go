package designs

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/collapsible"
	"github.com/zeiss/fiber-htmx/components/loading"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignRevisionsCardProps ...
type DesignRevisionsCardProps struct {
	ClassNames htmx.ClassNames
	DesignID   uuid.UUID
}

// DesignRevisionsCard ...
func DesignRevisionsCard(props DesignRevisionsCardProps) htmx.Node {
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
			collapsible.CollapseArrow(
				collapsible.CollapseProps{},
				htmx.HxTrigger("click once"),
				htmx.HxGet(fmt.Sprintf(utils.ListDesignRevisionsUrlFormat, props.DesignID)),
				htmx.HxTarget(".collapse-content"),
				htmx.HxIndicator("find .htmx-indicator"),
				collapsible.CollapseCheckbox(
					collapsible.CollapseCheckboxProps{},
				),
				collapsible.CollapseTitle(
					collapsible.CollapseTitleProps{
						ClassNames: htmx.ClassNames{
							tailwind.Flex:        true,
							tailwind.ItemsCenter: true,
						},
					},
					htmx.Text("Revisions"),
					loading.SpinnerSmall(
						loading.SpinnerProps{
							ClassNames: htmx.ClassNames{
								tailwind.Mx2:     true,
								"htmx-indicator": true,
							},
						},
					),
				),
				collapsible.CollapseContent(
					collapsible.CollapseContentProps{},
					htmx.Div(
						htmx.Text("Loading..."),
					),
				),
			),
		),
	)
}
