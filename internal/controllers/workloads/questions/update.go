package questions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/expr-lang/expr"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/utilx"
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

// Prepare ...
// nolint:gocyclo
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

	if !w.form.DoesNotApply {
		choices := make(map[int]bool)
		for _, c := range w.form.Choices {
			choices[conv.Int(c)] = true
		}

		question := models.Question{ID: w.params.QuestionID}
		err = w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
			return tx.GetLensQuestion(ctx, &question)
		})
		if err != nil {
			return err
		}

		env := map[string]bool{
			"default": true,
		}

		for _, c := range question.Choices {
			_, ok := choices[c.ID]
			env[string(c.Ref)] = utilx.IfElse(ok, true, false)
		}

		for _, r := range question.Risks {
			rule := r.Condition

			program, err := expr.Compile(rule, expr.Env(env))
			if err != nil {
				return err
			}

			output, err := expr.Run(program, env)
			if err != nil {
				return err
			}

			v, ok := output.(bool)
			if !ok {
				return fmt.Errorf("expected bool, got %T", v)
			}

			w.answer.RiskID = cast.Ptr(r.ID)

			if v {
				break
			}
		}
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
			ID: conv.Int(choice),
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
