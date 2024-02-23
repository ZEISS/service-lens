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

	Tags   []*Tag     `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:workload;"`
	Team   authz.Team `json:"team" gorm:"foreignKey:TeamID;"`
	TeamID uuid.UUID  `json:"team_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
