package workloads

import (
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
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
											htmx.Text("Review Owner"),
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
											htmx.Text("The email address of the person responsible for reviewing the workload."),
										),
									),
									forms.TextInputBordered(
										forms.TextInputProps{
											Name: "review_owner",
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
								htmx.Text("Environment"),
							),
							htmx.Input(
								htmx.Attribute("type", "hidden"),
								htmx.ID("environment"),
								htmx.Attribute("name", "environment_id"),
								htmx.Value(""),
							),
							dropdowns.Dropdown(
								dropdowns.DropdownProps{},
								htmx.HyperScript("on click from (closest <a/>) set (previous <input/>).value to 'test'"),
								dropdowns.DropdownButton(
									dropdowns.DropdownButtonProps{},
									htmx.Text("Select Environment"),
									htmx.HxGet("/workloads/partials/environments"),
									htmx.HxTarget("#environments-list"),
									htmx.HxSwap("innerHTML"),
									htmx.ID("environments-button"),
								),
								dropdowns.DropdownMenuItems(
									dropdowns.DropdownMenuItemsProps{
										TabIndex: 1,
									},
									htmx.ID("environments-list"),
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
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
							htmx.Input(
								htmx.Attribute("type", "hidden"),
								htmx.ID("profile"),
								htmx.Attribute("name", "profile_id"),
								htmx.Value(""),
							),
							dropdowns.Dropdown(
								dropdowns.DropdownProps{},
								htmx.HyperScript("on click from (closest <a/>) set (previous <input/>).value to 'test'"),
								dropdowns.DropdownButton(
									dropdowns.DropdownButtonProps{},
									htmx.Text("Select Profile"),
									htmx.HxGet("/workloads/partials/profiles"),
									htmx.HxTarget("#profiles-list"),
									htmx.HxSwap("innerHTML"),
									htmx.ID("profiles-button"),
								),
								dropdowns.DropdownMenuItems(
									dropdowns.DropdownMenuItemsProps{
										TabIndex: 1,
									},
									htmx.ID("profiles-list"),
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
									),
								),
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
								components.TagInput(
									components.TagInputProps{},
								),
							),
						),
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
