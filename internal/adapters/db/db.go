package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/tables"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ ports.ReadTx = (*readTxImpl)(nil)

type readTxImpl struct {
	conn *gorm.DB
}

// NewReadTx ...
func NewReadTx() seed.ReadTxFactory[ports.ReadTx] {
	return func(db *gorm.DB) (ports.ReadTx, error) {
		return &readTxImpl{conn: db}, nil
	}
}

// GetUser is a method that returns the profile of the current user
func (r *readTxImpl) GetUser(ctx context.Context, user *adapters.GothUser) error {
	return r.conn.Where(user).First(user).Error
}

// ListProfiles is a method that returns a list of profiles
func (r *readTxImpl) ListProfiles(ctx context.Context, team uuid.UUID, pagination *tables.Results[models.Profile]) error {
	return r.conn.
		Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).
		Where("team_id = ?", team).
		Find(&pagination.Rows).Error
}

// ListProfileQuestions is a method that returns a list of profile questions
func (r *readTxImpl) ListProfileQuestions(ctx context.Context, pagination *tables.Results[models.ProfileQuestion]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Preload("Choices").Find(&pagination.Rows).Error
}

// ListEnvironments is a method that returns a list of environments
func (r *readTxImpl) ListEnvironments(ctx context.Context, pagination *tables.Results[models.Environment]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// GetEnvironment is a method that returns an environment by ID
func (r *readTxImpl) GetEnvironment(ctx context.Context, environment *models.Environment) error {
	return r.conn.Where(environment).First(environment).Error
}

// GetProfile is a method that returns a profile by ID
func (r *readTxImpl) GetProfile(ctx context.Context, profile *models.Profile) error {
	return r.conn.
		Preload("Answers").
		Where(profile).
		First(profile).Error
}

// ListLenses is a method that returns a list of lenses
func (r *readTxImpl) ListLenses(ctx context.Context, pagination *tables.Results[models.Lens]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// GetLens is a method that returns a lens by ID
func (r *readTxImpl) GetLens(ctx context.Context, lens *models.Lens) error {
	return r.conn.
		Preload("Pillars").
		Preload("Pillars.Questions").
		First(lens).Error
}

// GetLensQuestion is a method that returns a lens question by ID
func (r *readTxImpl) GetLensQuestion(ctx context.Context, question *models.Question) error {
	return r.conn.Preload("Choices").First(question).Error
}

// ListWorkloads is a method that returns a list of workloads
func (r *readTxImpl) ListWorkloads(ctx context.Context, pagination *tables.Results[models.Workload]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).
		Preload("Lenses").
		Preload("Profile").
		Preload("Environment").
		Find(&pagination.Rows).Error
}

// GetWorkloadAnswer is a method that returns a workload answer by ID
func (r *readTxImpl) GetWorkloadAnswer(ctx context.Context, answer *models.WorkloadLensQuestionAnswer) error {
	return r.conn.
		Where(&models.WorkloadLensQuestionAnswer{WorkloadID: answer.WorkloadID, LensID: answer.LensID, QuestionID: answer.QuestionID}).
		Preload(clause.Associations).
		First(answer).Error
}

// GetTeam is a method that returns a team by ID
func (r *readTxImpl) GetTeam(ctx context.Context, team *adapters.GothTeam) error {
	return r.conn.Where(team).First(team).Error
}

// ListTeams is a method that returns a list of teams
func (r *readTxImpl) ListTeams(ctx context.Context, pagination *tables.Results[adapters.GothTeam]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// GetWorkload is a method that returns a workload by ID
func (r *readTxImpl) GetWorkload(ctx context.Context, workload *models.Workload) error {
	return r.conn.Preload(clause.Associations).Where(workload).First(workload).Error
}

type writeTxImpl struct {
	conn *gorm.DB
	readTxImpl
}

// NewWriteTx ...
func NewWriteTx() seed.ReadWriteTxFactory[ports.ReadWriteTx] {
	return func(db *gorm.DB) (ports.ReadWriteTx, error) {
		return &writeTxImpl{conn: db}, nil
	}
}

// CreateProfile is a method that creates a profile
func (rw *writeTxImpl) CreateProfile(ctx context.Context, profile *models.Profile) error {
	return rw.conn.Create(profile).Error
}

// UpdateProfile is a method that updates a profile
func (rw *writeTxImpl) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(profile).Error
}

// DeleteProfile is a method that deletes a profile
func (rw *writeTxImpl) DeleteProfile(ctx context.Context, profile *models.Profile) error {
	return rw.conn.Delete(profile).Error
}

// CreateEnvironment is a method that creates an environment
func (rw *writeTxImpl) CreateEnvironment(ctx context.Context, environment *models.Environment) error {
	return rw.conn.Create(environment).Error
}

// UpdateEnvironment is a method that updates an environment
func (rw *writeTxImpl) UpdateEnvironment(ctx context.Context, environment *models.Environment) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(environment).Error
}

// DeleteEnvironment is a method that deletes an environment
func (rw *writeTxImpl) DeleteEnvironment(ctx context.Context, environment *models.Environment) error {
	return rw.conn.Delete(environment).Error
}

// CreateLens is a method that creates a lens
func (rw *writeTxImpl) CreateLens(ctx context.Context, lens *models.Lens) error {
	return rw.conn.Create(lens).Error
}

// UpdateLens is a method that updates a lens
func (rw *writeTxImpl) UpdateLens(ctx context.Context, lens *models.Lens) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(lens).Error
}

// DeleteLens is a method that deletes a lens
func (rw *writeTxImpl) DeleteLens(ctx context.Context, lens *models.Lens) error {
	return rw.conn.Delete(lens).Error
}

// CreateWorkload is a method that creates a workload
func (rw *writeTxImpl) CreateWorkload(ctx context.Context, workload *models.Workload) error {
	return rw.conn.Create(workload).Error
}

// UpdateWorkload is a method that updates a workload
func (rw *writeTxImpl) UpdateWorkload(ctx context.Context, workload *models.Workload) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(workload).Error
}

// DeleteWorkload is a method that deletes a workload
func (rw *writeTxImpl) DeleteWorkload(ctx context.Context, workload *models.Workload) error {
	return rw.conn.Delete(workload).Error
}

// UpdateWorkloadAnswer is a method that updates a workload answer
func (rw *writeTxImpl) UpdateWorkloadAnswer(ctx context.Context, answer *models.WorkloadLensQuestionAnswer) error {
	err := rw.conn.
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

	return rw.conn.Model(answer).Association("Choices").Replace(answer.Choices)
}

// CreateTeam is a method that creates a team
func (rw *writeTxImpl) CreateTeam(ctx context.Context, team *adapters.GothTeam) error {
	return rw.conn.Create(team).Error
}

// UpdateTeam is a method that updates a team
func (rw *writeTxImpl) UpdateTeam(ctx context.Context, team *adapters.GothTeam) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(team).Error
}

// DeleteTeam is a method that deletes a team
func (rw *writeTxImpl) DeleteTeam(ctx context.Context, team *adapters.GothTeam) error {
	return rw.conn.Delete(team).Error
}
