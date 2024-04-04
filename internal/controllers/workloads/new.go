package workloads

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadNewController ...
type WorkloadNewController struct {
	db   ports.Repository
	team *authz.Team

	htmx.UnimplementedController
}

// NewWorkloadsNewController ...
func NewWorkloadsNewController(db ports.Repository) *WorkloadNewController {
	return &WorkloadNewController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadNewController) Prepare() error {
	team := w.Hx().Values(resolvers.ValuesKeyTeam).(*authz.Team)
	w.team = team

	return nil
}

// Post ...
func (w *WorkloadNewController) Post() error {
	hx := w.Hx()

	profileId, err := uuid.Parse(hx.Ctx().FormValue("profile"))
	if err != nil {
		return err
	}

	lensId, err := uuid.Parse(hx.Ctx().FormValue("lens"))
	if err != nil {
		return err
	}

	workload := &models.Workload{
		Name:        hx.Ctx().FormValue("name"),
		Description: hx.Ctx().FormValue("description"),
		Team:        *w.team,
		ProfileID:   profileId,
		Lenses:      []*models.Lens{{ID: lensId}},
	}

	err = w.db.StoreWorkload(hx.Ctx().Context(), workload)
	if err != nil {
		return err
	}

	hx.Redirect(fmt.Sprintf("/%s/workloads/%s", w.team.Slug, workload.ID))

	return nil
}

// Get ...
func (w *WorkloadNewController) Get() error {
	hx := w.Hx()

	profiles, err := w.db.ListProfiles(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return err
	}

	profilesItems := make([]htmx.Node, len(profiles))
	for i, profile := range profiles {
		profilesItems[i] = htmx.Option(
			htmx.Attribute("value", profile.ID.String()),
			htmx.Text(profile.Name),
		)
	}

	lenses, err := w.db.ListLenses(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return err
	}

	lensesItems := make([]htmx.Node, len(lenses))
	for i, lens := range lenses {
		lensesItems[i] = htmx.Option(
			htmx.Attribute("value", lens.ID.String()),
			htmx.Text(lens.Name),
		)
	}

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				htmx.FormElement(
					htmx.HxPost(""),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"w-full": true,
								"my-4":   true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Properties"),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{
										"py-4": true,
									},
								},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"-my-4": true,
											},
										},
										htmx.Text("Name"),
									),
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("A unique identifier for the workload."),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name: "name",
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("The name must be from 3 to 100 characters. At least 3 characters must be non-whitespace."),
									),
								),
								forms.FormControl(
									forms.FormControlProps{
										ClassNames: htmx.ClassNames{
											"py-4": true,
										},
									},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"-my-4": true,
												},
											},
											htmx.Text("Description"),
										),
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
												},
											},
											htmx.Text("A brief description of the workload to document its scope and intended purpose."),
										),
									),
									forms.TextInputBordered(
										forms.TextInputProps{
											Name: "description",
										},
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
												},
											},
											htmx.Text("The description must be from 3 to 1024 characters."),
										),
									),
								),
							),
						),
					),
				),
				cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"w-full": true,
							"my-4":   true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Profile"),
						),
						dropdowns.Dropdown(
							dropdowns.DropdownProps{},
							htmx.Input(
								htmx.Attribute("type", "hidden"),
								htmx.ID("profile-input"),
								htmx.Attribute("name", "profile"),
								htmx.Value("good"),
								htmx.HyperScript("on newprofile set @value to 'tag'"),
							),
							dropdowns.DropdownButton(
								dropdowns.DropdownButtonProps{
									ClassNames: htmx.ClassNames{
										"m-1": true,
									},
								},
								htmx.Role("button"),
								htmx.Text("Select Profile"),
							),
							dropdowns.DropdownMenuItems(
								dropdowns.DropdownMenuItemsProps{},
								dropdowns.DropdownMenuItem(
									dropdowns.DropdownMenuItemProps{},
									htmx.Text("Profile One"),
									htmx.DataAttribute("profile", "1"),
									htmx.HyperScript("on click send newprofile(tag: ()) to #profile-input"),
								),
							),
						),
						htmx.Select(
							htmx.ClassNames{
								"select":   true,
								"max-w-xs": true,
								"block":    true,
							},
							htmx.Attribute("name", "profile"),
							htmx.Group(profilesItems...),
						),
					),
				),
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text("Lens"),
					),
				),
				cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"w-full": true,
							"my-4":   true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Tags (Optional)"),
						),
					),
				),
				htmx.Select(
					htmx.ClassNames{
						"select":   true,
						"max-w-xs": true,
						"block":    true,
					},
					htmx.Attribute("name", "lens"),
					htmx.Group(lensesItems...),
				),
				buttons.OutlinePrimary(
					buttons.ButtonProps{},
					htmx.Attribute("type", "submit"),
					htmx.Text("Create Workload"),
				),
			),
		),
	)
}
