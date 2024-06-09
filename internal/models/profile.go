package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Profile represents a business profile.
type Profile struct {
	// ID ...
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" query:"id" param:"id"`
	// Name of the profile.
	Name string `json:"name" gorm:"uniqueIndex:idx_profile_name_team" form:"name" validate:"required,min=3,max=100"`
	// Description of the profile.
	Description string `json:"description" form:"description" validate:"required,min=3,max=1024"`
	// Questions ...
	Questions []ProfileQuestion `json:"questions" gorm:"many2many:profiles_questions;" form:"questions"`

	Tags []*Tag `json:"tags" gorm:"polymorphic:Taggable;"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Question represents a business profile question.
type ProfileQuestion struct {
	ID            int    `json:"id" gorm:"primary_key"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	MulipleChoice bool   `json:"multiple_choice"`

	Choices   []ProfileQuestionChoice `json:"choices"`
	ProfileID uuid.UUID               `json:"pillar_id"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// ProfileQuestionChoice is a model for a choice.
type ProfileQuestionChoice struct {
	ID          int         `json:"id" gorm:"primary_key"`
	Ref         QuestionRef `json:"ref"`
	Title       string      `json:"title"`
	Description string      `json:"description"`

	ProfileQuestion   ProfileQuestion `json:"profile_question" gorm:"foreignkey:ProfileQuestionID;"`
	ProfileQuestionID int             `json:"profile_question_id"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Answer represents a business profile question answer.
type ProfileQuestionAnswer struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Ref         string `json:"ref"`
	Tile        string `json:"tile"`
	Description string `json:"description"`

	Question   ProfileQuestion `json:"question" gorm:"foreignkey:QuestionID;"`
	QuestionID int             `json:"question_id"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
