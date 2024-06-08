package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Workload ...
type Workload struct {
	ID            uuid.UUID   `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	ProfileID     uuid.UUID   `json:"profile_id"`
	Profile       Profile     `json:"profile"`
	Lenses        []*Lens     `json:"lenses" gorm:"many2many:workload_lenses;"`
	EnvironmentID uuid.UUID   `json:"environment_id"`
	Environment   Environment `json:"environment"`

	Answers []*WorkloadLensQuestionAnswer `json:"answers" gorm:"foreignKey:WorkloadID;"`

	Tags []*Tag `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:workload;"`

	ReviewOwner string `json:"review_owner" validate:"required,min=1,max=255"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// WorkloadLensQuestionAnswer represents a business profile question answer.
type WorkloadLensQuestionAnswer struct {
	ID                 int    `json:"id" gorm:"primary_key"`
	DoesNotApply       bool   `json:"does_not_apply"`
	DoesNotApplyReason string `json:"does_not_apply_reason"`
	Notes              string `json:"notes"`

	QuestionID int      `json:"question_id" gorm:"uniqueIndex:idx_workload_lens_question_answer;"`
	Question   Question `json:"question"`

	LensID uuid.UUID `json:"lens_id" gorm:"uniqueIndex:idx_workload_lens_question_answer;"`
	Lens   Lens      `json:"lens"`

	WorkloadID uuid.UUID `json:"workload_id" gorm:"uniqueIndex:idx_workload_lens_question_answer;"`
	Workload   Workload  `json:"workload" `

	Choices []*Choice `json:"choices" gorm:"many2many:workload_lens_question_answer_choices;"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}
