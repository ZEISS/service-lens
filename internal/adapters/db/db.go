package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		&adapters.GothTeam{},
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

var (
	_ ports.ReadTx      = (*datastoreTx)(nil)
	_ ports.ReadWriteTx = (*datastoreTx)(nil)
)

type datastoreTx struct {
	tx *gorm.DB
}

// GetUser is a method that returns the profile of the current user
func (t *datastoreTx) GetUser(ctx context.Context, user *adapters.GothUser) error {
	return t.tx.First(user).Error
}

// ListProfiles is a method that returns a list of profiles
func (t *datastoreTx) ListProfiles(ctx context.Context, team uuid.UUID, pagination *tables.Results[models.Profile]) error {
	return t.tx.
		Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).
		Where("team_id = ?", team).
		Find(&pagination.Rows).Error
}

// GetProfile is a method that returns a profile by ID
func (t *datastoreTx) GetProfile(ctx context.Context, profile *models.Profile) error {
	return t.tx.
		Preload("Answers").
		Where(profile).
		First(profile).Error
}

// CreateProfile is a method that creates a profile
func (t *datastoreTx) CreateProfile(ctx context.Context, profile *models.Profile) error {
	return t.tx.Create(profile).Error
}

// UpdateProfile is a method that updates a profile
func (t *datastoreTx) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(profile).Error
}

// DeleteProfile is a method that deletes a profile
func (t *datastoreTx) DeleteProfile(ctx context.Context, profile *models.Profile) error {
	return t.tx.Delete(profile).Error
}

// ListProfileQuestions is a method that returns a list of profile questions
func (t *datastoreTx) ListProfileQuestions(ctx context.Context, pagination *tables.Results[models.ProfileQuestion]) error {
	return t.tx.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).Preload("Choices").Find(&pagination.Rows).Error
}

// ListEnvironments is a method that returns a list of environments
func (t *datastoreTx) ListEnvironments(ctx context.Context, pagination *tables.Results[models.Environment]) error {
	return t.tx.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).Find(&pagination.Rows).Error
}

// GetEnvironment is a method that returns an environment by ID
func (t *datastoreTx) GetEnvironment(ctx context.Context, environment *models.Environment) error {
	return t.tx.First(environment).Error
}

// CreateEnvironment is a method that creates an environment
func (t *datastoreTx) CreateEnvironment(ctx context.Context, environment *models.Environment) error {
	return t.tx.Create(environment).Error
}

// UpdateEnvironment is a method that updates an environment
func (t *datastoreTx) UpdateEnvironment(ctx context.Context, environment *models.Environment) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(environment).Error
}

// DeleteEnvironment is a method that deletes an environment
func (t *datastoreTx) DeleteEnvironment(ctx context.Context, environment *models.Environment) error {
	return t.tx.Delete(environment).Error
}

// ListLenses is a method that returns a list of lenses
func (t *datastoreTx) ListLenses(ctx context.Context, pagination *tables.Results[models.Lens]) error {
	return t.tx.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).Find(&pagination.Rows).Error
}

// GetLens is a method that returns a lens by ID
func (t *datastoreTx) GetLens(ctx context.Context, lens *models.Lens) error {
	return t.tx.
		Preload("Pillars").
		Preload("Pillars.Questions").
		First(lens).Error
}

// GetLensQuestion is a method that returns a lens question by ID
func (t *datastoreTx) GetLensQuestion(ctx context.Context, question *models.Question) error {
	return t.tx.Preload("Choices").First(question).Error
}

// CreateLens is a method that creates a lens
func (t *datastoreTx) CreateLens(ctx context.Context, lens *models.Lens) error {
	return t.tx.Create(lens).Error
}

// UpdateLens is a method that updates a lens
func (t *datastoreTx) UpdateLens(ctx context.Context, lens *models.Lens) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(lens).Error
}

// DeleteLens is a method that deletes a lens
func (t *datastoreTx) DeleteLens(ctx context.Context, lens *models.Lens) error {
	return t.tx.Delete(lens).Error
}

// ListWorkloads is a method that returns a list of workloads
func (t *datastoreTx) ListWorkloads(ctx context.Context, pagination *tables.Results[models.Workload]) error {
	return t.tx.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).
		Preload("Lenses").
		Preload("Profile").
		Preload("Environment").
		Find(&pagination.Rows).Error
}

// GetWorkload is a method that returns a workload by ID
func (t *datastoreTx) GetWorkload(ctx context.Context, workload *models.Workload) error {
	return t.tx.Preload(clause.Associations).First(workload).Error
}

// CreateWorkload is a method that creates a workload
func (t *datastoreTx) CreateWorkload(ctx context.Context, workload *models.Workload) error {
	return t.tx.Create(workload).Error
}

// UpdateWorkload is a method that updates a workload
func (t *datastoreTx) UpdateWorkload(ctx context.Context, workload *models.Workload) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(workload).Error
}

// DeleteWorkload is a method that deletes a workload
func (t *datastoreTx) DeleteWorkload(ctx context.Context, workload *models.Workload) error {
	return t.tx.Delete(workload).Error
}

// UpdateWorkloadAnswer is a method that updates a workload answer
func (t *datastoreTx) UpdateWorkloadAnswer(ctx context.Context, answer *models.WorkloadLensQuestionAnswer) error {
	err := t.tx.
		Debug().
		Session(&gorm.Session{FullSaveAssociations: true}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "workload_id"}, {Name: "lens_id"}, {Name: "question_id"}},
			UpdateAll: true,
		}).
		Where(&models.WorkloadLensQuestionAnswer{WorkloadID: answer.WorkloadID, LensID: answer.LensID, QuestionID: answer.QuestionID}).
		Omit("Choices.*").
		Save(answer).Error
	if err != nil {
		return err
	}

	return t.tx.Model(answer).Association("Choices").Replace(answer.Choices)
}

// GetWorkloadAnswer is a method that returns a workload answer by ID
func (t *datastoreTx) GetWorkloadAnswer(ctx context.Context, answer *models.WorkloadLensQuestionAnswer) error {
	return t.tx.
		Where(&models.WorkloadLensQuestionAnswer{WorkloadID: answer.WorkloadID, LensID: answer.LensID, QuestionID: answer.QuestionID}).
		Preload(clause.Associations).
		First(answer).Error
}

// GetTeam is a method that returns a team by ID
func (t *datastoreTx) GetTeam(ctx context.Context, team *adapters.GothTeam) error {
	return t.tx.First(team).Error
}

// ListTeams is a method that returns a list of teams
func (t *datastoreTx) ListTeams(ctx context.Context, pagination *tables.Results[adapters.GothTeam]) error {
	return t.tx.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).Find(&pagination.Rows).Error
}

// CreateTeam is a method that creates a team
func (t *datastoreTx) CreateTeam(ctx context.Context, team *adapters.GothTeam) error {
	return t.tx.Create(team).Error
}

// UpdateTeam is a method that updates a team
func (t *datastoreTx) UpdateTeam(ctx context.Context, team *adapters.GothTeam) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(team).Error
}

// DeleteTeam is a method that deletes a team
func (t *datastoreTx) DeleteTeam(ctx context.Context, team *adapters.GothTeam) error {
	return t.tx.Delete(team).Error
}
