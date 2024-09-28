package design_edit_title

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

var _ = htmx.Controller(&UpdateControllerImpl{})

// UpdateControllerImpl ...
type UpdateControllerImpl struct {
	Design models.Design
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewUpdateController ...
func NewUpdateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *UpdateControllerImpl {
	return &UpdateControllerImpl{store: store}
}

// Error ...
func (l *UpdateControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *UpdateControllerImpl) Prepare() error {
	err := l.BindParams(&l.Design)
	if err != nil {
		return err
	}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &l.Design)
	})
	if err != nil {
		return err
	}

	err = l.BindBody(&l.Design)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateDesign(ctx, &l.Design)
	})
}

// Prepare ...
func (l *UpdateControllerImpl) Put() error {
	return l.Render(
		htmx.FormElement(
			htmx.HxGet(fmt.Sprintf(utils.EditTitleUrlFormat, l.Design.ID)),
			htmx.HxTarget("this"),
			htmx.HxSwap("outerHTML"),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"my-2": true,
						"mx-2": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					htmx.ID("body"),
					htmx.H1(htmx.Text(l.Design.Title)),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Outline(
							buttons.ButtonProps{},
							htmx.HxGet(fmt.Sprintf(utils.EditTitleUrlFormat, l.Design.ID)),
							htmx.Text("Edit"),
						),
					),
				),
			),
		),
	)
}
