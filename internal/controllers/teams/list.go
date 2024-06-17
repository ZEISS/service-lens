package teams

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/teams"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// TeamListControllerImpl ...
type TeamListControllerImpl struct {
	teams tables.Results[adapters.GothTeam]
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamListController ...
func NewTeamListController(store ports.Datastore) *TeamListControllerImpl {
	return &TeamListControllerImpl{store: store}
}

// Error ...
func (w *TeamListControllerImpl) Error(err error) error {
	fmt.Println(err)

	return err
}

// Prepare ...
func (w *TeamListControllerImpl) Prepare() error {
	if err := w.BindQuery(&w.teams); err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTeams(ctx, &w.teams)
	})
}

// Get ...
func (w *TeamListControllerImpl) Get() error {
	return w.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: w.Path(),
					User: w.Session().User,
				},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						teams.TeamsTable(
							teams.TeamsTableProps{
								Teams:  w.teams.GetRows(),
								Offset: w.teams.GetOffset(),
								Limit:  w.teams.GetLimit(),
								Total:  w.teams.GetTotalRows(),
							},
						),
					),
				),
			),
		),
	)
}
