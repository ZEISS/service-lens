package settings

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// SettingsIndexController ...
type SettingsIndexController struct {
	db ports.Repository

	htmx.DefaultController
}

// NewSettingsIndexController ...
func NewSettingsIndexController(db ports.Repository) *SettingsIndexController {
	return &SettingsIndexController{
		db: db,
	}
}

// Prepare ...
func (m *SettingsIndexController) Prepare() error {
	if err := m.BindValues(utils.User(m.db), utils.Team(m.db)); err != nil {
		return err
	}

	return nil
}

// Get ...
func (m *SettingsIndexController) Get() error {
	return m.Hx().RenderComp(
		components.Page(
			m.DefaultCtx(),
			components.PageProps{},
			components.Layout(
				m.DefaultCtx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
				),
			),
		),
	)
}
