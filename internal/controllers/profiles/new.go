package profiles

import (
	"context"
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// NewProfileControllerImpl ...
type NewProfileControllerImpl struct {
	questions tables.Results[models.ProfileQuestion]
	store     ports.Datastore
	htmx.DefaultController
}

// NewProfileController ...
func NewProfileController(store ports.Datastore) *NewProfileControllerImpl {
	return &NewProfileControllerImpl{store: store}
}

// Prepare ...
func (p *NewProfileControllerImpl) Prepare() error {
	return p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfileQuestions(ctx, &p.questions)
	})
}

// New ...
func (p *NewProfileControllerImpl) Get() error {
	return p.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: p.Path(),
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
					htmx.Group(
						htmx.ForEach(p.questions.GetRows(), func(q *models.ProfileQuestion, profileIdx int) htmx.Node {
							return cards.CardBordered(
								cards.CardProps{
									ClassNames: htmx.ClassNames{
										"w-full": true,
										"my-4":   true,
									},
								},
								cards.Body(
									cards.BodyProps{},
									htmx.Group(

										cards.Title(
											cards.TitleProps{},
											htmx.Text(q.Title),
										),
										htmx.Group(
											htmx.ForEach(q.GetChoices(), func(c *models.ProfileQuestionChoice, choiceIdx int) htmx.Node {
												return forms.FormControl(
													forms.FormControlProps{},
													forms.FormControlLabel(
														forms.FormControlLabelProps{},
														forms.FormControlLabelText(
															forms.FormControlLabelTextProps{},
															htmx.Text(c.Title),
														),
														forms.Radio(
															forms.RadioProps{
																Name:    fmt.Sprintf("answers.%d.ChoiceID", profileIdx),
																Value:   utils.IntStr(c.ID),
																Checked: choiceIdx == 0, // todo(katallaxie): should be a default option in the model
															},
														),
													),
												)
											})...),
									),
								),
							)
						})...,
					),
					// 		forms.FormControl(
					// 			forms.FormControlProps{},
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Envision Adoption Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:    "",
					// 						Value:   "1",
					// 						Checked: true,
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Align Adoption Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.0.ChoiceID",
					// 						Value: "2",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Launch Adoption Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.0.ChoiceID",
					// 						Value: "3",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Scale Adoption Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.0.ChoiceID",
					// 						Value: "4",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Post-Adoption Optimization Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.0.ChoiceID",
					// 						Value: "5",
					// 					},
					// 				),
					// 			),
					// 		),
					// 	),
					// ),
					// cards.CardBordered(
					// 	cards.CardProps{
					// 		ClassNames: htmx.ClassNames{
					// 			"w-full": true,
					// 			"my-4":   true,
					// 		},
					// 	},
					// 	cards.Body(
					// 		cards.BodyProps{},
					// 		cards.Title(
					// 			cards.TitleProps{},
					// 			htmx.Text("What is the business value that workloads in this profile represent for your team, organization, or company?"),
					// 		),
					// 		forms.FormControl(
					// 			forms.FormControlProps{},
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Business-Critical Workloads"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:    "questions.1.ChoiceID",
					// 						Value:   "1",
					// 						Checked: true,
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Strategic External-facing Workloads"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.1.ChoiceID",
					// 						Value: "2",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Strategic Internal-facing Workloads"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.1.ChoiceID",
					// 						Value: "3",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Internal Business Workloads"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.1.ChoiceID",
					// 						Value: "4",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("General Use Workloads"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.1.ChoiceID",
					// 						Value: "5",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Experimentation or Testing Workloads"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.1.ChoiceID",
					// 						Value: "6",
					// 					},
					// 				),
					// 			),
					// 		),
					// 	),
					// ),
					// cards.CardBordered(
					// 	cards.CardProps{
					// 		ClassNames: htmx.ClassNames{
					// 			"w-full": true,
					// 			"my-4":   true,
					// 		},
					// 	},
					// 	cards.Body(
					// 		cards.BodyProps{},
					// 		cards.Title(
					// 			cards.TitleProps{},
					// 			htmx.Text("What is the current architectural and operational lifecycle phase of the workloads in this profile?"),
					// 		),
					// 		forms.FormControl(
					// 			forms.FormControlProps{},
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Plan / Requirements Gathering / Design Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:    "questions.2.ChoiceID",
					// 						Value:   "1",
					// 						Checked: true,
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Development / Build / Implementation Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.2.ChoiceID",
					// 						Value: "2",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Testing / Pre-Production Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.2.ChoiceID",
					// 						Value: "3",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Deploy / Production Launch Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.2.ChoiceID",
					// 						Value: "4",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Maintenance / Optimization Phase"),
					// 				),
					// 				forms.Radio(
					// 					forms.RadioProps{
					// 						Name:  "questions.2.ChoiceID",
					// 						Value: "5",
					// 					},
					// 				),
					// 			),
					// 		),
					// 	),
					// ),
					// cards.CardBordered(
					// 	cards.CardProps{
					// 		ClassNames: htmx.ClassNames{
					// 			"w-full": true,
					// 			"my-4":   true,
					// 		},
					// 	},
					// 	cards.Body(
					// 		cards.BodyProps{},
					// 		cards.Title(
					// 			cards.TitleProps{},
					// 			htmx.Text("What are the improvement priorities for Well-Architected Framework Reviews (WAFRs) in this profile?"),
					// 		),
					// 		forms.FormControl(
					// 			forms.FormControlProps{},
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Evaluate organizational cloud strategy and priorities"),
					// 				),
					// 				forms.Checkbox(
					// 					forms.CheckboxProps{
					// 						Name:    "questions.3.ChoiceID",
					// 						Value:   "1",
					// 						Checked: false,
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Improve operational readiness"),
					// 				),
					// 				forms.Checkbox(
					// 					forms.CheckboxProps{
					// 						Name:    "questions.3.ChoiceID",
					// 						Value:   "2",
					// 						Checked: true,
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Improve operational efficiency"),
					// 				),
					// 				forms.Checkbox(
					// 					forms.CheckboxProps{
					// 						Name:    "questions.3.ChoiceID",
					// 						Value:   "3",
					// 						Checked: true,
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Improve operational incident response"),
					// 				),
					// 				forms.Checkbox(
					// 					forms.CheckboxProps{
					// 						Name:  "questions.3.ChoiceID",
					// 						Value: "4",
					// 					},
					// 				),
					// 			),
					// 			forms.FormControlLabel(
					// 				forms.FormControlLabelProps{},
					// 				forms.FormControlLabelText(
					// 					forms.FormControlLabelTextProps{
					// 						ClassNames: htmx.ClassNames{
					// 							"-my-4": true,
					// 						},
					// 					},
					// 					htmx.Text("Improve monitoring and observability"),
					// 				),
					// 				forms.Checkbox(
					// 					forms.CheckboxProps{
					// 						Name:  "questions.3.ChoiceID",
					// 						Value: "5",
					// 					},
					// 				),
					// 			),
					// 		),
					// 	),
					// ),
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
								htmx.Text("Tags - Optional"),
							),
						),
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"my-4": true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							buttons.OutlinePrimary(
								buttons.ButtonProps{},
								htmx.Attribute("type", "submit"),
								htmx.Text("Create Profile"),
							),
						),
					),
				),
			),
		),
	)
}
