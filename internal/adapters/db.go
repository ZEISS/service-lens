package adapters

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
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

var _ ports.Repository = (*DB)(nil)

// NewDB ...
func NewDB(conn *gorm.DB) *DB {
	return &DB{conn}
}

// NewProfile ...
func (d *DB) NewProfile(ctx context.Context, profile *models.Profile) error {
	return d.conn.WithContext(ctx).Create(profile).Error
}

// FetchProfile ...
func (d *DB) FetchProfile(ctx context.Context, id uuid.UUID) (*models.Profile, error) {
	profile := &models.Profile{}
	err := d.conn.WithContext(ctx).Where("id = ?", id).First(profile).Error
	return profile, err
}

// AddLens ...
func (d *DB) AddLens(ctx context.Context) error {
	return nil
}
