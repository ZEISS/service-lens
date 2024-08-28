package workflows

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// WorkflowStepProps ...
type WorkflowStepProps struct {
	ClassNames htmx.ClassNames
	State      models.WorkflowState
	WorkflowID uuid.UUID
}

// WorkflowStep ...
func WorkflowStep(props WorkflowStepProps, children ...htmx.Node) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.CursorPointer: true,
				},
				props.ClassNames,
			),
		},
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("step"),
			htmx.Value(conv.String(props.State.ID)),
		),
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text(props.State.Name),
			),
			htmx.Text(props.State.Description),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxDelete(fmt.Sprintf(utils.DeleteWorkflowStepUrlFormat, conv.String(props.WorkflowID), props.State.ID)),
					htmx.HxConfirm("Are you sure you want to delete this step?"),
					htmx.HxTarget("closest .card"),
					htmx.HxSwap("outerHTML swap:1s"),
					htmx.Text("Delete"),
				),
			),
		),
		htmx.Group(children...),
	)
}
