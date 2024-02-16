package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Team ...
type Team struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`

	Profiles []*Profile `json:"profiles" gorm:"many2many:profiles_teams;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
