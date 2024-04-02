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

// WorkloadLensEditControllerGetParams ...
type WorkloadLensEditControllerGetParams struct {
	ID       uuid.UUID `json:"id" xml:"id" form:"id"`
	Team     string    `json:"team" xml:"team" form:"team"`
	Lens     uuid.UUID `json:"lens" xml:"lens" form:"lens"`
	Question int       `json:"question" xml:"question" form:"question"`
}

// WorkloadLensEditController ...
type WorkloadLensEditController struct {
	db       ports.Repository
	team     *authz.Team
	lens     *models.Lens
	question models.Question
	answers  *models.WorkloadLensQuestionAnswer

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

	params := &WorkloadLensEditControllerGetParams{}
	if err := hx.Context().ParamsParser(params); err != nil {
		return nil
	}

	lens, err := w.db.GetLensByID(hx.Context().Context(), team.Slug, params.Lens)
	if err != nil {
		return err
	}
	w.lens = lens

	answers, err := w.db.ListAnswers(hx.Context().Context(), params.ID, params.Lens, params.Question)
	if err != nil {
		return err
	}
	w.answers = answers

	for _, pillar := range lens.Pillars {
		for _, question := range pillar.Questions {
			if question.ID == params.Question {
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
									Answer:   w.answers,
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

// Post ...
func (w *WorkloadLensEditController) Post() error {
	return nil
}

// EditFormProps ...
type EditFormProps struct {
	Question models.Question
	Answer   *models.WorkloadLensQuestionAnswer
}

// EditFormComponent ...
func EditFormComponent(p EditFormProps) htmx.Node {
	choices := make([]htmx.Node, len(p.Question.Choices))

	for _, choice := range p.Question.Choices {
		var checked bool
		for _, answer := range p.Answer.Choices {
			if answer.ID == choice.ID {
				checked = true
			}
		}

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
						Name:    choice.Ref,
						Value:   choice.Ref,
						Checked: checked,
					},
					htmx.Text(choice.Title),
				),
			),
		)

		choices = append(choices, input)
	}

	return htmx.Form(
		htmx.Method("POST"),
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
			buttons.ButtonProps{
				Type: "submit",
			},
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
