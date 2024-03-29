package workloads

import (
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadPillarController ...
type WorkloadPillarController struct {
	db   ports.Repository
	team *authz.Team
	lens *models.Lens

	htmx.UnimplementedController
}

// NewWorkloadLensController ...
func NewWorkloadPillarController(db ports.Repository) *WorkloadPillarController {
	return &WorkloadPillarController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadPillarController) Prepare() error {
	hx := w.Hx()

	team := hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)
	w.team = team

	lensID, err := uuid.Parse(hx.Context().Params("lens"))
	if err != nil {
		return err
	}

	lens, err := w.db.GetLensByID(hx.Context().Context(), team.Slug, lensID)
	if err != nil {
		return err
	}
	w.lens = lens

	return nil
}

// Get ...
func (w *WorkloadPillarController) Get() error {
	hx := w.Hx()

	pillars := make([]htmx.Node, len(w.lens.Pillars))

	for _, pillar := range w.lens.Pillars {
		menuItems := make([]htmx.Node, len(pillar.Questions))

		for _, question := range pillar.Questions {
			menuItem := menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{},
					htmx.Text(question.Title),
				),
			)

			menuItems = append(menuItems, menuItem)
		}

		menu := menus.MenuItem(
			menus.MenuItemProps{},
			menus.MenuCollapsible(
				menus.MenuCollapsibleProps{
					Open: true,
				},
				menus.MenuCollapsibleSummary(
					menus.MenuCollapsibleSummaryProps{},
					htmx.Text(pillar.Name),
				),
				htmx.Group(menuItems...),
			),
		)

		pillars = append(pillars, menu)
	}

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},

					// drawers.Drawer(
					// 	drawers.DrawerProps{
					// 		ID: "pillars-drawer",
					// 		ClassNames: htmx.ClassNames{
					// 			"drawer-open": true,
					// 		},
					// 	},
					// 	drawers.DrawerContent(
					// 		drawers.DrawerContentProps{
					// 			ID: "pillars-drawer",
					// 			ClassNames: htmx.ClassNames{
					// 				"px-8": true,
					// 			},
					// 		},
					// 		htmx.Text("Drawer this is the new content for the drawer"),
					// 	),
					// 	drawers.DrawerSide(
					// 		drawers.DrawerSideProps{},
					// 		menus.Menu(
					// 			menus.MenuProps{
					// 				ClassNames: htmx.ClassNames{
					// 					"w-full": true,
					// 				},
					// 			},
					// 			htmx.Group(pillars...),
					// 		),
					// 	),
					// ),
				),
			),
		),
	)
}
