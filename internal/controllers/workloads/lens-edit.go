package workloads

import (
	"fmt"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadLensEditController ...
type WorkloadLensEditController struct {
	db       ports.Repository
	team     *authz.Team
	lens     *models.Lens
	question models.Question

	htmx.UnimplementedController
}

// NewWorkloadLensEditController ...
func NewWorkloadLensEditController(db ports.Repository) *WorkloadLensEditController {
	return &WorkloadLensEditController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadLensEditController) Prepare() error {
	hx := w.Hx()

	team := hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)
	w.team = team

	lensID, err := uuid.Parse(hx.Context().Params("lens"))
	if err != nil {
		return err
	}

	questionID, err := hx.Context().ParamsInt("question")
	if err != nil {
		return err
	}

	lens, err := w.db.GetLensByID(hx.Context().Context(), team.Slug, lensID)
	if err != nil {
		return err
	}
	w.lens = lens

	for _, pillar := range lens.Pillars {
		for _, question := range pillar.Questions {
			if question.ID == questionID {
				w.question = question
			}
		}
	}

	return nil
}

// Get ...
func (w *WorkloadLensEditController) Get() error {
	hx := w.Hx()

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					drawers.Drawer(
						drawers.DrawerProps{
							ID: "pillars-drawer",
							ClassNames: htmx.ClassNames{
								"drawer-open": true,
							},
						},
						drawers.DrawerContent(
							drawers.DrawerContentProps{
								ID: "pillars-drawer",
								ClassNames: htmx.ClassNames{
									"px-8": true,
								},
							},
							htmx.H1(htmx.Text(w.question.Title)),
							htmx.Text(w.question.Description),
							EditFormComponent(
								EditFormProps{
									Question: w.question,
								},
							),
						),
						drawers.DrawerSide(
							drawers.DrawerSideProps{},
							EditMenuComponent(
								EditMenuProps{
									Lens: w.lens,
								},
							),
						),
					),
				),
			),
		),
	)
}

// EditFormProps ...
type EditFormProps struct {
	Question models.Question
}

// EditFormComponent ...
func EditFormComponent(p EditFormProps) htmx.Node {
	choices := make([]htmx.Node, len(p.Question.Choices))

	for _, choice := range p.Question.Choices {
		input := forms.FormControl(
			forms.FormControlProps{},
			forms.FormControlLabel(
				forms.FormControlLabelProps{},
				forms.FormControlLabelText(
					forms.FormControlLabelTextProps{},
					htmx.Text(choice.Title),
				),
				forms.Checkbox(
					forms.CheckboxProps{
						Name:  choice.Ref,
						Value: choice.Ref,
					},
					htmx.Text(choice.Title),
				),
			),
		)

		choices = append(choices, input)
	}

	return htmx.Form(
		htmx.Method("PUT"),
		htmx.Group(choices...),
		forms.FormControl(
			forms.FormControlProps{},
			forms.FormControlLabel(
				forms.FormControlLabelProps{},
				forms.FormControlLabelText(
					forms.FormControlLabelTextProps{},
					htmx.Text("Question does not apply to this workload"),
				),
				forms.Toggle(
					forms.ToggleProps{},
				),
			),
		),
		forms.TextareaBordered(
			forms.TextareaProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
				Placeholder: "Opional notes",
			},
		),
		buttons.OutlinePrimary(
			buttons.ButtonProps{},
			htmx.Type("submit"),
			htmx.Text("Save"),
		),
	)
}

// EditMenuProps ...
type EditMenuProps struct {
	Lens *models.Lens
}

// EditMenuComponent ...
func EditMenuComponent(p EditMenuProps) htmx.Node {
	pillars := make([]htmx.Node, len(p.Lens.Pillars))
	for _, pillar := range p.Lens.Pillars {
		questions := make([]htmx.Node, len(pillar.Questions))

		for _, question := range pillar.Questions {
			questions = append(questions, menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href: fmt.Sprintf("%d", question.ID),
					},
					htmx.Text(question.Title),
				),
			))
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
				htmx.Group(questions...),
			),
		)

		pillars = append(pillars, menu)
	}

	return menus.Menu(
		menus.MenuProps{
			ClassNames: htmx.ClassNames{
				"w-full": true,
			},
		},
		htmx.Group(pillars...),
	)
}
