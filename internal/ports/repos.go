package ports

import (
	"context"
	"io"

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
	// ListProfiles is a method that returns a list of profiles
	ListProfiles(ctx context.Context, profiles *tables.Results[models.Profile]) error
	// GetProfile is a method that returns a profile by ID
	GetProfile(ctx context.Context, profile *models.Profile) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	ReadTx

	// CreateProfile is a method that creates a profile
	CreateProfile(ctx context.Context, profile *models.Profile) error
	// UpdateProfile is a method that updates a profile
	UpdateProfile(ctx context.Context, profile *models.Profile) error
}

// Repository is the interface that wraps the basic methods to interact with the database.
type Repository interface {
	Lenses
	Profiles
	Teams
	Users
	Workloads
	Environments
}
