package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Team is a group of users. Teams can be used to group lenses, reviews, profiles.
type Team struct {
	// ID is the primary key of the team.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" params:"id" validate:"require,uuid"`
	// Name is the name of the team.
	Name string `json:"name" form:"name" validate:"required,alphanum,gt=3,lt=255"`
	// Slug is the unique identifier of the team.
	Slug string `json:"slug" gorm:"uniqueIndex" form:"slug"  validate:"required,alphanum,gt=3,lt=255,lowercase"`
	// Description is the description of the team.
	Description *string `json:"description" form:"description" validate:"omitempty,max=255"`
	// CreatedAt is the time the team was created.
	CreatedAt time.Time
	// UpdatedAt is the time the team was last updated.
	UpdatedAt time.Time
	// DeletedAt is the time the team was deleted.
	DeletedAt gorm.DeletedAt
}