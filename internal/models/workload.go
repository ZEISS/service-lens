package models

import (
	"time"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"gorm.io/gorm"
)

// Workload ...
type Workload struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProfileID   uuid.UUID `json:"profile_id"`
	Profile     Profile   `json:"profile"`
	Lenses      []*Lens   `json:"lenses" gorm:"many2many:workload_lenses;"`

	Answers []*WorkloadLensQuestionAnswer `json:"answers" gorm:"foreignKey:WorkloadID;"`

	Tags   []*Tag     `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:workload;"`
	Team   authz.Team `json:"team" gorm:"foreignKey:TeamID;"`
	TeamID uuid.UUID  `json:"team_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// WorkloadLensQuestionAnswer represents a business profile question answer.
type WorkloadLensQuestionAnswer struct {
	ID                 int    `json:"id" gorm:"primary_key"`
	DoesNotApply       bool   `json:"does_not_apply"`
	DoesNotApplyReason string `json:"does_not_apply_reason"`
	Notes              string `json:"notes"`

	Question   Question `json:"question" gorm:"foreignkey:QuestionID;"`
	QuestionID int      `json:"question_id"`

	LensID uuid.UUID `json:"lens_id"`
	Lens   Lens      `json:"lens"`

	WorkloadID uuid.UUID `json:"workload_id"`
	Workload   Workload  `json:"workload"`

	Choices []*Choice `json:"choices" gorm:"many2many:workload_lens_question_answer_choices;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
