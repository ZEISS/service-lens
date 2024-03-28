package lenses

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// LensIndexController ...
type LensIndexController struct {
	db   ports.Repository
	lens *models.Lens

	htmx.UnimplementedController
}

// NewLensIndexController ...
func NewLensIndexController(db ports.Repository) *LensIndexController {
	return &LensIndexController{
		db: db,
	}
}

// Prepare ...
func (l *LensIndexController) Prepare() error {
	// id, err := uuid.Parse(l.Hx().Context().Params("id"))
	// if err != nil {
	// 	return err
	// }

	return nil
}

// Get ...
func (l *LensIndexController) Get() error {
	hx := l.Hx()

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
				),
			),
		),
	)
}
