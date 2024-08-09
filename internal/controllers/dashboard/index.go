package dashboard

import (
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/stats"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// ShowDashboardController ...
type ShowDashboardController struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewShowDashboardController ...
func NewShowDashboardController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowDashboardController {
	return &ShowDashboardController{
		store: store,
	}
}

// Get ...
func (d *ShowDashboardController) Get() error {
	return d.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: d.Path(),
				User: d.Session().User,
			},
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"m-2": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text("Dashboard"),
					),
					stats.Stats(
						stats.StatsProps{},
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Designs"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text("0"),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Profiles"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text("0"),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Workloads"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text("0"),
							),
						),
					),
				),
			),
		),
	)
}
