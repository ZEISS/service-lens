package ports

import (
	"context"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/models"
)

// User ...
type User interface {
	// GetUserByID ...
	GetUserByID(ctx context.Context, id uuid.UUID) (*authz.User, error)
	// ListUsers ...
	ListUsers(ctx context.Context, pagination *models.Pagination) ([]*authz.User, error)
	// AddUser ...
	AddUser(ctx context.Context, user *authz.User) (*authz.User, error)
	// UpdateUser ...
	UpdateUser(ctx context.Context, user *authz.User) (*authz.User, error)
	// DeleteUser ...
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
