package models

import "github.com/google/uuid"

// OwnableType ...
type OwnableType string

// OwnableTypeTemplate is the template ownable type
var OwnableTypeTemplate OwnableType = "template"

// OwnerType ...
type OwnerType string

var (
	// OwnerTypeUser is the user owner type
	OwnerTypeUser OwnerType = "user"
	// OwnerTypeTeam is the team owner type
	OwnerTypeTeam OwnerType = "team"
)

// Ownable is the interface for the ownable
type Ownable struct {
	// ID is the primary key of the workable
	ID int `json:"id" gorm:"primary_key;type:bigint;auto_increment;not null"`
	// OwnableID  is the foreign key of the ownable
	OwnableID uuid.UUID `json:"ownable_id" gorm:"type:uuid;not null"`
	// OwnableType is the type of the ownable
	OwnableType OwnableType `json:"ownable_type" gorm:"type:varchar(255);not null"`
	// OwnerID is the foreign key of the owner
	OwnerID uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	// OwnerType is the type of the owner
	OwnerType OwnerType `json:"owner_type" gorm:"type:varchar(255);not null"`
}
