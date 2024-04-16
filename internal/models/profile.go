package models

import (
	"time"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"gorm.io/gorm"
)

// Profile represents a business profile.
type Profile struct {
	ID          uuid.UUID         `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string            `json:"name" gorm:"uniqueIndex:idx_profile_name_team"`
	Description string            `json:"description"`
	Questions   []ProfileQuestion `json:"questions" gorm:"many2many:profiles_questions;"`

	Tags   []*Tag     `json:"tags" gorm:"polymorphic:Taggable;"`
	Team   authz.Team `json:"team" gorm:"foreignKey:TeamID;"`
	TeamID uuid.UUID  `json:"team_id" gorm:"uniqueIndex:idx_profile_name_team"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
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

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
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

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
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

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
