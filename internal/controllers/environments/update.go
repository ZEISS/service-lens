package environments

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// EnvironmentUpdateControllerImpl ...
type EnvironmentUpdateControllerImpl struct {
	environment models.Environment
	store       seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewEnvironmentUpdateController ...
func NewEnvironmentUpdateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EnvironmentUpdateControllerImpl {
	return &EnvironmentUpdateControllerImpl{
		environment: models.Environment{},
		store:       store,
	}
}

// Prepare ...
func (p *EnvironmentUpdateControllerImpl) Prepare() error {
	validate = validator.New()

	err := p.BindParams(&p.environment)
	if err != nil {
		return err
	}

	err = p.BindBody(&p.environment)
	if err != nil {
		return err
	}

	err = validate.Struct(&p.environment)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateEnvironment(ctx, &p.environment)
	})
}

// Put ...
func (p *EnvironmentUpdateControllerImpl) Put() error {
	return p.Redirect(fmt.Sprintf("/environments/%s", p.environment.ID))
}
