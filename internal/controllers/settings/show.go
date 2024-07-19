package settings

import (
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// SettingsShowControllerImpl ...
type SettingsShowControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewSettingsShowController ...
func NewSettingsShowController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *SettingsShowControllerImpl {
	return &SettingsShowControllerImpl{
		store: store,
	}
}

// Prepare ...
func (m *SettingsShowControllerImpl) Prepare() error {
	return nil
}

// Get ...
func (m *SettingsShowControllerImpl) Get() error {
	return m.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: m.Path(),
				},
				components.Wrap(
					components.WrapProps{},
				),
			),
		),
	)
}
