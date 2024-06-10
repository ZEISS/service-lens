package workloads

import (
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	utilz "github.com/zeiss/fiber-htmx/components/utils"
)

// WorkloadNewControllerImpl ...
type WorkloadNewControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewWorkloadController ...
func NewWorkloadController(store ports.Datastore) *WorkloadNewControllerImpl {
	return &WorkloadNewControllerImpl{store: store}
}

// // Post ...
// func (w *WorkloadNewControllerImpl) Post() error {
// 	query := NewWorkloadNewControllerQuery()
// 	if err := w.BindQuery(query); err != nil {
// 		return err
// 	}

// 	validate := validator.New(validator.WithRequiredStructEnabled())
// 	err := validate.Struct(query)
// 	if err != nil {
// 		return err
// 	}

// 	team := w.Values(utils.ValuesKeyTeam).(*authz.Team)

// 	workload := &models.Workload{
// 		Description: query.Description,
// 		Lenses:      []*models.Lens{{ID: query.Lens}},
// 		Name:        query.Name,
// 		ProfileID:   query.Profile,
// 		Team:        *team,
// 	}

// 	err = w.db.CreateWorkload(w.Context(), workload)
// 	if err != nil {
// 		return err
// 	}

// 	w.Hx().Redirect(fmt.Sprintf("/%s/workloads/%s", team.Slug, workload.ID))

// 	return nil
// }

// Get ...
func (w *WorkloadNewControllerImpl) Get() error {

	// team := w.Values(utils.ValuesKeyTeam).(*authz.Team)

	// environments, err := w.db.ListEnvironment(w.Context(), team.Slug, models.Pagination[*models.Environment]{Limit: 10, Offset: 0})
	// if err != nil {
	// 	return err
	// }

	// profiles, err := w.db.ListProfiles(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: 10, Offset: 0})
	// if err != nil {
	// 	return err
	// }

	profilesItems := make([]htmx.Node, 0)
	// for i, profile := range profiles {
	// 	profilesItems[i] = htmx.Option(
	// 		htmx.Attribute("value", profile.ID.String()),
	// 		htmx.Text(profile.Name),
	// 	)
	// }

	// lenses, err := w.db.ListLenses(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: 10, Offset: 0})
	// if err != nil {
	// 	return err
	// }

	lensesItems := make([]htmx.Node, 0)
	// for i, lens := range lenses {
	// 	lensesItems[i] = htmx.Option(
	// 		htmx.Attribute("value", lens.ID.String()),
	// 		htmx.Text(lens.Name),
	// 	)
	// }

	return w.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: w.Path(),
				},
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
									forms.TextareaBordered(
										forms.TextareaProps{
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
							components.MultiSelect(
								components.MultiSelectProps{},
							),
							// dropdowns.Dropdown(
							// 	dropdowns.DropdownProps{},
							// 	htmx.Input(
							// 		htmx.Attribute("type", "hidden"),
							// 		htmx.ID("profile-input"),
							// 		htmx.Attribute("name", "profile"),
							// 		htmx.Value("good"),
							// 		htmx.HyperScript("on newprofile set @value to 'tag'"),
							// 	),
							// 	dropdowns.DropdownButton(
							// 		dropdowns.DropdownButtonProps{
							// 			ClassNames: htmx.ClassNames{
							// 				"m-1": true,
							// 			},
							// 		},
							// 		htmx.Role("button"),
							// 		htmx.Text("Select Profile"),
							// 	),
							// 	dropdowns.DropdownMenuItems(
							// 		dropdowns.DropdownMenuItemsProps{},
							// 		dropdowns.DropdownMenuItem(
							// 			dropdowns.DropdownMenuItemProps{},
							// 			htmx.Text("Profile One"),
							// 			htmx.DataAttribute("profile", "1"),
							// 			// htmx.HyperScript("on click send newprofile(tag: ()) to #profile-input"),
							// 		),
							// 	),

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
								htmx.Text("Environment"),
							),
							utilz.Map(func(el *models.Environment) htmx.Node {
								return forms.FormControl(
									forms.FormControlProps{},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"-my-4": true,
												},
											},
											htmx.Text(el.Name),
										),
										forms.Radio(
											forms.RadioProps{
												Name:  "environment",
												Value: el.ID.String(),
											},
										),
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
													"-my-4":            true,
												},
											},
											htmx.Text(el.Description),
										),
									),
								)
							}),
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
								components.TagInput(
									components.TagInputProps{},
								),
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
		),
	)
}
