package models

import (
	"time"

	"github.com/zeiss/pkg/conv"
	"gorm.io/gorm"
)

// SettingType ...
type SettingType string

// Scan implements the sql.Scanner interface
func (st *SettingType) Scan(value interface{}) error {
	*st = SettingType(value.([]byte))
	return nil
}

// Value implements the driver.Valuer interface
func (st SettingType) Value() (interface{}, error) {
	return conv.String(st), nil
}

// SettingType constants
const (
	SettingTypeString SettingType = "STRING"
	SettingTypeInt    SettingType = "INT"
	SettingTypeFloat  SettingType = "FLOAT"
	SettingTypeBool   SettingType = "BOOL"
)

// Setting ...
type Setting struct {
	// ID is the primary key
	ID string `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// Name is the name of the setting
	Name string `json:"name" gorm:"type:varchar(255);unique_index" validate:"required,min=3"`
	// Value is the value of the setting
	Value string `json:"value" gorm:"type:text" validate:"required,min=3"`
	// NotNull is the not null constraint
	NotNull bool `json:"not_null" gorm:"type:boolean;default:false"`
	// SettingType is the type of the setting
	SettingType SettingType `json:"setting_type" sql:"type:ENUM('STRING', 'INT', 'FLOAT', 'BOOL');default:'STRING'"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at" gore:"index"`
}
