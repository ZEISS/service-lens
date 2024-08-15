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

// DesignStatsControllerImpl ...
type DesignStatsControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewDesignStatsController ...
func NewDesignStatsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *DesignStatsControllerImpl {
	return &DesignStatsControllerImpl{store: store}
}

// Get ...
func (d *DesignStatsControllerImpl) Get() error {
	return d.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var total int64

				err := d.store.ReadTx(d.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTotalNumberOfDesigns(ctx, &total)
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
