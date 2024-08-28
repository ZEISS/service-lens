package models

import (
	"slices"
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
	// CreatedAt is the created at field
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the updated at field
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the soft delete field
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// GetStates returns the states of the workflow
func (w Workflow) GetStates() []WorkflowState {
	slices.SortFunc(w.States, func(i, j WorkflowState) int {
		return i.Order - j.Order
	})

	return w.States
}

// WorkflowState is the model for the workflow_state table
type WorkflowState struct {
	// ID is the primary key of the workflow status
	ID int `json:"id" gorm:"primary_key;type:bigint;auto_increment;not null"`
	// WorkflowID is the foreign key of the workflow
	WorkflowID uuid.UUID `json:"workflow_id" gorm:"type:uuid;idx_workflow_state"`
	// Name is the name of the workflow status
	Name string `json:"name" gorm:"type:varchar(255);unique_index:idx_workflow_state"`
	// Description is the description of the workflow status
	Description string `json:"description" gorm:"type:text"`
	// Order is the order of the state in the workflow
	Order int `json:"order" gorm:"type:int"`
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
	WorkflowID uuid.UUID `json:"workflow_id" gorm:"type:uuid;not null"`
	// Workflow is the workflow of the transition
	Workflow Workflow `json:"workflow" gorm:"foreignKey:WorkflowID;references:ID"`
	// CurrentStateID is the foreign key of the current state
	CurrentStateID int `json:"current_state_id" gorm:"type:bigint;not null"`
	// CurrentState is the current state of the transition
	CurrentState WorkflowState `json:"current_state" gorm:"foreignKey:CurrentStateID;references:ID"`
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
	// WorkflowTransitionID is the foreign key of the workflow transition
	WorkflowTransitionID int `json:"workflow_transition_id" gorm:"type:bigint"`
	// WorkflowTransition is the workflow transition of the workable
	WorkflowTransition WorkflowTransition `json:"workflow_transition" gorm:"foreignKey:WorkflowTransitionID;references:ID"`
}
