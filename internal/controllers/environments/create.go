package environments

import (
	"context"
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

var validate *validator.Validate

// CreateEnvironmentControllerImpl ...
type CreateEnvironmentControllerImpl struct {
	environment models.Environment
	store       ports.Datastore
	htmx.DefaultController
}

// NewCreateEnvironmentController ...
func NewCreateEnvironmentController(store ports.Datastore) *CreateEnvironmentControllerImpl {
	return &CreateEnvironmentControllerImpl{
		environment:       models.Environment{},
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *CreateEnvironmentControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.environment)
	if err != nil {
		return err
	}

	err = validate.Struct(l)
	if err != nil {

		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateEnvironment(ctx, &l.environment)
	})
}

// Post ...
func (l *CreateEnvironmentControllerImpl) Post() error {
	return l.Redirect(fmt.Sprintf("/environments/%s", l.environment.ID))
}
