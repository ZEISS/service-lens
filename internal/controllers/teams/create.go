package teams

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

const (
	listTeamURL = "/site/teams"
)

var validate *validator.Validate

// CreateTeamControllerImpl ...
type CreateTeamControllerImpl struct {
	team  adapters.GothTeam
	store ports.Datastore
	htmx.DefaultController
}

// NewCreateTeamController ...
func NewCreateTeamController(store ports.Datastore) *CreateTeamControllerImpl {
	return &CreateTeamControllerImpl{
		store: store,
	}
}

// Error ...
func (l *CreateTeamControllerImpl) Error(err error) error {
	fmt.Println()
	return err
}

// Prepare ...
func (l *CreateTeamControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.team)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.team)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateTeam(ctx, &l.team)
	})
}

// Post ...
func (l *CreateTeamControllerImpl) Post() error {
	return l.Redirect(listTeamURL)
}
