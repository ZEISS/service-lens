package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Workload ...
type Workload struct {
	// ID ...
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name adds a name to the workload.
	Name string `json:"name" form:"name" validate:"required,min=1,max=255"`
	// Description adds a description to the workload.
	Description string `json:"description" form:"description" validate:"required,min=3,max=1024"`
	// ProfileID is the foreign key for the profile.
	ProfileID uuid.UUID `json:"profile_id" form:"profile_id" validate:"required,uuid"`
	Profile   Profile   `json:"profile" validate:"-"`
	// Lenses are the attached lenses
	Lenses []*Lens `json:"lenses" gorm:"many2many:workload_lenses;"`
	// EnvironmentID is the foreign key for the environment.
	EnvironmentID uuid.UUID   `json:"environment_id" form:"environment_id" validate:"required,uuid"`
	Environment   Environment `json:"environment" validate:"-"`
	// Answers are the attached answers
	Answers []*WorkloadLensQuestionAnswer `json:"answers" gorm:"foreignKey:WorkloadID;" form:"answers"`
	// Tags are the attached tags
	Tags []Tag `json:"tags" gorm:"many2many:workload_tags;"`
	// ReviewOwner is the owner of the review.
	ReviewOwner string `json:"review_owner" form:"review_owner" validate:"required,email"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// TotalAnswers ...
func (w *Workload) TotalAnswers() int {
	return len(w.Answers)
}

// TotalQuestions ...
func (w *Workload) TotalQuestions() int {
	total := 0

	for _, l := range w.Lenses {
		for _, p := range l.Pillars {
			total += len(p.Questions)
		}
	}

	return total
}

// TotalHighRisks ...
func (w *Workload) TotalHighRisks() int {
	total := 0

	for _, a := range w.Answers {
		if a.Risk != nil && a.Risk.Risk == "HIGH_RISK" {
			total++
		}
	}

	return total
}

// TotalMediumRisks ...
func (w *Workload) TotalMediumRisks() int {
	total := 0

	for _, a := range w.Answers {
		if a.Risk != nil && a.Risk.Risk == "MEDIUM_RISK" {
			total++
		}
	}

	return total
}

// TotalLenses ...
func (w *Workload) TotalLenses() int {
	return len(w.Lenses)
}

// WorkloadLensQuestionAnswer represents a business profile question answer.
type WorkloadLensQuestionAnswer struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key"`
	// DoesNotApply is a flag to indicate if the question does not apply.
	DoesNotApply bool `json:"does_not_apply" form:"does_not_apply"`
	// DoesNotApplyReason is the reason the question does not apply.
	DoesNotApplyReason string `json:"does_not_apply_reason" form:"does_not_apply_reason" validate:"max=1024"`
	// Notes are additional notes.
	Notes string `json:"notes" form:"notes" validate:"max=2048"`
	// QuestionID is the foreign key for the question.
	QuestionID int `json:"question_id" gorm:"uniqueIndex:idx_workload_lens_question_answer;" params:"question" form:"question_id" validate:"required"`
	// Question is the question.
	Question Question `json:"question"`
	// LensID is the foreign key for the lens.
	LensID uuid.UUID `json:"lens_id" gorm:"uniqueIndex:idx_workload_lens_question_answer;" params:"lens" form:"lens_id" validate:"required,uuid"`
	// Lens is the lens.
	Lens Lens `json:"lens"`
	// WorkloadID is the foreign key for the workload.
	WorkloadID uuid.UUID `json:"workload_id" gorm:"uniqueIndex:idx_workload_lens_question_answer;" params:"workload" form:"workload_id" validate:"required,uuid"`
	// Workload is the workload.
	Workload Workload `json:"workload" validate:"-"`
	// Choices are the selected choices.
	Choices []Choice `json:"choices" gorm:"many2many:workload_lens_question_answer_choices;" form:"choices"`
	// RiskID is the foreign key for the risk.
	RiskID *int `json:"risk_id" form:"risk_id" validate:"required"`
	// Risk is the risk associated with the answer.
	Risk *Risk `json:"risk" form:"risk"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// IsChecked ...
func (a *WorkloadLensQuestionAnswer) IsChecked(choiceID int) bool {
	for _, c := range a.Choices {
		if c.ID == choiceID {
			return true
		}
	}
	return false
}
