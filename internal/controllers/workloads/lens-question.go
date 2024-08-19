package workloads

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/collapsible"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

const (
	updateWorkloadAnswerURL = "/workloads/%s/lenses/%s/question/%d"
)

// LensQuestionParams
type LensQuestionParams struct {
	QuestionID int       `params:"question"`
	WorkloadID uuid.UUID `params:"workload"`
	LensID     uuid.UUID `params:"lens"`
}

// WorkloadLensEditQuestionControllerImpl ...
type WorkloadLensEditQuestionControllerImpl struct {
	params   LensQuestionParams
	question models.Question
	answer   models.WorkloadLensQuestionAnswer
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadLensEditQuestionController ...
func NewWorkloadLensEditQuestionController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadLensEditQuestionControllerImpl {
	return &WorkloadLensEditQuestionControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *WorkloadLensEditQuestionControllerImpl) Prepare() error {
	err := w.BindParams(&w.params)
	if err != nil {
		return err
	}
	w.answer.LensID = w.params.LensID
	w.answer.WorkloadID = w.params.WorkloadID
	w.answer.QuestionID = w.params.QuestionID
	w.question.ID = w.params.QuestionID

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		err := tx.GetLensQuestion(ctx, &w.question)
		if err != nil {
			return err
		}

		return tx.GetWorkloadAnswer(ctx, &w.answer)
	})
}

// Get ...
func (w *WorkloadLensEditQuestionControllerImpl) Get() error {
	return w.Render(
		htmx.Form(
			htmx.HxPut(fmt.Sprintf(updateWorkloadAnswerURL, w.params.WorkloadID, w.params.LensID, w.params.QuestionID)),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text(w.question.Title),
					),
				),
			),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					collapsible.CollapseArrow(
						collapsible.CollapseProps{},
						collapsible.CollapseCheckbox(
							collapsible.CollapseCheckboxProps{},
						),
						collapsible.CollapseTitle(
							collapsible.CollapseTitleProps{},
							htmx.Text("Additional Information"),
						),
						collapsible.CollapseContent(
							collapsible.CollapseContentProps{},
							htmx.Text(w.question.Description),
						),
					),
				),
			),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{
								"flex":            true,
								"justify-between": true,
								"flex-row":        true,
							},
						},
						forms.FormControlLabel(
							forms.FormControlLabelProps{},
							forms.FormControlLabelText(
								forms.FormControlLabelTextProps{},
								htmx.Text("Question does not apply to this workload"),
							),
						),
						forms.Toggle(
							forms.ToggleProps{
								Name:    "does_not_apply",
								Value:   "true",
								Checked: w.answer.DoesNotApply,
							},
							htmx.HyperScript(`on change if me.checked set disabled of <input[type=checkbox][name=choices]/> to true
								remove .hidden from next <label/>
								else set disabled of <input[type=checkbox][name=choices]/> to false
								add .hidden to next <label/>`),
						),
					),
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{
								"hidden": !w.answer.DoesNotApply,
							},
						},
						htmx.ID("does-not-apply-reason"),
						forms.SelectBordered(
							forms.SelectProps{
								ClassNames: htmx.ClassNames{
									"w-full": true,
								},
							},
							forms.Option(
								forms.OptionProps{
									Selected: w.answer.DoesNotApplyReason == "",
									Disabled: true,
								},
								htmx.Text("Select a reason"),
							),
							forms.Option(
								forms.OptionProps{
									Selected: w.answer.DoesNotApplyReason == "OUT_OF_SCOPE",
								},
								htmx.Text("Out of scope"),
								htmx.Value("OUT_OF_SCOPE"),
							),
							forms.Option(
								forms.OptionProps{
									Selected: w.answer.DoesNotApplyReason == "BUSINESS_PRIORITIES",
								},
								htmx.Text("Business Priorities"),
								htmx.Value("BUSINESS_PRIORITIES"),
							),
							forms.Option(
								forms.OptionProps{
									Selected: w.answer.DoesNotApplyReason == "ARCHITECTURE_CONSTRAINTS",
								},
								htmx.Text("Architecture Constraints"),
								htmx.Value("ARCHITECTURE_CONSTRAINTS"),
							),
							forms.Option(
								forms.OptionProps{
									Selected: w.answer.DoesNotApplyReason == "OTHER",
								},
								htmx.Text("Other"),
								htmx.Value("OTHER"),
							),
							htmx.Name("does_not_apply_reason"),
						),
					),
					htmx.Group(htmx.ForEach(w.question.Choices, func(choice models.Choice, choiceIdx int) htmx.Node {
						return forms.FormControl(
							forms.FormControlProps{},
							forms.FormControlLabel(
								forms.FormControlLabelProps{},
								forms.FormControlLabelText(
									forms.FormControlLabelTextProps{},
									htmx.Text(choice.Title),
								),
								forms.Checkbox(
									forms.CheckboxProps{
										Name:     "choices",
										Value:    utils.IntStr(choice.ID),
										Checked:  w.answer.IsChecked(choice.ID), // todo(katallaxie): should be a default option in the model
										Disabled: w.answer.DoesNotApply || (choice.Ref == models.NoneOfTheseQuestionRef && w.answer.DoesNotApply),
									},
								),
							),
						)
					})...),
				),
			),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					forms.FormControl(
						forms.FormControlProps{},
						forms.FormControlLabel(
							forms.FormControlLabelProps{},
							forms.FormControlLabelText(
								forms.FormControlLabelTextProps{},
								htmx.Text("Notes"),
							),
						),
						forms.TextareaBordered(
							forms.TextareaProps{
								Name:        "notes",
								Placeholder: "Optional notes",
							},
							htmx.Text(w.answer.Notes),
						),
						forms.FormControlLabelAltText(
							forms.FormControlLabelAltTextProps{},
							htmx.Text("Optional - Can be from 3 to 2048 characters."),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Attribute("type", "submit"),
							htmx.Text("Save & Next"),
						),
					),
				),
			),
		),
	)
}
