package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Profile represents a business profile.
type Profile struct {
	// ID ...
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" query:"id" params:"id"`
	// Name of the profile.
	Name string `json:"name" form:"name" validate:"required,min=3,max=100"`
	// Description of the profile.
	Description string `json:"description" form:"description" validate:"required,min=3,max=1024"`
	// Answers are the possible answers for the question.
	Answers []ProfileQuestionAnswer `json:"answers" form:"answers" gorm:"many2many:profiles_answers;"`
	// Tags ...
	Tags []Tag `json:"tags" gorm:"many2many:profile_tags;"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// GetAnswers returns the answers of the profile.
func (p *Profile) GetAnswers() []*ProfileQuestionAnswer {
	answers := make([]*ProfileQuestionAnswer, len(p.Answers))
	for i, answer := range p.Answers {
		answer := answer
		answers[i] = &answer
	}

	return answers
}

// IsChoosen returns true if the question is answered.
func (p *Profile) IsChoosen(id int) bool {
	for _, answer := range p.Answers {
		if answer.ChoiceID == id {
			return true
		}
	}

	return false
}

// Question represents a business profile question.
type ProfileQuestion struct {
	// ID is the identifier of the question.
	ID int `json:"id" gorm:"primary_key"`
	// Title is the title of the question.
	Title string `json:"title" form:"title" validate:"required,min=3,max=100"`
	// Description is the description of the question.
	Description string `json:"description" form:"description" validate:"required,min=3,max=1024"`
	// MultipleChoice is a flag to indicate if the question is multiple choice.
	MulipleChoice bool `json:"multiple_choice" form:"multiple_choice" validate:"required"`
	// Choices are the possible choices for the question.
	Choices []ProfileQuestionChoice `json:"choices" form:"choices"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// GetChoices returns the choices of the question.
func (q *ProfileQuestion) GetChoices() []*ProfileQuestionChoice {
	choices := make([]*ProfileQuestionChoice, len(q.Choices))
	for i, choice := range q.Choices {
		choice := choice
		choices[i] = &choice
	}
	return choices
}

// ProfileQuestionChoice is a model for a choice.
type ProfileQuestionChoice struct {
	// ID is the identifier of the choice.
	ID int `json:"id" gorm:"primary_key"`
	// Ref is the reference of the choice.
	Ref QuestionRef `json:"ref"`
	// Title is the title of the choice.
	Title string `json:"title"`
	// Description is the description of the choice.
	Description string `json:"description"`
	// ProfileQuestion is the question of the choice.
	ProfileQuestion ProfileQuestion `json:"profile_question" gorm:"foreignkey:ProfileQuestionID;"`
	// ProfileQuestionID is the identifier of the question.
	ProfileQuestionID int `json:"profile_question_id"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// ProfileQuestionAnswer represents a business profile question answer.
type ProfileQuestionAnswer struct {
	// ID is the identifier of the answer.
	ID int `json:"id" gorm:"primary_key" form:"id"`
	// Choice is the choice of the answer.
	Choice ProfileQuestionChoice `json:"choice" gorm:"foreignkey:ChoiceID;"`
	// ChoiceID is the identifier of the choice.
	ChoiceID int `json:"choice_id"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
