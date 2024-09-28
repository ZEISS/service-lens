package dashboard

import (
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/dashboard"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/loading"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
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

// Error ...
func (d *ShowDashboardController) Error(err error) error {
	return toasts.Error(err.Error())
}

// Get ...
func (d *ShowDashboardController) Get() error {
	return d.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        d.Path(),
				User:        d.Session().User,
				Development: d.IsDevelopment(),
			},
			func() htmx.Node {
				return htmx.Fragment(
					dashboard.WelcomeCard(
						dashboard.WelcomeCardProps{},
					),
					stats.Stats(
						stats.StatsProps{
							ClassNames: htmx.ClassNames{
								tailwind.Shadow: false,
								tailwind.M2:     true,
							},
						},
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Designs"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.HxGet(utils.DashboardStatsDesignUrlFormat),
								htmx.HxTrigger("load"),
								loading.Spinner(
									loading.SpinnerProps{},
								),
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
								htmx.HxGet(utils.DashboardStatsProfileUrlFormat),
								htmx.HxTrigger("load"),
								loading.Spinner(
									loading.SpinnerProps{},
								),
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
								htmx.HxGet(utils.DashboardStatsWorkloadUrlFormat),
								htmx.HxTrigger("load"),
								loading.Spinner(
									loading.SpinnerProps{},
								),
							),
						),
					),
				)
			},
		),
	)
}
