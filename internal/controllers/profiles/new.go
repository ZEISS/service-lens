package profiles

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
)

// NewProfileControllerImpl ...
type NewProfileControllerImpl struct {
	questions tables.Results[models.ProfileQuestion]
	store     seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewProfileController ...
func NewProfileController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewProfileControllerImpl {
	return &NewProfileControllerImpl{store: store}
}

// Prepare ...
func (p *NewProfileControllerImpl) Prepare() error {
	return p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfileQuestions(ctx, &p.questions)
	})
}

// Post ...
func (p *NewProfileControllerImpl) Post() error {
	var profile models.Profile
	validate = validator.New()

	err := p.BindBody(&profile)
	if err != nil {
		return err
	}

	err = validate.Struct(&profile)
	if err != nil {
		return err
	}

	err = p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateProfile(ctx, &profile)
	})
	if err != nil {
		return toasts.Error(err.Error())
	}

	return p.Redirect(fmt.Sprintf(utils.ShowProfileUrlFormat, profile.ID))
}

// New ...
func (p *NewProfileControllerImpl) Get() error {
	return p.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       "New Profile",
				Path:        p.Path(),
				User:        p.Session().User,
				Development: p.IsDevelopment(),
			},
			func() htmx.Node {
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
								htmx.Text("Add Profile"),
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
										Placeholder: "Experiment, Production, Migration ...",
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
											Placeholder: "This is a profile for ...",
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
									cards.Actions(
										cards.ActionsProps{},
										buttons.Button(
											buttons.ButtonProps{},
											htmx.Attribute("type", "submit"),
											htmx.Text("Add"),
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
																Value:   conv.String(c.ID),
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
					components.AddTags(
						components.AddTagsProps{},
					),
				)
			},
		),
	)
}
