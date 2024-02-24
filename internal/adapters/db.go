package adapters

import (
	"context"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
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
		&models.Lens{},
		&models.Pillar{},
		&models.Question{},
		&models.Resource{},
		&models.Choice{},
		&models.Risk{},
		&models.Workload{},
		&models.Tag{},
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
	if err != nil {
		return nil, err
	}

	return profile, err
}

// AddLens ...
func (d *DB) AddLens(ctx context.Context, lens *models.Lens) (*models.Lens, error) {
	err := d.conn.WithContext(ctx).Create(lens).Error
	if err != nil {
		return nil, err
	}

	return lens, nil
}

// GetLensByID ...
func (d *DB) GetLensByID(ctx context.Context, id uuid.UUID) (*models.Lens, error) {
	lens := &models.Lens{
		ID: id,
	}
	err := d.conn.WithContext(ctx).Preload("Tags").Preload("Pillars").Find(lens).Error
	if err != nil {
		return nil, err
	}

	return lens, err
}

// ListLenses ...
func (d *DB) ListLenses(ctx context.Context, pagination *models.Pagination) ([]*models.Lens, error) {
	lenses := []*models.Lens{}
	err := d.conn.WithContext(ctx).Preload("Tags").Limit(pagination.Limit).Offset(pagination.Offset).Find(&lenses).Error
	if err != nil {
		return nil, err
	}

	return lenses, nil
}

// ListProfiles ...
func (d *DB) ListProfiles(ctx context.Context, pagination *models.Pagination) ([]*models.Profile, error) {
	profiles := []*models.Profile{}
	err := d.conn.WithContext(ctx).Where("team_id = ?", "244b0756-4ace-44f2-b5cd-cd6fa8918e5f").Limit(pagination.Limit).Offset(pagination.Offset).Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

// ListWorkloads ...
func (d *DB) ListWorkloads(ctx context.Context, pagination *models.Pagination) ([]*models.Workload, error) {
	workloads := []*models.Workload{}

	err := d.conn.WithContext(ctx).Where("name LIKE ?", "%"+pagination.Search+"%").Limit(pagination.Limit).Offset(pagination.Offset).Find(&workloads).Error
	if err != nil {
		return nil, err
	}

	return workloads, nil
}

// ShowWorkload ...
func (d *DB) ShowWorkload(ctx context.Context, id uuid.UUID) (*models.Workload, error) {
	workload := &models.Workload{
		ID: id,
	}

	err := d.conn.WithContext(ctx).Preload("Lenses").Preload("Tags").Find(workload).Error
	if err != nil {
		return nil, err
	}

	return workload, nil
}

// StoreWorkload ...
func (d *DB) StoreWorkload(ctx context.Context, workload *models.Workload) error {
	return d.conn.WithContext(ctx).Create(workload).Error
}

// DestroyWorkload ...
func (d *DB) DestroyWorkload(ctx context.Context, id uuid.UUID) error {
	return d.conn.WithContext(ctx).Delete(&models.Workload{}, id).Error
}

// AddTeam ...
func (d *DB) AddTeam(ctx context.Context, team *authz.Team) (*authz.Team, error) {
	err := d.conn.WithContext(ctx).Create(team).Error
	if err != nil {
		return nil, err
	}

	return team, nil
}

// GetTeamBySlug ...
func (d *DB) GetTeamBySlug(ctx context.Context, slug string) (*authz.Team, error) {
	team := &authz.Team{}

	err := d.conn.WithContext(ctx).Where("slug = ?", slug).First(team).Error
	if err != nil {
		return nil, err
	}

	return team, err
}

// ListTeams ...
func (d *DB) ListTeams(ctx context.Context, pagination *models.Pagination) ([]*authz.Team, error) {
	teams := []*authz.Team{}
	err := d.conn.WithContext(ctx).Limit(pagination.Limit).Offset(pagination.Offset).Find(&teams).Error
	if err != nil {
		return nil, err
	}

	return teams, nil
}
