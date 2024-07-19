package questions

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

// QuestionUpdateForm ...
type QuestionUpdateForm struct {
	Choices            []string `form:"choices"`
	Notes              string   `form:"notes"`
	DoesNotApply       bool     `form:"does_not_apply"`
	DoesNotApplyReason string   `form:"does_not_apply_reason"`
}

// QuestionUpdateParams ...
type QuestionUpdateParams struct {
	QuestionID int       `params:"question"`
	WorkloadID uuid.UUID `params:"workload"`
	LensID     uuid.UUID `params:"lens"`
}

var validate *validator.Validate

// WorkloadUpdateAnswerControllerImpl ...
type WorkloadUpdateAnswerControllerImpl struct {
	params QuestionUpdateParams
	form   QuestionUpdateForm
	answer models.WorkloadLensQuestionAnswer
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadUpdateAnswerController ...
func NewWorkloadUpdateAnswerController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadUpdateAnswerControllerImpl {
	return &WorkloadUpdateAnswerControllerImpl{
		store: store,
	}
}

// Error ...
func (w *WorkloadUpdateAnswerControllerImpl) Error(err error) error {
	return err
}

// Prepare ...
func (w *WorkloadUpdateAnswerControllerImpl) Prepare() error {
	validate = validator.New()

	err := w.BindParams(&w.params)
	if err != nil {
		return err
	}
	w.answer.WorkloadID = w.params.WorkloadID
	w.answer.LensID = w.params.LensID
	w.answer.QuestionID = w.params.QuestionID

	err = w.BindBody(&w.form)
	if err != nil {
		return err
	}

	err = validate.Struct(&w.form)
	if err != nil {
		return err
	}
	w.answer.DoesNotApply = w.form.DoesNotApply
	w.answer.DoesNotApplyReason = w.form.DoesNotApplyReason
	w.answer.Notes = w.form.Notes

	for _, choice := range w.form.Choices {
		w.answer.Choices = append(w.answer.Choices, models.Choice{
			ID: utils.StrInt(choice),
		})
	}

	return w.store.ReadWriteTx(w.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateWorkloadAnswer(ctx, &w.answer)
	})
}

// Put ...
func (w *WorkloadUpdateAnswerControllerImpl) Put() error {
	return w.Ctx().SendStatus(http.StatusNoContent)
}
