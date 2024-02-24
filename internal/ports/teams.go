package ports

import (
	"context"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/models"
)

// Teams ...
type Teams interface {
	GetTeamBySlug(ctx context.Context, slug string) (*authz.Team, error)
	GetTeamByID(ctx context.Context, id uuid.UUID) (*authz.Team, error)
	AddTeam(ctx context.Context, team *authz.Team) (*authz.Team, error)
	ListTeams(ctx context.Context, pagination *models.Pagination) ([]*authz.Team, error)
}
