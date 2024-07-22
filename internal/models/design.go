package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

// Design ...
type Design struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// Title of the design
	Title string `json:"title" form:"title" validate:"required,min=3,max=255"`
	// Body of the design in markdown, HTML, or plain text
	Body string `form:"body" gorm:"type:text"`
	// Tags are the tags associated with the environment
	Tags []Tag `json:"tags" gorm:"polymorphic:Taggable;"`
	// AuthorID is the foreign key to the author
	AuthorID uuid.UUID `json:"author_id"`
	// Author is the author
	Author adapters.GothUser `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	// Comments are the comments associated with the design
	Comments []DesignComment `json:"comments"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// DesignRevision ...
type DesignRevision struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// DesignID is the foreign key to the design
	DesignID uuid.UUID `json:"design_id" gorm:"type:uuid;index" params:"design_id"`
	// Design is the design
	Design Design `json:"design" gorm:"foreignKey:DesignID;references:ID"`
	// Title of the design revision
	Title string `json:"title" form:"title" validate:"required,min=3,max=1024"`
	// Body of the design in markdown, HTML, or plain text
	Body string `gorm:"type:text"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
}

// DesignComment ...
type DesignComment struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// DesignID is the foreign key to the design
	DesignID uuid.UUID `json:"design_id" gorm:"type:uuid;index" params:"id"`
	// Design is the design
	Design Design `json:"design" gorm:"foreignKey:DesignID;references:ID" validate:"-"`
	// Comment is the comment
	Comment string `json:"comment" form:"comment" gorm:"type:text" validate:"required,min=3"`
	// AuthorID is the foreign key to the author
	AuthorID uuid.UUID `json:"author_id"`
	// Author is the author
	Author adapters.GothUser `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	// ParentID is the foreign key to the parent comment
	ParentID *uuid.UUID `json:"parent_id" gorm:"type:uuid;index" params:"parent_id"`
	// Parent is the parent comment
	Parent *DesignComment `json:"parent" gorm:"foreignKey:ParentID;references:ID"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}

// DesignCommentRevision ...
type DesignCommentRevision struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// DesignCommentID is the foreign key to the design comment
	DesignCommentID uuid.UUID `json:"design_comment_id" gorm:"type:uuid;index" params:"design_comment_id"`
	// DesignComment is the design comment
	DesignComment DesignComment `json:"design_comment" gorm:"foreignKey:DesignCommentID;references:ID"`
	// Comment is the comment
	Comment string `json:"comment" form:"comment" gorm:"type:text" validate:"required,min=3"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
}
