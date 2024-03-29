package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaggableType string

// TaggableType ...
const (
	LensType    TaggableType = "Lens"
	ProfileTyp  TaggableType = "Profile"
	WorkloadTyp TaggableType = "Workload"
)

// Tag ...
type Tag struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	TaggableID   uuid.UUID    `json:"taggable_id"`
	TaggableType TaggableType `json:"taggable_type"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
