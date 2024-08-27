package workflows

import (
	"context"

	"github.com/zeiss/service-lens/internal/components/workflows"
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

// Delete ...
func (l *StepControllerImpl) Delete() error {
	var params struct {
		WorkflowID uuid.UUID `json:"workflow_id" params:"id"`
		ID         int       `json:"id" params:"step_id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteWorkflowState(ctx, &models.WorkflowState{ID: params.ID, WorkflowID: params.WorkflowID})
	})
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
		WorkflowID:  params.WorkflowID,
		Name:        params.Name,
		Description: params.Description,
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateWorkflowState(ctx, &state)
	})
	if err != nil {
		return err
	}

	return l.Render(
		htmx.Fragment(
			workflows.WorkflowStep(
				workflows.WorkflowStepProps{
					State:      state,
					WorkflowID: state.WorkflowID,
				},
				htmx.ID("steps"),
				htmx.HxSwapOob("beforeend focus-scroll:true"),
			),
		),
	)
}
