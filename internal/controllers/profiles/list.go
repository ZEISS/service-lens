package profiles

import (
	"context"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/profiles"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// ProfileListControllerImpl ...
type ProfileListControllerImpl struct {
	profiles tables.Results[models.Profile]
	store    ports.Datastore
	htmx.DefaultController
}

// NewProfilesListController ...
func NewProfilesListController(store ports.Datastore) *ProfileListControllerImpl {
	return &ProfileListControllerImpl{store: store}
}

// Prepare ...
func (w *ProfileListControllerImpl) Prepare() error {
	if err := w.BindQuery(&w.profiles); err != nil {
		return err
	}

	team := utils.FromContextTeam(w.Ctx())

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfiles(ctx, team.ID, &w.profiles)
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
								Team:     utils.FromContextTeam(w.Ctx()).Slug,
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
