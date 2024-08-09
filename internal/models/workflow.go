package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkableType is the type of the workable
type WorkableType string

const (
	// WorkableTypeDesign is the design workable type
	WorkableTypeDesign WorkableType = "design"
)

// Workflow a workflow that can be used to manage the state of a model
type Workflow struct {
	// ID ...
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the workflow
	Name string `json:"name" gorm:"type:varchar(255);unique_index"`
	// Description is the description of the workflow
	Description string `json:"description" gorm:"type:text"`
	// States is the list of states in the workflow
	States []WorkflowState `json:"states" gorm:"foreignKey:WorkflowID;references:ID"`
	// Transitions is the list of transitions in the workflow
	Transitions []WorkflowTransition `json:"transitions" gorm:"foreignKey:WorkflowID;references:ID"`
	// CreatedAt is the created at field
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the updated at field
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the soft delete field
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// WorkflowState is the model for the workflow_state table
type WorkflowState struct {
	// ID is the primary key of the workflow status
	ID int `json:"id" gorm:"primary_key;type:bigint;auto_increment;not null"`
	// WorkflowID is the foreign key of the workflow
	WorkflowID int `json:"workflow_id" gorm:"primary_key;type:bigint;not null"`
	// Name is the name of the workflow status
	Name string `json:"name" gorm:"type:varchar(255)"`
	// Description is the description of the workflow status
	Description string `json:"description" gorm:"type:text"`
	// CreatedAt is the created at field
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the updated at field
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the soft delete field
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// WorkflowTransition is the model for the workflow_transition table
type WorkflowTransition struct {
	// ID is the primary key of the workflow transition
	ID int `json:"id" gorm:"primary_key;type:bigint;auto_increment;not null"`
	// WorkflowID is the foreign key of the workflow
	WorkflowID int `json:"workflow_id" gorm:"type:bigint;not null"`
	// CurrentStateId is the foreign key of the current state
	CurrentStateId int `json:"current_state_id" gorm:"type:bigint;not null"`
	// NextStateID is the foreign key of the next state
	NextStateID int `json:"next_state_id" gorm:"type:bigint;not null"`
	// CreatedAt is the created at field
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the updated at field
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the soft delete field
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Workable is model that can be used with a workflow
type Workable struct {
	// ID is the primary key of the workable
	ID int `json:"id" gorm:"primary_key;type:bigint;auto_increment;not null"`
	// WorkableID is the foreign key of the workable
	WorkableID uuid.UUID `json:"workable_id" gorm:"type:uuid;not null"`
	// WorkableType is the type of the workable
	WorkableType WorkableType `json:"workable_type" gorm:"type:varchar(255);not null"`
	// WorkflowID is the foreign key of the workflow
	WorkflowID int `json:"workflow_id" gorm:"type:bigint;not null"`
	// Workflow is the workflow of the workable
	Workflow Workflow `json:"workflow" gorm:"foreignKey:WorkflowID;references:ID"`
}
