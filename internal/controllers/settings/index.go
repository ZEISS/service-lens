package settings

import (
	authz "github.com/zeiss/fiber-authz"
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
	if err := m.BindValues(utils.User(m.db)); err != nil {
		return err
	}

	return nil
}

// Get ...
func (m *SettingsIndexController) Get() error {
	return m.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: m.Values(utils.ValuesKeyUser).(*authz.User),
				},
				components.Wrap(
					components.WrapProps{},
				),
			),
		),
	)
}
