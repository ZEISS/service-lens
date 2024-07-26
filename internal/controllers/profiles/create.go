package profiles

import (
	"context"
	"fmt"

	goth "github.com/zeiss/fiber-goth"
	"github.com/zeiss/fiber-goth/adapters"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

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
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	team    adapters.GothTeam
	htmx.DefaultController
}

// NewCreateProfileController ...
func NewCreateProfileController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateProfileControllerImpl {
	return &CreateProfileControllerImpl{
		profile: models.Profile{},
		store:   store,
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

	session, err := goth.SessionFromContext(l.Ctx())
	if err != nil {
		return err
	}
	l.team = session.User.TeamBySlug(l.Ctx().Params("t_slug"))

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
	return l.Redirect(fmt.Sprintf(listProfilesURL, l.team.ID))
}
