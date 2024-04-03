package workloads

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/collapsible"
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

	var noneOfThese bool
	for _, choice := range p.Question.Choices {
		if choice.Ref == "none_of_these" {
			noneOfThese = true
		}
	}

	for _, choice := range p.Question.Choices {
		var checked bool
		for _, answer := range p.Answer.Choices {
			if answer.ID == choice.ID {
				checked = true
			}
		}

		checkbox := CheckboxComponent(
			CheckboxProps{
				Title:    choice.Title,
				Value:    strconv.Itoa(choice.ID),
				Checked:  checked,
				Disabled: (noneOfThese && choice.Ref != "none_of_these") || p.Answer.DoesNotApply,
			},
		)

		choices = append(choices, checkbox)
	}

	return htmx.Form(
		htmx.HxPost(""),
		htmx.HxSwap("none"),
		htmx.Group(choices...),
		DoesNotApplyComponent(
			DoesNotApplyProps{
				Checked: p.Answer.DoesNotApply,
			},
		),
		AdditionalInformationComponent(
			AdditionalInformationProps{
				Description: p.Question.Description,
			},
		),
		forms.TextareaBordered(
			forms.TextareaProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
				Placeholder: "Optional notes",
				Name:        "notes",
			},
			htmx.Text(p.Answer.Notes),
		),
		buttons.OutlinePrimary(
			buttons.ButtonProps{
				Type: "submit",
			},

			htmx.Text("Save"),
		),
	)
}

// CheckboxProps ...
type CheckboxProps struct {
	Title    string
	Value    string
	Checked  bool
	Ref      string
	Disabled bool
}

// CheckboxComponent ...
func CheckboxComponent(p CheckboxProps) htmx.Node {
	return forms.FormControl(
		forms.FormControlProps{},
		forms.FormControlLabel(
			forms.FormControlLabelProps{},
			forms.FormControlLabelText(
				forms.FormControlLabelTextProps{},
				htmx.Text(p.Title),
			),
			forms.Checkbox(
				forms.CheckboxProps{
					Name:     "choices",
					Value:    p.Value,
					Checked:  p.Checked,
					Disabled: p.Disabled,
				},
			),
		),
	)
}

// DoesNotApplyProps ...
type DoesNotApplyProps struct {
	Checked bool
}

// DoesNotApplyComponent ...
func DoesNotApplyComponent(p DoesNotApplyProps) htmx.Node {
	return forms.FormControl(
		forms.FormControlProps{},
		forms.FormControlLabel(
			forms.FormControlLabelProps{},
			forms.FormControlLabelText(
				forms.FormControlLabelTextProps{},
				htmx.Text("Question does not apply to this workload"),
			),
			forms.Toggle(
				forms.ToggleProps{
					Name:    "does_not_apply",
					Value:   "1",
					Checked: p.Checked,
				},
				htmx.HyperScript("on change if me.checked set disabled of <input[type=checkbox][name=choices]/> to true else set disabled of <input[type=checkbox][name=choices]/> to false"),
			),
		),
	)
}

// AdditionalInformationProps ...
type AdditionalInformationProps struct {
	Description string
}

// AdditionalInformationComponent ...
func AdditionalInformationComponent(p AdditionalInformationProps) htmx.Node {
	return collapsible.CollapseArrow(
		collapsible.CollapseProps{},
		collapsible.CollapseCheckbox(
			collapsible.CollapseCheckboxProps{},
		),
		collapsible.CollapseTitle(
			collapsible.CollapseTitleProps{},
			htmx.Text("Additional Information"),
		),
		collapsible.CollapseContent(
			collapsible.CollapseContentProps{},
			htmx.Text(p.Description),
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
