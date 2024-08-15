package stats

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/stats"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// ProfileStatsControllerImpl ...
type ProfileStatsControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewProfileStatsController ...
func NewProfileStatsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProfileStatsControllerImpl {
	return &ProfileStatsControllerImpl{store: store}
}

// Get ...
func (d *ProfileStatsControllerImpl) Get() error {
	return d.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var total int64

				err := d.store.ReadTx(d.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTotalNumberOfProfiles(ctx, &total)
				})
				errorx.Panic(err)

				return stats.Value(
					stats.ValueProps{},
					htmx.Text(conv.String(total)),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		),
	)
}
