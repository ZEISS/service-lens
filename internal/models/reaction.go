package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

// ReactableType ...
type ReactableType string

const (
	// ReactableTypeDesign ...
	ReactableTypeDesign ReactableType = "design"
	// ReactableTypeDesignComment ...
	ReactableTypeDesignComment ReactableType = "design_comment"
)

// Reaction ...
type Reaction struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"type:bigint;primaryKey;unique;autoIncrement"`
	// ReactableID is the ID of the reactable
	ReactableID uuid.UUID `json:"reactable_id" unique_index:"idx_reactable_id_reactable_type_reactor_id"`
	// ReactableType is the type of the taggable
	ReactableType ReactableType `json:"reactable_type" unique_index:"idx_reactable_id_reactable_type_reactor_id"`
	// Value is the value of the reaction
	Value string `json:"value"`
	// ReactorID is the ID of the reactor
	ReactorID uuid.UUID `json:"reactor_id" unique_index:"idx_reactable_id_reactable_type_reactor_id"`
	// Reactor is the reactor
	Reactor adapters.GothUser `json:"reactor" gorm:"foreignKey:ReactorID;references:ID"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
