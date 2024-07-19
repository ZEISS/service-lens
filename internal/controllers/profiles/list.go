package profiles

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/profiles"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// ProfileListControllerImpl ...
type ProfileListControllerImpl struct {
	profiles tables.Results[models.Profile]
	team     adapters.GothTeam
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewProfilesListController ...
func NewProfilesListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProfileListControllerImpl {
	return &ProfileListControllerImpl{store: store}
}

// Prepare ...
func (w *ProfileListControllerImpl) Prepare() error {
	if err := w.BindQuery(&w.profiles); err != nil {
		return err
	}

	w.team = w.Session().User.TeamBySlug(w.Ctx().Params("t_slug"))

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfiles(ctx, w.team.ID, &w.profiles)
	})
}

// Get ...
func (w *ProfileListControllerImpl) Get() error {
	return w.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: w.Path(),
				},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						profiles.ProfilesTable(
							profiles.ProfilesTableProps{
								Team:     w.team.Slug,
								Profiles: w.profiles.GetRows(),
								Offset:   w.profiles.GetOffset(),
								Limit:    w.profiles.GetLimit(),
								Total:    w.profiles.GetTotalRows(),
							},
						),
					),
				),
			),
		),
	)
}
