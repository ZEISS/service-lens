package designs_edit_body

import (
	"bytes"
	"context"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
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
	return toasts.RenderToasts(
		l.Ctx(),
		toasts.Toasts(
			toasts.ToastsProps{},
			toasts.ToastAlertError(
				toasts.ToastProps{},
				htmx.Text(err.Error()),
			),
		),
	)
}

// Prepare ...
func (l *UpdateControllerImpl) Prepare() error {
	err := l.BindParams(&l.Design)
	if err != nil {
		return err
	}

	err = l.BindBody(&l.Design)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		err := tx.UpdateDesign(ctx, &l.Design)
		if err != nil {
			return err
		}

		markdown := goldmark.New(
			goldmark.WithRendererOptions(
				html.WithXHTML(),
				html.WithUnsafe(),
			),
			goldmark.WithExtensions(
				extension.GFM,
			),
		)

		var b bytes.Buffer
		err = markdown.Convert([]byte(l.Design.Body), &b)
		if err != nil {
			return err
		}

		l.Design.Body = b.String()

		return err
	})
}

// Prepare ...
func (l *UpdateControllerImpl) Put() error {
	return l.Render(
		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
			},
			htmx.HxTarget("this"),
			htmx.HxSwap("outerHTML"),
			htmx.ID("body"),
			cards.Body(
				cards.BodyProps{},
				htmx.Div(
					htmx.Raw(l.Design.Body),
				),
				cards.Actions(
					cards.ActionsProps{},
					buttons.Outline(
						buttons.ButtonProps{},
						htmx.HxGet(fmt.Sprintf(utils.EditBodyUrlFormat, l.Design.ID)),
						htmx.Text("Edit"),
					),
				),
			),
		),
	)
}
