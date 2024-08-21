package workloads

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/loading"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/pkg/slices"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadNewControllerImpl ...
type WorkloadNewControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadController ...
func NewWorkloadController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadNewControllerImpl {
	return &WorkloadNewControllerImpl{store: store}
}

// Get ...
func (w *WorkloadNewControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "New Workload",
				Path:  w.Path(),
				User:  w.Session().User,
			},
			func() htmx.Node {
				lenses := tables.Results[models.Lens]{}
				profiles := tables.Results[models.Profile]{}
				environments := tables.Results[models.Environment]{}

				errorx.Panic(w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					if err := tx.ListLenses(ctx, &lenses); err != nil {
						return err
					}

					if err := tx.ListProfiles(ctx, &profiles); err != nil {
						return err
					}

					return tx.ListEnvironments(ctx, &environments)
				}))

				return htmx.FormElement(
					htmx.HxPost(""),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Properties"),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Name"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:        "name",
										Placeholder: "Shop System, Payment Gateway, etc.",
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
									forms.FormControlProps{},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{},
											htmx.Text("Description"),
										),
									),
									forms.TextareaBordered(
										forms.TextareaProps{
											Name:        "description",
											Placeholder: "This is a shop system that processes payments.",
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
									forms.FormControlProps{},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{},
											htmx.Text("Review Owner"),
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
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Environment"),
									),
								),
								dropdowns.Dropdown(
									dropdowns.DropdownProps{},
									alpine.XData(`{
                      selected: {},
                      onOptionClick(id, name) {
                           this.selected = { id, name };
                        },
                    }`),
									htmx.Div(
										htmx.ClassNames{
											tailwind.Flex:          true,
											tailwind.SpaceX4:       true,
											tailwind.JustifyCenter: true,
										},
										forms.TextInputBordered(
											forms.TextInputProps{
												Placeholder: "Search an environment ...",
												Name:        "search",
											},
											alpine.XModel("selected.name"),
											alpine.XRef("button"),
											htmx.HxPost(utils.WorkloadSearchEnvironmentsUrlFormat),
											htmx.HxTarget("#search-results"),
											htmx.HxTrigger("keyup changed delay:500ms"),
											htmx.HxIndicator(".htmx-indicator"),
										),
										loading.Spinner(
											loading.SpinnerProps{
												ClassNames: htmx.ClassNames{
													"htmx-indicator": true,
												},
											},
										),
									),
									htmx.Input(
										htmx.Name("environment_id"),
										htmx.Type("hidden"),
										alpine.XModel("selected.id"),
									),
									dropdowns.DropdownMenuItems(
										dropdowns.DropdownMenuItemsProps{
											ClassNames: htmx.ClassNames{
												tailwind.WFull: true,
											},
										},
										htmx.ID("search-results"),
										htmx.IfElse(
											!slices.Size(0, environments.GetRows()),
											htmx.Group(
												htmx.ForEach(environments.GetRows(), func(e *models.Environment, idx int) htmx.Node {
													return dropdowns.DropdownMenuItem(
														dropdowns.DropdownMenuItemProps{},
														htmx.A(
															htmx.Text(e.Name),
															htmx.Value(e.ID.String()),
															alpine.XOn("click", "onOptionClick($event.target.getAttribute('value'), $event.target.innerText)"),
														),
													)
												})...,
											),
											dropdowns.DropdownMenuItem(
												dropdowns.DropdownMenuItemProps{},
												htmx.A(
													htmx.Text("No environment found"),
												),
											),
										),
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
										htmx.Text("The environment that will be used to monitor the workload."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Profle"),
									),
								),
								dropdowns.Dropdown(
									dropdowns.DropdownProps{},
									alpine.XData(`{
                      selected: {},
                      onOptionClick(id, name) {
                           this.selected = { id, name };
                        },
                    }`),
									htmx.Div(
										htmx.ClassNames{
											tailwind.Flex:          true,
											tailwind.SpaceX4:       true,
											tailwind.JustifyCenter: true,
										},
										forms.TextInputBordered(
											forms.TextInputProps{
												Placeholder: "Search a profile ...",
												Name:        "search",
											},
											alpine.XModel("selected.name"),
											alpine.XRef("button"),
											htmx.HxPost(utils.WorkloadSearchProfilesUrlFormat),
											htmx.HxTarget("#search-results"),
											htmx.HxTrigger("keyup changed delay:500ms"),
											htmx.HxIndicator(".htmx-indicator"),
										),
										loading.Spinner(
											loading.SpinnerProps{
												ClassNames: htmx.ClassNames{
													"htmx-indicator": true,
												},
											},
										),
									),
									htmx.Input(
										htmx.Name("profile_id"),
										htmx.Type("hidden"),
										alpine.XModel("selected.id"),
									),
									dropdowns.DropdownMenuItems(
										dropdowns.DropdownMenuItemsProps{
											ClassNames: htmx.ClassNames{
												tailwind.WFull: true,
											},
										},
										htmx.ID("search-results"),
										htmx.IfElse(
											!slices.Size(0, profiles.GetRows()),
											htmx.Group(
												htmx.ForEach(profiles.GetRows(), func(e *models.Profile, idx int) htmx.Node {
													return dropdowns.DropdownMenuItem(
														dropdowns.DropdownMenuItemProps{},
														htmx.A(
															htmx.Text(e.Name),
															htmx.Value(e.ID.String()),
															alpine.XOn("click", "onOptionClick($event.target.getAttribute('value'), $event.target.innerText)"),
														),
													)
												})...,
											),
											dropdowns.DropdownMenuItem(
												dropdowns.DropdownMenuItemProps{},
												htmx.A(
													htmx.Text("No profile found"),
												),
											),
										),
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
										htmx.Text("The profile that will be used to monitor the workload."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Lens"),
									),
								),
								dropdowns.Dropdown(
									dropdowns.DropdownProps{},
									alpine.XData(`{
                      selected: {},
                      onOptionClick(id, name) {
                           this.selected = { id, name };
                        },
                    }`),
									htmx.Div(
										htmx.ClassNames{
											tailwind.Flex:          true,
											tailwind.SpaceX4:       true,
											tailwind.JustifyCenter: true,
										},
										forms.TextInputBordered(
											forms.TextInputProps{
												Placeholder: "Search a lens ...",
												Name:        "search",
											},
											alpine.XModel("selected.name"),
											alpine.XRef("button"),
											htmx.HxPost(utils.WorkloadSearchLensesUrlFormat),
											htmx.HxTarget("#search-results"),
											htmx.HxTrigger("keyup changed delay:500ms"),
											htmx.HxIndicator(".htmx-indicator"),
										),
										loading.Spinner(
											loading.SpinnerProps{
												ClassNames: htmx.ClassNames{
													"htmx-indicator": true,
												},
											},
										),
									),
									htmx.Input(
										htmx.Name("lens_id"),
										htmx.Type("hidden"),
										alpine.XModel("selected.id"),
									),
									dropdowns.DropdownMenuItems(
										dropdowns.DropdownMenuItemsProps{
											ClassNames: htmx.ClassNames{
												tailwind.WFull: true,
											},
										},
										htmx.ID("search-results"),
										htmx.IfElse(
											!slices.Size(0, lenses.GetRows()),
											htmx.Group(
												htmx.ForEach(lenses.GetRows(), func(e *models.Lens, idx int) htmx.Node {
													return dropdowns.DropdownMenuItem(
														dropdowns.DropdownMenuItemProps{},
														htmx.A(
															htmx.Text(e.Name),
															htmx.Value(e.ID.String()),
															alpine.XOn("click", "onOptionClick($event.target.getAttribute('value'), $event.target.innerText)"),
														),
													)
												})...,
											),
											dropdowns.DropdownMenuItem(
												dropdowns.DropdownMenuItemProps{},
												htmx.A(
													htmx.Text("No lenses found"),
												),
											),
										),
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
										htmx.Text("This is the lens that will be used to monitor the workload."),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{
										Type: "submit",
									},
									htmx.Text("Create Workload"),
								),
							),
						),
					),
					components.AddTags(components.AddTagsProps{}),
				)
			},
		),
	)
}
