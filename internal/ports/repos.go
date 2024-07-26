package ports

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
)

// Migration is a method that runs the migration.
type Migration interface {
	// Migrate is a method that runs the migration.
	Migrate(context.Context) error
}

// Datastore provides methods for transactional operations.
type Datastore interface {
	// ReadTx starts a read only transaction.
	ReadTx(context.Context, func(context.Context, ReadTx) error) error
	// ReadWriteTx starts a read write transaction.
	ReadWriteTx(context.Context, func(context.Context, ReadWriteTx) error) error

	io.Closer
	Migration
}

// ReadTx provides methods for transactional read operations.
type ReadTx interface {
	// GetUser is a method that returns the profile of the current user
	GetUser(ctx context.Context, user *adapters.GothUser) error
	// GetDesign is a method that returns a design by ID
	GetDesign(ctx context.Context, design *models.Design) error
	// ListDesigns is a method that returns a list of designs
	ListDesigns(ctx context.Context, designs *tables.Results[models.Design]) error
	// ListProfiles is a method that returns a list of profiles
	ListProfiles(ctx context.Context, team uuid.UUID, profiles *tables.Results[models.Profile]) error
	// ListProfileQuestions is a method that returns a list of profile questions
	ListProfileQuestions(ctx context.Context, questions *tables.Results[models.ProfileQuestion]) error
	// GetProfile is a method that returns a profile by ID
	GetProfile(ctx context.Context, profile *models.Profile) error
	// ListEnvironments is a method that returns a list of environments
	ListEnvironments(ctx context.Context, environments *tables.Results[models.Environment]) error
	// GetEnvironment is a method that returns an environment by ID
	GetEnvironment(ctx context.Context, environment *models.Environment) error
	// ListLenses is a method that returns a list of lenses
	ListLenses(ctx context.Context, lenses *tables.Results[models.Lens]) error
	// GetLens is a method that returns a lens by ID
	GetLens(ctx context.Context, lens *models.Lens) error
	// ListWorkloads is a method that returns a list of workloads
	ListWorkloads(ctx context.Context, workloads *tables.Results[models.Workload]) error
	// GetWorkload is a method that returns a workload by ID
	GetWorkload(ctx context.Context, workload *models.Workload) error
	// GetLensQuestion is a method that returns a lens question by ID
	GetLensQuestion(ctx context.Context, question *models.Question) error
	// GetWorkloadAnswer is a method that returns a workload answer by ID
	GetWorkloadAnswer(ctx context.Context, answer *models.WorkloadLensQuestionAnswer) error
	// ListTeams is a method that returns a list of teams
	ListTeams(ctx context.Context, teams *tables.Results[adapters.GothTeam]) error
	// GetTeam is a method that returns a team by ID
	GetTeam(ctx context.Context, team *adapters.GothTeam) error
	// ListTags is a method that returns a list of tags
	ListTags(ctx context.Context, tags *tables.Results[models.Tag]) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	ReadTx

	// CreateDesign is a method that creates a design
	CreateDesign(ctx context.Context, design *models.Design) error
	// UpdateDesign is a method that updates a design
	UpdateDesign(ctx context.Context, design *models.Design) error
	// CreateDesignComment is a method that creates a design comment
	CreateDesignComment(ctx context.Context, comment *models.DesignComment) error
	// CreateProfile is a method that creates a profile
	CreateProfile(ctx context.Context, profile *models.Profile) error
	// UpdateProfile is a method that updates a profile
	UpdateProfile(ctx context.Context, profile *models.Profile) error
	// DeleteProfile is a method that deletes a profile
	DeleteProfile(ctx context.Context, profile *models.Profile) error
	// CreateEnvironment is a method that creates an environment
	CreateEnvironment(ctx context.Context, environment *models.Environment) error
	// UpdateEnvironment is a method that updates an environment
	UpdateEnvironment(ctx context.Context, environment *models.Environment) error
	// DeleteEnvironment is a method that deletes an environment
	DeleteEnvironment(ctx context.Context, environment *models.Environment) error
	// CreateLens is a method that creates a lens
	CreateLens(ctx context.Context, lens *models.Lens) error
	// UpdateLens is a method that updates a lens
	UpdateLens(ctx context.Context, lens *models.Lens) error
	// DeleteLens is a method that deletes a lens
	DeleteLens(ctx context.Context, lens *models.Lens) error
	// CreateWorkload is a method that creates a workload
	CreateWorkload(ctx context.Context, workload *models.Workload) error
	// UpdateWorkload is a method that updates a workload
	UpdateWorkload(ctx context.Context, workload *models.Workload) error
	// DeleteWorkload is a method that deletes a workload
	DeleteWorkload(ctx context.Context, workload *models.Workload) error
	// UpdateWorkloadAnswer is a method that updates a workload answer
	UpdateWorkloadAnswer(ctx context.Context, answer *models.WorkloadLensQuestionAnswer) error
	// CreateTeam is a method that creates a team
	CreateTeam(ctx context.Context, team *adapters.GothTeam) error
	// UpdateTeam is a method that updates a team
	UpdateTeam(ctx context.Context, team *adapters.GothTeam) error
	// DeleteTeam is a method that deletes a team
	DeleteTeam(ctx context.Context, team *adapters.GothTeam) error
	// CreateTag is a method that creates a tag
	CreateTag(ctx context.Context, tag *models.Tag) error
	// UpdateTag is a method that updates a tag
	UpdateTag(ctx context.Context, tag *models.Tag) error
	// DeleteTag is a method that deletes a tag
	DeleteTag(ctx context.Context, tag *models.Tag) error
}
