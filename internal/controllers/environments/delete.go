package environments

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// EnvironmentDeleteControllerImpl ...
type EnvironmentDeleteControllerImpl struct {
	environment models.Environment
	store       ports.Datastore
	htmx.DefaultController
}

// NewEnvironmentDeleteController ...
func NewEnvironmentDeleteController(store ports.Datastore) *EnvironmentDeleteControllerImpl {
	return &EnvironmentDeleteControllerImpl{
		environment: models.Environment{},
		store:       store,
	}
}

// Prepare ...
func (p *EnvironmentDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.environment)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteEnvironment(ctx, &p.environment)
	})
}

// Delete ...
func (p *EnvironmentDeleteControllerImpl) Delete() error {
	return p.Redirect("/environments")
}
