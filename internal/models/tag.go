package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaggableType string

// TaggableType ...
const (
	DesignType     TaggableType = "Design"
	EnvironmentTyp TaggableType = "Environment"
	LensType       TaggableType = "Lens"
	ProfileTyp     TaggableType = "Profile"
	WorkloadTyp    TaggableType = "Workload"
)

// Tag ...
type Tag struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key"`
	// Name is the tag name.
	Name string `json:"name"`
	// Value is the tag value.
	Value string `json:"value"`
	// TaggableID is the foreign key of the taggable
	TaggableID uuid.UUID `json:"taggable_id"`
	// TaggableType is the type of the taggable
	TaggableType TaggableType `json:"taggable_type"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
