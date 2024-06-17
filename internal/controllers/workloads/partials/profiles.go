package partials

import (
	"context"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// ProfilePartialListControllerImpl ...
type ProfilePartialListControllerImpl struct {
	profiles tables.Results[models.Profile]
	store    ports.Datastore
	htmx.UnimplementedController
}

// NewProfilePartialListController ...
func NewProfilePartialListController(store ports.Datastore) *ProfilePartialListControllerImpl {
	return &ProfilePartialListControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *ProfilePartialListControllerImpl) Prepare() error {
	err := w.BindQuery(&w.profiles)
	if err != nil {
		return err
	}

	team := utils.FromContextTeam(w.Ctx())

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfiles(ctx, team.ID, &w.profiles)
	})
}

// Get ...
func (w *ProfilePartialListControllerImpl) Get() error {
	return w.Render(
		htmx.Fragment(
			htmx.ForEach(w.profiles.GetRows(), func(e *models.Profile, profileIdx int) htmx.Node {
				return dropdowns.DropdownMenuItem(
					dropdowns.DropdownMenuItemProps{},
					htmx.A(
						htmx.Text(e.Name),
						htmx.DataAttribute("profile", e.ID.String()),
						htmx.HyperScript(`on click set (previous <input/>).value to my @data-profile then put my innerHTML into #profiles-button`),
					),
				)
			})...,
		),
	)
}
