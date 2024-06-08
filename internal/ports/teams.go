package ports

import (
	"context"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// Teams ...
type Teams interface {
	// GetTeamBySlug ...
	GetTeamBySlug(ctx context.Context, slug string) (*authz.Team, error)
	// GetTeamByID ...
	GetTeamByID(ctx context.Context, id uuid.UUID) (*authz.Team, error)
	// CreateTeam ...
	CreateTeam(ctx context.Context, team *authz.Team, user *authz.User) (*authz.Team, error)
	// ListTeams ...
	ListTeams(ctx context.Context, pagination tables.Results[*authz.Team]) (*tables.Results[*authz.Team], error)
	// DeleteTeam
	DeleteTeam(ctx context.Context, id uuid.UUID) error
	// TotalCountWorkloads returns the total number of workloads.
	TotalCountWorkloads(ctx context.Context, teamSlug string) (int, error)
	// TotalCountLenses returns the total number of lenses.
	TotalCountLenses(ctx context.Context, teamSlug string) (int, error)
	// TotalCountProfiles returns the total number of questions.
	TotalCountProfiles(ctx context.Context, teamSlug string) (int, error)
}
