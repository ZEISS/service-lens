package environments

import (
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/environments"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// EnvironmentControllerImpl ...
type EnvironmentControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewEnvironmentController ...
func NewEnvironmentController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EnvironmentControllerImpl {
	return &EnvironmentControllerImpl{store: store}
}

// New ...
func (p *EnvironmentControllerImpl) Get() error {
	return p.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				User:        p.Session().User,
				Title:       "New Environment",
				Path:        p.Path(),
				Development: p.IsDevelopment(),
			},
			func() htmx.Node {
				return environments.NewForm(
					environments.NewFormProps{},
				)
			},
		),
	)
}
