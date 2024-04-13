package settings

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// SettingsIndexController ...
type SettingsIndexController struct {
	db  ports.Repository
	ctx htmx.Ctx

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
	ctx, err := htmx.NewDefaultContext(m.Hx().Ctx(), utils.Team(m.Hx().Ctx(), m.db), utils.User(m.Hx().Ctx(), m.db))
	if err != nil {
		return err
	}
	m.ctx = ctx

	return nil
}

// Get ...
func (m *SettingsIndexController) Get() error {
	return m.Hx().RenderComp(
		components.Page(
			m.ctx,
			components.PageProps{},
			components.Layout(
				m.ctx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
				),
			),
		),
	)
}
