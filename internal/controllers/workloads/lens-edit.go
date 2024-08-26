package workloads

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

const (
	showLensQuestionURL = "/workloads/%s/lenses/%s/question/%d"
)

// WorkloadLensEditControllerImpl ...
type WorkloadLensEditControllerImpl struct {
	workload models.Workload
	lens     models.Lens
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadLensEditController ...
func NewWorkloadLensEditController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadLensEditControllerImpl {
	return &WorkloadLensEditControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *WorkloadLensEditControllerImpl) Prepare() error {
	err := w.BindParams(&w.workload)
	if err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		if err := tx.GetWorkload(ctx, &w.workload); err != nil {
			return err
		}

		id, err := uuid.Parse(w.Ctx().Params("lens"))
		if err != nil {
			return err
		}
		w.lens.ID = id

		return tx.GetLens(ctx, &w.lens)
	})
}

// Get ...
func (w *WorkloadLensEditControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       w.lens.Name,
				Path:        w.Path(),
				User:        w.Session().User,
				Development: w.IsDevelopment(),
			},
			func() htmx.Node {
				return drawers.Drawer(
					drawers.DrawerProps{
						ID: "pillars-drawer",
						ClassNames: htmx.ClassNames{
							"drawer-open": true,
						},
					},
					drawers.DrawerContent(
						drawers.DrawerContentProps{},
						htmx.ID("pillars-drawer-content"),
						cards.CardBordered(
							cards.CardProps{
								ClassNames: htmx.ClassNames{
									tailwind.M4: true,
								},
							},
							cards.Body(
								cards.BodyProps{},
								cards.Title(
									cards.TitleProps{},
									htmx.Text(w.lens.Name),
								),
							),
						),
					),
					drawers.DrawerSide(
						drawers.DrawerSideProps{
							ClassNames: htmx.ClassNames{
								tailwind.MinHScreen: true,
								tailwind.HFull:      true,
								"overflow-y-auto":   true,
								"max-w-80":          true,
							},
						},
						htmx.Div(
							htmx.Role("navigation"),
							menus.Menu(
								menus.MenuProps{
									ClassNames: htmx.ClassNames{
										"w-full":      true,
										"bg-base-200": false,
									},
								},
								htmx.Group(
									htmx.ForEach(w.lens.GetPillars(), func(p *models.Pillar, pillarIdx int) htmx.Node {
										return menus.MenuItem(
											menus.MenuItemProps{},
											menus.MenuCollapsible(
												menus.MenuCollapsibleProps{
													Open: pillarIdx == 0,
												},
												menus.MenuCollapsibleSummary(
													menus.MenuCollapsibleSummaryProps{},
													htmx.Text(p.Name),
												),
												htmx.Group(
													htmx.ForEach(p.GetQuestions(), func(q *models.Question, questionIdx int) htmx.Node {
														return menus.MenuItem(
															menus.MenuItemProps{},
															menus.MenuLink(
																menus.MenuLinkProps{},
																htmx.Text(q.Title),
																htmx.HxTarget("#pillars-drawer-content"),
																htmx.HxGet(fmt.Sprintf(showLensQuestionURL, w.workload.ID, w.lens.ID, q.ID)),
																htmx.HxSwap("innerHTML"),
															),
														)
													})...,
												),
											),
										)
									})...,
								),
							),
						),
					),
				)
			},
		),
	)
}
