package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

type database struct {
	conn *gorm.DB
}

// NewDatastore returns a new instance of db.
func NewDB(conn *gorm.DB) (ports.Datastore, error) {
	return &database{
		conn: conn,
	}, nil
}

// Close closes the database connection.
func (d *database) Close() error {
	sqlDB, err := d.conn.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// RunMigrations runs the database migrations.
func (d *database) Migrate(ctx context.Context) error {
	return d.conn.WithContext(ctx).AutoMigrate(
		&adapters.GothUser{},
		&adapters.GothAccount{},
		&adapters.GothSession{},
		&adapters.GothVerificationToken{},
		&models.ProfileQuestion{},
		&models.ProfileQuestionChoice{},
		&models.ProfileQuestionAnswer{},
		&models.Environment{},
		&models.Profile{},
		&models.Lens{},
		&models.Pillar{},
		&models.Question{},
		&models.Resource{},
		&models.Choice{},
		&models.Risk{},
		&models.Workload{},
		&models.Tag{},
		&models.WorkloadLensQuestionAnswer{},
	)
}

// ReadWriteTx starts a read only transaction.
func (d *database) ReadWriteTx(ctx context.Context, fn func(context.Context, ports.ReadWriteTx) error) error {
	tx := d.conn.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(ctx, &datastoreTx{tx}); err != nil {
		tx.Rollback()
	}

	if err := tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// ReadTx starts a read only transaction.
func (d *database) ReadTx(ctx context.Context, fn func(context.Context, ports.ReadTx) error) error {
	tx := d.conn.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(ctx, &datastoreTx{tx}); err != nil {
		tx.Rollback()
	}

	if err := tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

var _ ports.ReadTx = (*datastoreTx)(nil)
var _ ports.ReadWriteTx = (*datastoreTx)(nil)

type datastoreTx struct {
	tx *gorm.DB
}

// GetUser is a method that returns the profile of the current user
func (t *datastoreTx) GetUser(ctx context.Context, user *adapters.GothUser) error {
	return t.tx.First(user).Error
}

// ListProfiles is a method that returns a list of profiles
func (t *datastoreTx) ListProfiles(ctx context.Context, pagination *tables.Results[models.Profile]) error {
	return t.tx.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).Find(&pagination.Rows).Error
}

// GetProfile is a method that returns a profile by ID
func (t *datastoreTx) GetProfile(ctx context.Context, profile *models.Profile) error {
	return t.tx.First(profile).Error
}

// CreateProfile is a method that creates a profile
func (t *datastoreTx) CreateProfile(ctx context.Context, profile *models.Profile) error {
	return t.tx.Create(profile).Error
}

// UpdateProfile is a method that updates a profile
func (t *datastoreTx) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(profile).Error
}
