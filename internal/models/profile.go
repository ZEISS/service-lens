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

// ProfileQuestion represents a business profile questions.
type ProfileQuestions struct {
	ProfileID  uuid.UUID `json:"profile_id" gorm:"primary_key"`
	QuestionID int       `json:"question_id" gorm:"primary_key"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// TableName ...
func (ProfileQuestions) TableName() string {
	return "profiles_questions"
}

// Question represents a business profile question.
type ProfileQuestion struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`

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
