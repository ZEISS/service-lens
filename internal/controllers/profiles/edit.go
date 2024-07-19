package profiles

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// ProfileEditController ...
type ProfileEditController struct {
	profile models.Profile
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewProfileEditController ...
func NewProfileEditController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProfileEditController {
	return &ProfileEditController{
		profile: models.Profile{},
		store:   store,
	}
}

// Prepare ...
func (p *ProfileEditController) Prepare() error {
	err := p.BindParams(p.profile)
	if err != nil {
		return err
	}

	return p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetProfile(ctx, &p.profile)
	})
}

// // Post ...
// func (p *ProfileEditController) Post() error {
// 	query := NewDefaultProfileNewControllerQuery()
// 	if err := p.BindQuery(query); err != nil {
// 		return err
// 	}

// 	p.profile.Description = query.Description

// 	err := p.db.UpdateProfile(p.Context(), p.profile)
// 	if err != nil {
// 		return err
// 	}

// 	// team := p.Values(utils.ValuesKeyTeam).(*authz.Team)

// 	// p.Hx().Redirect(fmt.Sprintf("/%s/profiles/%s", team.Slug, p.profile.ID))

// 	return nil
// }

// New ...
func (p *ProfileEditController) Get() error {
	return p.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
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
										Name:     "name",
										Value:    p.profile.Name,
										Disabled: true,
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
										Name:  "description",
										Value: p.profile.Description,
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
				buttons.OutlinePrimary(
					buttons.ButtonProps{},
					htmx.Attribute("type", "submit"),
					htmx.Text("Update Profile"),
				),
			),
		),
	)
}
