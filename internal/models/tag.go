package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

type TaggableType string

// TaggableType ...
const (
	LensType       TaggableType = "Lens"
	ProfileTyp     TaggableType = "Profile"
	WorkloadTyp    TaggableType = "Workload"
	EnvironmentTyp TaggableType = "Environment"
)

// Tag ...
type Tag struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key"`
	// Name is the tag name.
	Name string `json:"name"`
	// TaggableID is the foreign key of the taggable
	TaggableID uuid.UUID `json:"taggable_id"`
	// TaggableType is the type of the taggable
	TaggableType TaggableType `json:"taggable_type"`
	// Team is the team that owns the environment
	Team adapters.GothTeam `json:"owner" gorm:"foreignKey:TeamID"`
	// TeamID is the foreign key of the owner
	TeamID uuid.UUID `json:"owner_id" gorm:"type:uuid;index"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
