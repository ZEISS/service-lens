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

	Answers []*WorkloadLensQuestionAnswer `json:"answers" gorm:"foreignKey:WorkloadID;"`

	// Tags are the attached tags
	Tags []*Tag `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:workload;"`
	// ReviewOwner is the owner of the review.
	ReviewOwner string `json:"review_owner" form:"review_owner" validate:"required,email"`
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
