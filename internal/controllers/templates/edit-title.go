package templates

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components/templates"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

var _ = htmx.Controller(&EditTitleControllerImpl{})

// EditTitleControllerImpl ...
type EditTitleControllerImpl struct {
	Template models.Template
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewEditTitleController ...
func NewEditTitleController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EditTitleControllerImpl {
	return &EditTitleControllerImpl{store: store}
}

// Prepare ...
func (l *EditTitleControllerImpl) Prepare() error {
	err := l.BindParams(&l.Template)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTemplate(ctx, &l.Template)
	})
}

// Put ...
func (l *EditTitleControllerImpl) Put() error {
	err := l.BindBody(&l.Template)
	if err != nil {
		return err
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateTemplate(ctx, &l.Template)
	})
	if err != nil {
		return err
	}

	return l.Render(
		templates.TemplateTitleCard(
			templates.TemplateTitleCardProps{
				Template: l.Template,
			},
		),
	)
}

// Get ...
func (l *EditTitleControllerImpl) Get() error {
	return l.Render(
		htmx.FormElement(
			htmx.HxPut(fmt.Sprintf(utils.EditTemplateTitleUrlFormat, l.Template.ID)),
			htmx.HxTarget("this"),
			htmx.HxSwap("outerHTML"),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{},
						},
						forms.TextInputBordered(
							forms.TextInputProps{
								Name:        "name",
								Placeholder: "Name",
								Value:       l.Template.Name,
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
								htmx.Text("The name must be from 3 to 2048 characters."),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						links.Link(
							links.LinkProps{
								ClassNames: htmx.ClassNames{
									"btn":       true,
									"btn-ghost": true,
								},
								Href: fmt.Sprintf(utils.ShowTemplateUrlFormat, l.Template.ID),
							},
							htmx.Text("Cancel"),
						),
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Attribute("type", "submit"),
							htmx.Text("Save"),
						),
					),
				),
			),
		),
	)
}
