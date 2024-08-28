package designs

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/collapsible"
	"github.com/zeiss/fiber-htmx/components/tailwind"
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
				// htmx.HxTrigger("click once"),
				// htmx.HxGet("/designs/"+props.DesignID.String()+"/revisions"),
				collapsible.CollapseCheckbox(
					collapsible.CollapseCheckboxProps{},
				),
				collapsible.CollapseTitle(
					collapsible.CollapseTitleProps{},
					htmx.Text("Revisions"),
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
