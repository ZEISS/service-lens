package lenses

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// LensIndexControllerParams ...
type LensIndexControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultLensIndexControllerParams ...
func NewDefaultLensIndexControllerParams() *LensIndexControllerParams {
	return &LensIndexControllerParams{}
}

// LensIndexController ...
type LensIndexController struct {
	db     ports.Repository
	lens   *models.Lens
	params *LensIndexControllerParams

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
	l.params = NewDefaultLensIndexControllerParams()
	if err := l.Hx().Ctx().ParamsParser(l.params); err != nil {
		return err
	}

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

// Delete ...
func (l *LensIndexController) Delete() error {
	err := l.db.DestroyLens(l.Hx().Ctx().Context(), l.params.ID)
	if err != nil {
		return err
	}

	return nil
}
