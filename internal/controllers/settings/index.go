package settings

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// SettingsIndexController ...
type SettingsIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewSettingsIndexController ...
func NewSettingsIndexController(db ports.Repository) *SettingsIndexController {
	return &SettingsIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (m *SettingsIndexController) Get() error {
	return m.Hx().RenderComp(
		components.Page(
			m.Hx(),
			components.PageProps{},
			components.Layout(
				m.Hx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
				),
			),
		),
	)
}
