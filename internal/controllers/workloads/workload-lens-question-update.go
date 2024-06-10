package workloads

import (
	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadLensQuestionUpdateControllerParams ...
type WorkloadLensQuestionUpdateControllerParams struct {
	ID       uuid.UUID `json:"id" xml:"id" form:"id"`
	Team     string    `json:"team" xml:"team" form:"team"`
	Lens     uuid.UUID `json:"lens" xml:"lens" form:"lens"`
	Question int       `json:"question" xml:"question" form:"question"`
}

// WorkloadLensQuestionUpdateControllerBody ...
type WorkloadLensQuestionUpdateControllerBody struct {
	Choices      []int  `json:"choices" xml:"choices" form:"choices"`
	DoesNotApply bool   `json:"does_not_apply" xml:"does_not_apply" form:"does_not_apply"`
	Notes        string `json:"notes" xml:"notes" form:"notes"`
}

// WorkloadLensQuestionUpdateController ...
type WorkloadLensQuestionUpdateController struct {
	db     ports.Repository
	params *WorkloadLensQuestionUpdateControllerParams
	body   *WorkloadLensQuestionUpdateControllerBody

	htmx.UnimplementedController
}

// WorkloadLensQuestionUpdateController ...
func NewWorkloadLensQuestionUpdateController(db ports.Repository) *WorkloadLensQuestionUpdateController {
	return &WorkloadLensQuestionUpdateController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadLensQuestionUpdateController) Prepare() error {
	params := &WorkloadLensQuestionUpdateControllerParams{}
	if err := w.BindParams(params); err != nil {
		return err
	}
	w.params = params

	body := &WorkloadLensQuestionUpdateControllerBody{}
	if err := w.BindBody(body); err != nil {
		return err
	}
	w.body = body

	return nil
}

// Post ...
func (w *WorkloadLensQuestionUpdateController) Post() error {
	err := w.db.UpdateAnswers(w.Context(), w.params.ID, w.params.Lens, w.params.Question, w.body.Choices, w.body.DoesNotApply, w.body.Notes)
	if err != nil {
		return err
	}

	return nil
}
