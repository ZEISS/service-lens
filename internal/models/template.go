package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Template ...
type Template struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// Name of the template
	Name string `json:"name" form:"name" validate:"required,min=3,max=1024"`
	// Description of the template
	Description string `json:"description" form:"description" gorm:"type:text" validate:"required,min=3"`
	// Body of the template in markdown, HTML, or plain text
	Body string `json:"body" gorm:"type:text"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// NewTemplate ...
func NewTemplate() *Template {
	return &Template{}
}
