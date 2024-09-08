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
	// ID is the primary key.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Version is the lens version.
	Version int `json:"version" gorm:"uniqueIndex:idx_lens_name_version_team"`
	// Name is the lens name.
	Name string `json:"name" gorm:"many2many:taggables;uniqueIndex:idx_lens_name_version_team"`
	// Description is the lens description.
	Description string `json:"description"`
	// Pillars are the lens pillars.
	Pillars []Pillar `json:"pillars"`
	// IsDraft is the lens draft status.
	IsDraft bool `json:"is_draft"`
	// Tags are the tags associated with the lens.
	Tags []Tag `json:"tags" gorm:"many2many:lens_tags;"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// GetQuestions ...
func (l *Lens) GetQuestions() []Question {
	questions := make([]Question, 0)
	for _, pillar := range l.Pillars {
		for _, question := range pillar.Questions {
			question := question
			questions = append(questions, question)
		}
	}

	return questions
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

// TotalQuestions ...
func (l *Lens) TotalQuestions() int {
	return len(l.GetQuestions())
}

// TotalPillars ...
func (l *Lens) TotalPillars() int {
	return len(l.Pillars)
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
	// Choices are the question choices.
	Choices []Choice `json:"choices"`
	// Risks are the question risks.
	Risks []Risk `json:"risks"`
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
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key" form:"id" query:"id" params:"id" validate:"required"`
	// Ref is the choice reference.
	Ref QuestionRef `json:"ref"`
	// Title is the choice title.
	Title string `json:"title"`
	// Description is the choice description.
	Description string `json:"description"`
	// QuestionID is the question ID.
	QuestionID int `json:"question_id"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// Risk is a model for a risk.
type Risk struct {
	// ID is the primary key.
	ID int `json:"id" gorm:"primary_key"`
	// Ref is the risk reference.
	Ref string `json:"ref"`
	// Title is the risk title.
	Risk string `json:"risk"`
	// Condition is the risk condition.
	Condition string `json:"condition"`
	// QuestionID is the question ID.
	QuestionID int `json:"question_id"`
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
