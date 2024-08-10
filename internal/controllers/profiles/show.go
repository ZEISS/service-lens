package profiles

import (
	"context"
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// ProfileShowControllerImpl ...
type ProfileShowControllerImpl struct {
	questions tables.Results[models.ProfileQuestion]
	profile   models.Profile
	store     seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewProfileShowController ...
func NewProfileShowController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProfileShowControllerImpl {
	return &ProfileShowControllerImpl{
		store: store,
	}
}

// Prepare ...
func (p *ProfileShowControllerImpl) Prepare() error {
	err := p.BindParams(&p.profile)
	if err != nil {
		return err
	}

	err = p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfileQuestions(ctx, &p.questions)
	})
	if err != nil {
		return err
	}

	return p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetProfile(ctx, &p.profile)
	})
}

// Get ...
func (p *ProfileShowControllerImpl) Get() error {
	return p.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: p.Path(),
				User: p.Session().User,
			},
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"m-2": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text("Overview"),
					),
					htmx.Div(
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("ID"),
							),
							htmx.H3(
								htmx.Text(p.profile.ID.String()),
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("Name"),
							),
							htmx.H3(
								htmx.Text(p.profile.Name),
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("Description"),
							),
							htmx.H3(
								htmx.Text(p.profile.Description),
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("Created at"),
							),
							htmx.H3(
								htmx.Text(
									p.profile.CreatedAt.Format("2006-01-02 15:04:05"),
								),
							),
						),
						htmx.Div(
							htmx.ClassNames{
								"flex":     true,
								"flex-col": true,
								"py-2":     true,
							},
							htmx.H4(
								htmx.ClassNames{
									"text-gray-500": true,
								},
								htmx.Text("Updated at"),
							),
							htmx.H3(
								htmx.Text(
									p.profile.UpdatedAt.Format("2006-01-02 15:04:05"),
								),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						links.Button(
							links.LinkProps{
								ClassNames: htmx.ClassNames{
									"btn-outline": true,
									"btn-primary": true,
								},
								Href: fmt.Sprintf("%s/edit", p.profile.ID),
							},
							htmx.Text("Edit"),
						),
						buttons.OutlinePrimary(
							buttons.ButtonProps{},
							htmx.HxDelete(""),
							htmx.HxConfirm("Are you sure you want to delete this profile?"),
							htmx.Text("Delete"),
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
														Name:     fmt.Sprintf("answers.%d.ChoiceID", profileIdx),
														Value:    utils.IntStr(c.ID),
														Checked:  p.profile.IsChoosen(c.ID),
														Disabled: true,
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
		),
	)
}
