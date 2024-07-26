package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Environment ...
type Environment struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// Name of the environment
	Name string `json:"name" form:"name" validate:"required,min=3,max=255"`
	// Description of the environment
	Description string `json:"description" form:"description" validate:"max=1024"`
	// Tags are the tags associated with the environment
	Tags []Tag `json:"tags" gorm:"many2many:environment_tags;"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}
