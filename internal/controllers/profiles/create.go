package profiles

import (
	"context"
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

var validate *validator.Validate

// CreateProfileControllerImpl ...
type CreateProfileControllerImpl struct {
	profile models.Profile
	store   ports.Datastore
	htmx.DefaultController
}

// NewCreateProfileController ...
func NewCreateProfileController(store ports.Datastore) *CreateProfileControllerImpl {
	return &CreateProfileControllerImpl{
		profile:           models.Profile{},
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *CreateProfileControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.profile)
	if err != nil {
		return err
	}

	err = validate.Struct(l)
	if err != nil {

		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateProfile(ctx, &l.profile)
	})
}

// Post ...
func (l *CreateProfileControllerImpl) Post() error {
	return l.Redirect(fmt.Sprintf("/profiles/%s", l.profile.ID))
}
