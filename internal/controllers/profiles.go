package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

var _ ports.Profiles = (*Profiles)(nil)

// Profiles ...
type Profiles struct {
	profiles ports.Profiles
}

// NewProfilesController ...
func NewProfilesController(profiles ports.Profiles) *Profiles {
	return &Profiles{profiles}
}

// FetchProfile ...
func (p *Profiles) FetchProfile(ctx context.Context, id uuid.UUID) (*models.Profile, error) {
	return p.profiles.FetchProfile(ctx, id)
}

// NewProfile ...
func (p *Profiles) NewProfile(ctx context.Context, profile *models.Profile) error {
	return p.profiles.NewProfile(ctx, profile)
}

// ListProfiles ...
func (p *Profiles) ListProfiles(ctx context.Context, pagination *models.Pagination) ([]*models.Profile, error) {
	return p.profiles.ListProfiles(ctx, pagination)
}
