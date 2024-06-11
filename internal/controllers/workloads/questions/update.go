package questions

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

// QuestionForm ...
type QuestionForm struct {
	Choices            []string `form:"choices"`
	DoesNotApply       bool     `form:"does_not_apply"`
	DoesNotApplyReason string   `form:"does_not_apply_reason"`
}

var validate *validator.Validate

// WorkloadUpdateAnswerControllerImpl ...
type WorkloadUpdateAnswerControllerImpl struct {
	answer models.WorkloadLensQuestionAnswer
	store  ports.Datastore
	htmx.DefaultController
}

// NewWorkloadUpdateAnswerController ...
func NewWorkloadUpdateAnswerController(store ports.Datastore) *WorkloadUpdateAnswerControllerImpl {
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

	err := w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		err := w.BindParams(&w.answer)
		if err != nil {
			return err
		}

		return tx.GetWorkloadAnswer(ctx, &w.answer)
	})
	if err != nil {
		return err
	}

	var form QuestionForm
	err = w.BindBody(&form)
	if err != nil {
		return err
	}

	err = validate.Struct(&form)
	if err != nil {
		return err
	}
	w.answer.DoesNotApply = form.DoesNotApply
	w.answer.DoesNotApplyReason = form.DoesNotApplyReason

	for _, choice := range form.Choices {
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
