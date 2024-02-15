package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Lens is a model for a lens.
type Lens struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Version     int        `json:"version"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Pillars     []Pillar   `json:"pillars"`
	Resources   []Resource `json:"resources"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Pillar is a model for a pillar.
type Pillar struct {
	ID          int        `json:"id" gorm:"primary_key"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
	Resources   []Resource `json:"resources"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Question is a model for a question.
type Question struct {
	ID          int        `json:"id" gorm:"type:uuid;primary_key;"`
	Ref         string     `json:"ref"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Resources   []Resource `json:"resources"`
	Choices     []Choice   `json:"choices"`
	Risks       []Risk     `json:"risks"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Resource is a model for a resource.
type Resource struct {
	ID          int    `json:"id" gorm:"primary_key"`
	URL         string `json:"url"`
	Description string `json:"description"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Choice is a model for a choice.
type Choice struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Ref         string `json:"ref"`
	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Risk is a model for a risk.
type Risk struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Ref       string `json:"ref"`
	Risk      string `json:"risk"`
	Condition string `json:"condition"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// UnmarshalJSON ...
func (l *Lens) UnmarshalJSON(data []byte) error {
	lens := struct {
		Version     int        `json:"version"`
		Name        string     `json:"name"`
		Description string     `json:"description"`
		Pillars     []Pillar   `json:"pillars"`
		Resources   []Resource `json:"resources"`
	}{}

	if err := json.Unmarshal(data, &lens); err != nil {
		return errors.WithStack(err)
	}

	l.Version = lens.Version
	l.Name = lens.Name
	l.Description = lens.Description
	l.Pillars = lens.Pillars
	l.Resources = lens.Resources

	return nil
}
