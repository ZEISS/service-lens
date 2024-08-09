package teams

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/teams"
	"github.com/zeiss/service-lens/internal/ports"
)

// NewTeamControllerImpl ...
type NewTeamControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTeamController ...
func NewTeamController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewTeamControllerImpl {
	return &NewTeamControllerImpl{store: store}
}

// Error ...
func (p *NewTeamControllerImpl) Error(err error) error {
	return toasts.RenderToasts(
		p.Ctx(),
		toasts.Toasts(
			toasts.ToastsProps{},
			toasts.ToastAlertError(
				toasts.ToastProps{},
				htmx.Text(err.Error()),
			),
		),
	)
}

// New ...
func (p *NewTeamControllerImpl) Get() error {
	return p.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				User: p.Session().User,
			},
			teams.NewForm(
				teams.NewFormProps{},
			),
		),
	)
}
