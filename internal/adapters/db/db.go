package db

import (
	"context"
	"slices"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/pkg/dbx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ ports.ReadTx = (*readTxImpl)(nil)

type readTxImpl struct {
	conn *gorm.DB
}

// NewReadTx ...
func NewReadTx() dbx.ReadTxFactory[ports.ReadTx] {
	return func(db *gorm.DB) (ports.ReadTx, error) {
		return &readTxImpl{conn: db}, nil
	}
}

// GetUser is a method that returns the profile of the current user
func (r *readTxImpl) GetUser(ctx context.Context, user *adapters.GothUser) error {
	return r.conn.Where(user).First(user).Error
}

// GetDesign is a method that returns a design by ID
func (r *readTxImpl) GetDesign(ctx context.Context, design *models.Design) error {
	return r.conn.
		Preload(clause.Associations).
		Preload("Comments.Author").
		Preload("Comments.Reactions").
		Preload("Comments.Reactions.Reactor").
		Where(design).
		First(design, design.ID).Error
}

// ListDesigns is a method that returns a list of designs
func (r *readTxImpl) ListDesigns(ctx context.Context, pagination *tables.Results[models.Design]) error {
	return r.conn.Debug().Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// ListProfiles is a method that returns a list of profiles
func (r *readTxImpl) ListProfiles(ctx context.Context, pagination *tables.Results[models.Profile]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
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
	return r.conn.Debug().Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// GetLens is a method that returns a lens by ID
func (r *readTxImpl) GetLens(ctx context.Context, lens *models.Lens) error {
	return r.conn.
		Preload(clause.Associations).
		Preload("Pillars").
		Preload("Pillars.Questions").
		Preload("Pillars.Questions.Choices").
		Preload("Pillars.Questions.Risks").
		First(lens).Error
}

// GetLensQuestion is a method that returns a lens question by ID
func (r *readTxImpl) GetLensQuestion(ctx context.Context, question *models.Question) error {
	return r.conn.
		Preload(clause.Associations).
		First(question, question.ID).
		Error
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
		FirstOrCreate(answer).Error
}

// GetWorkload is a method that returns a workload by ID
func (r *readTxImpl) GetWorkload(ctx context.Context, workload *models.Workload) error {
	return r.conn.
		Preload(clause.Associations).
		Preload("Lenses.Pillars").
		Preload("Lenses.Pillars.Questions").
		Preload("Lenses.Pillars.Questions.Choices").
		Preload("Lenses.Pillars.Questions.Risks").
		Preload("Answers.Choices").
		Preload("Answers.Risk").
		Where(workload).
		First(workload, workload.ID).
		Error
}

// ListTags is a method that returns a list of tags
func (r *readTxImpl) ListTags(ctx context.Context, pagination *tables.Results[models.Tag]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// ListWorkflows is a method that returns a list of workflows
func (r *readTxImpl) ListWorkflows(ctx context.Context, pagination *tables.Results[models.Workflow]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// ListTemplates is a method that returns a list of templates
func (r *readTxImpl) ListTemplates(ctx context.Context, pagination *tables.Results[models.Template]) error {
	return r.conn.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, r.conn)).Find(&pagination.Rows).Error
}

// GetTemplate is a method that returns a template by ID
func (r *readTxImpl) GetTemplate(ctx context.Context, template *models.Template) error {
	return r.conn.First(template, template.ID).Error
}

// GetTotalNumberOfProfiles is a method that returns the total number of profiles
func (r *readTxImpl) GetTotalNumberOfProfiles(ctx context.Context, total *int64) error {
	return r.conn.Model(&models.Profile{}).Count(total).Error
}

// GetTotalNumberOfDesigns is a method that returns the total number of designs
func (r *readTxImpl) GetTotalNumberOfDesigns(ctx context.Context, total *int64) error {
	return r.conn.Model(&models.Design{}).Count(total).Error
}

// GetTotalNumberOfWorkloads is a method that returns the total number of workloads
func (r *readTxImpl) GetTotalNumberOfWorkloads(ctx context.Context, total *int64) error {
	return r.conn.Model(&models.Workload{}).Count(total).Error
}

// GetWorkflow is a method that returns a workflow by ID
func (r *readTxImpl) GetWorkflow(ctx context.Context, workflow *models.Workflow) error {
	return r.conn.
		Preload(clause.Associations).
		First(workflow, workflow.ID).
		Error
}

// GetDesignComment is a method that returns a design comment by ID
func (r *readTxImpl) GetDesignComment(ctx context.Context, comment *models.DesignComment) error {
	return r.conn.
		Preload(clause.Associations).
		Preload("Reactions").
		Preload("Reactions.Reactor").
		First(comment, comment.ID).Error
}

// ListDesignCommentReactions is a method that returns a list of design comment reactions
func (r *readTxImpl) ListDesignCommentReactions(ctx context.Context, comment *models.DesignComment) error {
	return r.conn.Preload("Reactions").Preload("Reactions.Reactor").First(comment, comment.ID).Error
}

// ListDesignRevisions is a method that returns a list of design revisions
func (r *readTxImpl) ListDesignRevisions(ctx context.Context, designID uuid.UUID, pagination *tables.Results[models.DesignRevision]) error {
	return r.conn.Where("design_id = ?", designID).Find(&pagination.Rows).Error
}

// ListLensAnswers is a method that returns a list of lens answers
func (r *readTxImpl) ListLensAnswers(ctx context.Context, lensID uuid.UUID, answers *[]models.WorkloadLensQuestionAnswer) error {
	return r.conn.
		Preload(clause.Associations).
		Where("lens_id = ?", lensID).
		Find(answers).
		Error
}

type writeTxImpl struct {
	conn *gorm.DB
	readTxImpl
}

// NewWriteTx ...
func NewWriteTx() dbx.ReadWriteTxFactory[ports.ReadWriteTx] {
	return func(db *gorm.DB) (ports.ReadWriteTx, error) {
		return &writeTxImpl{conn: db}, nil
	}
}

// CreateTemplate is a method that creates a template
func (rw *writeTxImpl) CreateTemplate(ctx context.Context, template *models.Template) error {
	return rw.conn.Create(template).Error
}

// DeleteTemplate is a method that deletes a template
func (rw *writeTxImpl) DeleteTemplate(ctx context.Context, template *models.Template) error {
	return rw.conn.Delete(template).Error
}

// UpdateTemplate is a method that updates a template
func (rw *writeTxImpl) UpdateTemplate(ctx context.Context, template *models.Template) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(template).Error
}

// CreateDesign is a method that creates a design
func (rw *writeTxImpl) CreateDesign(ctx context.Context, design *models.Design) error {
	for i, tag := range design.Tags { // TODO: This is a workaround for a bug in GORM
		err := rw.conn.Debug().Where(&models.Tag{Name: tag.Name, Value: tag.Value}).FirstOrCreate(&design.Tags[i]).Error
		if err != nil {
			return err
		}
	}

	if design.Workable != nil {
		workflow := models.Workflow{}
		err := rw.conn.Debug().Preload(clause.Associations).First(&workflow, design.Workable.WorkflowTransition.WorkflowID).Error
		if err != nil {
			return err
		}

		states := workflow.GetStates()

		design.Workable.WorkflowTransition.CurrentStateID = states[0].ID
	}

	return rw.conn.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Create(design).Error
}

// DeleteDesign is a method that deletes a design
func (rw *writeTxImpl) DeleteDesign(ctx context.Context, design *models.Design) error {
	return rw.conn.Delete(design).Error
}

// UpdateDesign is a method that updates a design
func (rw *writeTxImpl) UpdateDesign(ctx context.Context, design *models.Design) error {
	src := models.Design{}
	src.ID = design.ID

	err := rw.conn.Where(src).First(&src).Error
	if err != nil {
		return err
	}

	tg := models.DesignRevision{
		DesignID: src.ID,
		Title:    src.Title,
		Body:     src.Body,
	}

	err = rw.conn.Create(&tg).Error
	if err != nil {
		return err
	}

	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(design).Error
}

// AddTagDesign is a method that adds a tag to a design
func (rw *writeTxImpl) AddTagDesign(ctx context.Context, designId uuid.UUID, tag *models.Tag) error {
	design := models.Design{}

	err := rw.conn.Debug().Preload(clause.Associations).First(&design, designId).Error
	if err != nil {
		return err
	}

	err = rw.conn.Debug().Where(&models.Tag{Name: tag.Name, Value: tag.Value}).FirstOrCreate(tag).Error
	if err != nil {
		return err
	}

	return rw.conn.Debug().Model(&design).Association("Tags").Append(tag)
}

// RemoveTagDesign is a method that removes a tag from a design
func (rw *writeTxImpl) RemoveTagDesign(ctx context.Context, designId uuid.UUID, tag *models.Tag) error {
	design := models.Design{}

	err := rw.conn.Debug().Preload(clause.Associations).First(&design, designId).Error
	if err != nil {
		return err
	}

	return rw.conn.Debug().Model(&design).Association("Tags").Delete(tag)
}

// CreateDesignComment is a method that creates a design comment
func (rw *writeTxImpl) CreateDesignComment(ctx context.Context, comment *models.DesignComment) error {
	return rw.conn.Preload(clause.Associations).Create(comment).Error
}

// CreateProfile is a method that creates a profile
func (rw *writeTxImpl) CreateProfile(ctx context.Context, profile *models.Profile) error {
	return rw.conn.Create(profile).Error
}

// UpdateProfile is a method that updates a profile
func (rw *writeTxImpl) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(profile).Error
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
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(environment).Error
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
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(workload).Error
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

// CreateTag is a method that creates a tag
func (rw *writeTxImpl) CreateTag(ctx context.Context, tag *models.Tag) error {
	return rw.conn.Create(tag).Error
}

// UpdateTag is a method that updates a tag
func (rw *writeTxImpl) UpdateTag(ctx context.Context, tag *models.Tag) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(tag).Error
}

// DeleteTag is a method that deletes a tag
func (rw *writeTxImpl) DeleteTag(ctx context.Context, tag *models.Tag) error {
	return rw.conn.Delete(tag).Error
}

// CreateWorkflow is a method that creates a workflow
func (rw *writeTxImpl) CreateWorkflow(ctx context.Context, workflow *models.Workflow) error {
	return rw.conn.Create(workflow).Error
}

// UpdateWorkflow is a method that updates a workflow
func (rw *writeTxImpl) UpdateWorkflow(ctx context.Context, workflow *models.Workflow) error {
	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(workflow).Error
}

// DeleteWorkflow is a method that deletes a workflow
func (rw *writeTxImpl) DeleteWorkflow(ctx context.Context, workflow *models.Workflow) error {
	return rw.conn.Delete(workflow).Error
}

// CreateReaction is a method that creates a reaction
func (rw *writeTxImpl) CreateReaction(ctx context.Context, reaction *models.Reaction) error {
	return rw.conn.Create(reaction).Error
}

// DeleteReaction is a method that deletes a reaction
func (rw *writeTxImpl) DeleteReaction(ctx context.Context, reaction *models.Reaction) error {
	return rw.conn.Delete(reaction, reaction.ID).Error
}

// CreateWorkflowState is a method that adds a workflow state
func (rw *writeTxImpl) CreateWorkflowState(ctx context.Context, state *models.WorkflowState) error {
	workflow := models.Workflow{}

	err := rw.conn.Preload(clause.Associations).First(&workflow, state.WorkflowID).Error
	if err != nil {
		return err
	}

	state.Order = len(workflow.States) - 1
	workflow.States = append(workflow.States, cast.Value(state))

	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&workflow).Error
}

// DeleteWorkflowState is a method that deletes a workflow state
func (rw *writeTxImpl) DeleteWorkflowState(ctx context.Context, state *models.WorkflowState) error {
	err := rw.conn.Debug().Delete(state, state.ID).Error
	if err != nil {
		return err
	}

	workflow := models.Workflow{}
	err = rw.conn.Debug().Preload(clause.Associations).First(&workflow, state.WorkflowID).Error
	if err != nil {
		return err
	}

	slices.SortFunc(workflow.States, func(i, j models.WorkflowState) int {
		return i.Order - j.Order
	})

	for i := range workflow.States {
		workflow.States[i].Order = i
	}

	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&workflow).Error
}

// UpdateWorkflowTransitions is a method that updates workflow transitions
func (rw *writeTxImpl) UpdateWorkflowStateOrder(ctx context.Context, workflowId uuid.UUID, transitions []int) error {
	workflow := models.Workflow{}

	err := rw.conn.Preload(clause.Associations).First(&workflow, workflowId).Error
	if err != nil {
		return err
	}

	for i := range transitions {
		for j := range workflow.States {
			if workflow.States[j].ID == transitions[i] {
				workflow.States[j].Order = i
			}
		}
	}

	return rw.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&workflow.States).Error
}

// DeleteDesignComment ...
func (rw *writeTxImpl) DeleteDesignComment(ctx context.Context, comment *models.DesignComment) error {
	return rw.conn.Debug().Delete(comment, comment.ID).Error
}

// PublishLens ...
func (rw *writeTxImpl) PublishLens(ctx context.Context, lensID uuid.UUID) error {
	return rw.conn.Debug().Model(&models.Lens{}).Where("id = ?", lensID).Update("is_draft", false).Error
}

// UnpublishLens ...
func (rw *writeTxImpl) UnpublishLens(ctx context.Context, lensID uuid.UUID) error {
	return rw.conn.Debug().Model(&models.Lens{}).Where("id = ?", lensID).Update("is_draft", true).Error
}

// AddTagWorkload ...
func (rw *writeTxImpl) AddTagWorkload(ctx context.Context, workloadId uuid.UUID, tag *models.Tag) error {
	workload := models.Workload{}

	err := rw.conn.Debug().Preload(clause.Associations).First(&workload, workloadId).Error
	if err != nil {
		return err
	}

	err = rw.conn.Debug().Where(&models.Tag{Name: tag.Name, Value: tag.Value}).FirstOrCreate(tag).Error
	if err != nil {
		return err
	}

	return rw.conn.Debug().Model(&workload).Association("Tags").Append(tag)
}

// RemoveTagWorkload ...
func (rw *writeTxImpl) RemoveTagWorkload(ctx context.Context, workloadId uuid.UUID, tag *models.Tag) error {
	workload := models.Workload{}

	err := rw.conn.Debug().Preload(clause.Associations).First(&workload, workloadId).Error
	if err != nil {
		return err
	}

	return rw.conn.Debug().Model(&workload).Association("Tags").Delete(tag)
}
