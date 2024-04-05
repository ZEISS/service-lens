package adapters

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
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
		&models.WorkloadLensQuestionAnswer{},
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

// DestroyLens ...
func (d *DB) DestroyLens(ctx context.Context, id uuid.UUID) error {
	return d.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&models.Lens{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// GetLensByID ...
func (d *DB) GetLensByID(ctx context.Context, id uuid.UUID) (*models.Lens, error) {
	lens := &models.Lens{ID: id}
	err := d.conn.WithContext(ctx).Preload("Tags").Preload("Pillars").Preload("Pillars.Questions").Preload("Pillars.Questions.Choices").Find(lens).Error
	if err != nil {
		return nil, err
	}

	return lens, err
}

// UpdateProfile ...
func (d *DB) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	return d.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Save(profile).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// DestroyProfile ...
func (d *DB) DestroyProfile(ctx context.Context, id uuid.UUID) error {
	return d.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&models.Profile{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// GetProfileByID ...
func (d *DB) GetProfileByID(ctx context.Context, id uuid.UUID) (*models.Profile, error) {
	profile := &models.Profile{ID: id}
	err := d.conn.WithContext(ctx).Preload("Questions").Find(profile).Error
	if err != nil {
		return nil, err
	}

	return profile, err
}

// GetPillarById ...
func (d *DB) GetPillarById(ctx context.Context, teamSlug string, lensId uuid.UUID, id int) (*models.Pillar, error) {
	pillar := &models.Pillar{
		ID:     id,
		LensID: lensId,
	}

	err := d.conn.WithContext(ctx).Preload("Questions").Find(pillar).Error
	if err != nil {
		return nil, err
	}

	return pillar, err
}

// ListLenses ...
func (d *DB) ListLenses(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Lens, error) {
	lenses := []*models.Lens{}
	err := d.conn.WithContext(ctx).Where("team_id = (?)", d.conn.WithContext(ctx).Select("id").Where("slug = ?", teamSlug).Table("teams")).Limit(pagination.Limit).Offset(pagination.Offset).Find(&lenses).Error
	if err != nil {
		return nil, err
	}

	return lenses, nil
}

// ListAnswers ...
func (d *DB) ListAnswers(ctx context.Context, workloadID uuid.UUID, lensID uuid.UUID, questionID int) (*models.WorkloadLensQuestionAnswer, error) {
	answers := &models.WorkloadLensQuestionAnswer{}
	err := d.conn.WithContext(ctx).Where("workload_id = ? AND lens_id = ? AND question_id = ?", workloadID, lensID, questionID).Preload("Choices").Find(&answers).Error
	if err != nil {
		return nil, err
	}

	return answers, nil
}

// UpdateAnswers ...
func (d *DB) UpdateAnswers(ctx context.Context, workloadID uuid.UUID, lensID uuid.UUID, questionID int, choices []int, doesNotApply bool, notes string) error {
	return d.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		answer := &models.WorkloadLensQuestionAnswer{
			WorkloadID: workloadID,
			LensID:     lensID,
			QuestionID: questionID,
		}

		err := tx.Where("workload_id = ? AND lens_id = ? AND question_id = ?", workloadID, lensID, questionID).FirstOrCreate(&answer).Error
		if err != nil {
			return err
		}

		answer.DoesNotApply = doesNotApply
		answer.Notes = notes

		err = tx.Save(&answer).Error
		if err != nil {
			return err
		}

		c := []*models.Choice{}
		for _, choice := range choices {
			c = append(c, &models.Choice{
				ID:         choice,
				QuestionID: questionID,
			})
		}

		err = tx.Model(&answer).Association("Choices").Replace(c)
		if err != nil {
			return err
		}

		return nil
	})
}

// ListProfiles ...
func (d *DB) ListProfiles(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Profile, error) {
	profiles := []*models.Profile{}
	err := d.conn.WithContext(ctx).Where("team_id = (?)", d.conn.WithContext(ctx).Select("id").Where("slug = ?", teamSlug).Table("teams")).Limit(pagination.Limit).Offset(pagination.Offset).Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

// ListWorkloads ...
func (d *DB) ListWorkloads(ctx context.Context, teamSlug string, pagination *models.Pagination) ([]*models.Workload, error) {
	workloads := []*models.Workload{}

	err := d.conn.WithContext(ctx).Where("team_id = (?)", d.conn.WithContext(ctx).Select("id").Where("slug = ?", teamSlug).Table("teams")).Limit(pagination.Limit).Offset(pagination.Offset).Find(&workloads).Error
	if err != nil {
		return nil, err
	}

	return workloads, nil
}

// ShowWorkload ...
func (d *DB) IndexWorkload(ctx context.Context, id uuid.UUID) (*models.Workload, error) {
	workload := &models.Workload{
		ID: id,
	}

	err := d.conn.WithContext(ctx).Preload("Lenses").Preload("Tags").Preload("Profile").Preload("Answers").Preload("Answers.Choices").Find(workload).Error
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
	return d.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&models.Workload{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// AddTeam ...
func (d *DB) AddTeam(ctx context.Context, team *authz.Team) (*authz.Team, error) {
	err := d.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Save(team).Error
		if err != nil {
			return err
		}

		return nil
	})
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

// GetTeamByID ...
func (d *DB) GetTeamByID(ctx context.Context, id uuid.UUID) (*authz.Team, error) {
	team := &authz.Team{
		ID: id,
	}

	err := d.conn.WithContext(ctx).Find(&team).Error
	if err != nil {
		return nil, err
	}

	return team, err
}

// GetUserByID ...
func (d *DB) GetUserByID(ctx context.Context, id uuid.UUID) (*authz.User, error) {
	user := &authz.User{
		User: &adapters.User{
			ID: id,
		},
	}

	err := d.conn.WithContext(ctx).Preload("Teams").Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

// ListUsers ...
func (d *DB) ListUsers(ctx context.Context, pagination *models.Pagination) ([]*authz.User, error) {
	users := []*authz.User{}
	err := d.conn.WithContext(ctx).Preload("Teams").Limit(pagination.Limit).Offset(pagination.Offset).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// AddUser ...
func (d *DB) AddUser(ctx context.Context, user *authz.User) (*authz.User, error) {
	err := d.conn.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser ...
func (d *DB) UpdateUser(ctx context.Context, user *authz.User) (*authz.User, error) {
	err := d.conn.WithContext(ctx).Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser ...
func (d *DB) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return d.conn.WithContext(ctx).Delete(&authz.User{}, id).Error
}

// DeleteTeam ...
func (d *DB) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	return d.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&authz.Team{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// TotalCountWorkloads ...
func (d *DB) TotalCountWorkloads(ctx context.Context, teamSlug string) (int, error) {
	var count int64
	err := d.conn.WithContext(ctx).Model(&models.Workload{}).Where("team_id = (?)", d.conn.WithContext(ctx).Select("id").Where("slug = ?", teamSlug).Table("teams")).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// TotalCountLenses ...
func (d *DB) TotalCountLenses(ctx context.Context, teamSlug string) (int, error) {
	var count int64
	err := d.conn.WithContext(ctx).Model(&models.Lens{}).Where("team_id = (?)", d.conn.WithContext(ctx).Select("id").Where("slug = ?", teamSlug).Table("teams")).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// TotalCountProfiles ...
func (d *DB) TotalCountProfiles(ctx context.Context, teamSlug string) (int, error) {
	var count int64
	err := d.conn.WithContext(ctx).Model(&models.Profile{}).Where("team_id = (?)", d.conn.WithContext(ctx).Select("id").Where("slug = ?", teamSlug).Table("teams")).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
