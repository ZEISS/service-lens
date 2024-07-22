package designs_edit_body

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

var _ = htmx.Controller(&ShowControllerImpl{})

// ShowControllerImpl ...
type ShowControllerImpl struct {
	Design models.Design
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewEditController ...
func NewEditController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowControllerImpl {
	return &ShowControllerImpl{store: store}
}

// Prepare ...
func (l *ShowControllerImpl) Prepare() error {
	err := l.BindParams(&l.Design)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &l.Design)
	})
}

// Prepare ...
func (l *ShowControllerImpl) Get() error {
	return l.Render(
		cards.CardBordered(
			cards.CardProps{},
			cards.Body(
				cards.BodyProps{},
				htmx.HxTarget("this"),
				htmx.HxSwap("outerHTML"),
				htmx.ID("body"),
				htmx.FormElement(
					htmx.HxPut(fmt.Sprintf(utils.EditBodyUrlFormat, l.Design.ID)),
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{},
						},
						forms.TextareaBordered(
							forms.TextareaProps{
								ClassNames: htmx.ClassNames{
									"h-64": true,
								},
								Name:        "body",
								Placeholder: "Start typing...",
							},
							htmx.Text(l.Design.Body),
						),
						forms.FormControlLabel(
							forms.FormControlLabelProps{},
							forms.FormControlLabelText(
								forms.FormControlLabelTextProps{
									ClassNames: htmx.ClassNames{
										"text-neutral-500": true,
									},
								},
								htmx.Text("Supports Markdown."),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Outline(
							buttons.ButtonProps{},
							htmx.Attribute("type", "submit"),
							htmx.Text("Update"),
						),
					),
				),
			),
		),
	)
}
