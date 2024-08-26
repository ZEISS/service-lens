package workflows

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// StepControllerImpl ...
type StepControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewStepController ...
func NewStepController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *StepControllerImpl {
	return &StepControllerImpl{store: store}
}

// Error ...
func (l *StepControllerImpl) Error(err error) error {
	fmt.Println(err)

	return err
}

// Post ...
func (l *StepControllerImpl) Post() error {
	var params struct {
		WorkflowID  uuid.UUID `json:"workflow_id" params:"id"`
		Name        string    `json:"name" forms:"name"`
		Description string    `json:"description" forms:"description"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	err = l.BindBody(&params)
	if err != nil {
		return err
	}

	state := models.WorkflowState{
		WorkflowID: params.WorkflowID,
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateWorkflowState(ctx, &state)
	})
	if err != nil {
		return err
	}

	return l.Render(
		htmx.Fragment(
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				htmx.ID("steps"),
				htmx.HxSwapOob("beforeend"),
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text(state.Name),
					),
					htmx.Text(state.Description),
				),
			),
		),
	)
}
