package ports

import (
	"context"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// Users ...
type Users interface {
	// GetUserByID ...
	GetUserByID(ctx context.Context, id uuid.UUID) (*authz.User, error)
	// ListUsers ...
	ListUsers(ctx context.Context, pagination tables.Results[*authz.User]) (*tables.Results[*authz.User], error)
	// AddUser ...
	AddUser(ctx context.Context, user *authz.User) (*authz.User, error)
	// UpdateUser ...
	UpdateUser(ctx context.Context, user *authz.User) (*authz.User, error)
	// DeleteUser ...
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
