package ports

import (
	"context"
	"io"

	"github.com/zeiss/fiber-goth/adapters"
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
	// GetProfile is a method that returns the profile of the current user
	GetProfile(ctx context.Context, user *adapters.GothUser) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	ReadTx
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
