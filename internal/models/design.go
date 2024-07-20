package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Design ...
type Design struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// Title of the design
	Title string `json:"title" form:"title" validate:"required,min=3,max=255"`
	// Tags are the tags associated with the environment
	Tags []*Tag `json:"tags" gorm:"polymorphic:Taggable;"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}
