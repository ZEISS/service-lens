package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// QuestionRef is a question reference.
type QuestionRef string

// String is a string representation of a question reference.
func (q QuestionRef) String() string {
	return string(q)
}

const (
	// NoneOfTheseQuestionRef is a question reference for none of these.
	NoneOfTheseQuestionRef QuestionRef = "none_of_these"
)

// Lens is a model for a lens.
type Lens struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Version     int       `json:"version" gorm:"uniqueIndex:idx_lens_name_version_team"`
	Name        string    `json:"name" gorm:"uniqueIndex:idx_lens_name_version_team"`
	Description string    `json:"description"`
	Pillars     []Pillar  `json:"pillars"`
	IsDraft     bool      `json:"is_draft"`

	Tags []*Tag `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:lens;"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// GetPillars ...
func (l *Lens) GetPillars() []*Pillar {
	pillars := make([]*Pillar, len(l.Pillars))
	for i, pillar := range l.Pillars {
		pillar := pillar
		pillars[i] = &pillar
	}

	return pillars

}

// Pillar is a model for a pillar.
type Pillar struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key"`
	// Name is the lens pillar name.
	Name string `json:"name"`
	// Description is the lens pillar description.
	Description string `json:"description"`
	// Questions are the lens pillar questions.
	Questions []Question `json:"questions"`
	// Resources are the lens pillar resources.
	Resources []Resource `json:"resources" gorm:"many2many:pillar_resources;"`
	// LensID is the lens ID.
	LensID uuid.UUID `json:"lens_id"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// GetQuestions ...
func (p *Pillar) GetQuestions() []*Question {
	questions := make([]*Question, len(p.Questions))
	for i, question := range p.Questions {
		question := question
		questions[i] = &question
	}

	return questions
}

// Question is a model for a question.
type Question struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key" params:"question" query:"question"`
	// Ref is the question reference.
	Ref string `json:"ref"`
	// Title is the question title.
	Title string `json:"title"`
	// Description is the question description.
	Description string `json:"description"`
	// Resources are the question resources.
	Resources []Resource `json:"resources" gorm:"many2many:question_resources;"`
	Choices   []Choice   `json:"choices"`
	Risks     []Risk     `json:"risks"`
	// PillarID is the pillar ID.
	PillarID int `json:"pillar_id"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// Resource is a model for a resource.
type Resource struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key"`
	// URL is the URL of the resource.
	URL string `json:"url"`
	// Description is the description of the resource.
	Description string `json:"description"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// Choice is a model for a choice.
type Choice struct {
	ID          int         `json:"id" gorm:"primary_key"`
	Ref         QuestionRef `json:"ref"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	QuestionID  int         `json:"question_id"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// Risk is a model for a risk.
type Risk struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Ref        string `json:"ref"`
	Risk       string `json:"risk"`
	Condition  string `json:"condition"`
	QuestionID int    `json:"question_id"`

	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
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

	return nil
}
