package teams

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/teams"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// TeamListControllerImpl ...
type TeamListControllerImpl struct {
	teams tables.Results[adapters.GothTeam]
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTeamListController ...
func NewTeamListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TeamListControllerImpl {
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
		components.DefaultLayout(
			components.DefaultLayoutProps{},
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"m-2": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
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
	)
}
