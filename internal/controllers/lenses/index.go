package lenses

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
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

	lens, err := l.db.GetLensByID(l.Hx().Ctx().Context(), l.params.ID)
	if err != nil {
		return err
	}
	l.lens = lens

	return nil
}

// Get ...
func (l *LensIndexController) Get() error {
	return l.Hx().RenderComp(
		components.Page(
			l.Hx(),
			components.PageProps{},
			components.Layout(
				l.Hx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Overview"),
							),
							htmx.Div(
								htmx.H1(
									htmx.Text(l.lens.Name),
								),
								htmx.P(
									htmx.Text(l.lens.Description),
								),
								htmx.Div(
									htmx.ClassNames{
										"flex":     true,
										"flex-col": true,
										"py-2":     true,
									},
									htmx.H4(
										htmx.ClassNames{
											"text-gray-500": true,
										},
										htmx.Text("Created at"),
									),
									htmx.H3(
										htmx.Text(
											l.lens.CreatedAt.Format("2006-01-02 15:04:05"),
										),
									),
								),
								htmx.Div(
									htmx.ClassNames{
										"flex":     true,
										"flex-col": true,
										"py-2":     true,
									},
									htmx.H4(
										htmx.ClassNames{
											"text-gray-500": true,
										},
										htmx.Text("Updated at"),
									),
									htmx.H3(
										htmx.Text(
											l.lens.UpdatedAt.Format("2006-01-02 15:04:05"),
										),
									),
								),
							),
						),
					),
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
