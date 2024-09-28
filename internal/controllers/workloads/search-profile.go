package workloads

import (
	"context"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/conv"
)

var _ = htmx.Controller(&SearchProfilesControllerImpl{})

// Search ...
type SearchProfilesControllerImpl struct {
	profiles tables.Results[models.Profile]
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewSearchProfilesController ...
func NewSearchProfilesController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *SearchProfilesControllerImpl {
	return &SearchProfilesControllerImpl{
		profiles: tables.Results[models.Profile]{SearchFields: []string{"name"}},
		store:    store,
	}
}

// Error ...
func (l *SearchProfilesControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *SearchProfilesControllerImpl) Prepare() error {
	var params struct {
		Profile string `json:"profile" form:"profile" query:"profile" validate:"required"`
	}

	err := l.BindQuery(&params)
	if err != nil {
		return err
	}
	l.profiles.Search = params.Profile

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfiles(ctx, &l.profiles)
	})
}

// Get ...
func (l *SearchProfilesControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.ForEach(l.profiles.GetRows(), func(e *models.Profile, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(conv.String(e.ID)),
					htmx.Text(e.Name),
				)
			})...,
		),
	)
}
