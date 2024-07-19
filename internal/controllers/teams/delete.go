package teams

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/ports"
)

// TeamDeleteControllerImpl ...
type TeamDeleteControllerImpl struct {
	team  adapters.GothTeam
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTeamDeleteController ...
func NewTeamDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TeamDeleteControllerImpl {
	return &TeamDeleteControllerImpl{
		team:  adapters.GothTeam{},
		store: store,
	}
}

// Prepare ...
func (p *TeamDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.team)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTeam(ctx, &p.team)
	})
}

// Delete ...
func (p *TeamDeleteControllerImpl) Delete() error {
	return p.Redirect("/site/teams")
}
