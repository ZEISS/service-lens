package models

import (
	"time"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"gorm.io/gorm"
)

// Environment ...
type Environment struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"uniqueIndex:idx_environment_name_team"`
	Description string    `json:"description"`

	Tags   []*Tag     `json:"tags" gorm:"polymorphic:Taggable;"`
	Team   authz.Team `json:"team" gorm:"foreignKey:TeamID;"`
	TeamID uuid.UUID  `json:"team_id" gorm:"uniqueIndex:idx_environment_name_team"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
