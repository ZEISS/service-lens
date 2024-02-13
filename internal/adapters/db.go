package adapters

import (
	"github.com/zeiss/service-lens/internal/models"
	"gorm.io/gorm"
)

// DB ...
type DB struct {
	conn *gorm.DB
}

// RunMigration ...
func (d *DB) RunMigration() error {
	return d.conn.AutoMigrate(
		&models.ProfileQuestionAnswer{},
		&models.ProfileQuestion{},
		&models.ProfileQuestions{},
		&models.Profile{},
	)
}

// NewDB ...
func NewDB(conn *gorm.DB) *DB {
	return &DB{conn}
}
