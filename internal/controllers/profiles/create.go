package profiles

import (
	"context"
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

const (
	listProfilesURL = "/teams/%s/profiles"
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

// Error ...
func (l *CreateProfileControllerImpl) Error(err error) error {
	fmt.Println(err)
	return err
}

// Prepare ...
func (l *CreateProfileControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.profile)
	if err != nil {
		return err
	}

	team := utils.FromContextTeam(l.Ctx())
	l.profile.TeamID = team.ID

	err = validate.Struct(&l.profile)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateProfile(ctx, &l.profile)
	})
}

// Post ...
func (l *CreateProfileControllerImpl) Post() error {
	return l.Redirect(fmt.Sprintf(listProfilesURL, utils.FromContextTeam(l.Ctx()).Slug))
}
