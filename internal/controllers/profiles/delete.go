package profiles

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// ProfileDeleteControllerImpl ...
type ProfileDeleteControllerImpl struct {
	profile models.Profile
	store   ports.Datastore
	htmx.DefaultController
}

// NewProfileDeleteController ...
func NewProfileDeleteController(store ports.Datastore) *ProfileDeleteControllerImpl {
	return &ProfileDeleteControllerImpl{
		profile: models.Profile{},
		store:   store,
	}
}

// Prepare ...
func (p *ProfileDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.profile)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteProfile(ctx, &p.profile)
	})
}

// Delete ...
func (p *ProfileDeleteControllerImpl) Delete() error {
	return p.Redirect("/profiles")
}
