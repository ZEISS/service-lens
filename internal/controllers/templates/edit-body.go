package templates

import (
	"bytes"
	"context"
	"fmt"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/builders"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

var _ = htmx.Controller(&EditBodyControllerImpl{})

// EditBodyControllerImpl ...
type EditBodyControllerImpl struct {
	// Template ...
	Template models.Template
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewEditBodyController ...
func NewEditBodyController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EditBodyControllerImpl {
	return &EditBodyControllerImpl{store: store}
}

// Error ...
func (l *EditBodyControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *EditBodyControllerImpl) Prepare() error {
	err := l.BindParams(&l.Template)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTemplate(ctx, &l.Template)
	})
}

// Put ...
func (l *EditBodyControllerImpl) Put() error {
	errorx.Panic(l.BindBody(&l.Template))
	errorx.Panic(l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		err := tx.UpdateTemplate(ctx, &l.Template)
		if err != nil {
			return err
		}

		markdown := goldmark.New(
			goldmark.WithRendererOptions(
				html.WithXHTML(),
				html.WithUnsafe(),
				renderer.WithNodeRenderers(
					util.Prioritized(
						builders.NewMarkdownBuilder(), 1),
				),
			),
			goldmark.WithExtensions(
				extension.GFM,
				emoji.Emoji,
			),
		)

		var b bytes.Buffer
		err = markdown.Convert([]byte(l.Template.Body), &b)
		if err != nil {
			return err
		}

		l.Template.Body = b.String()

		return err
	}))

	return l.Render(
		htmx.Fragment(
			htmx.Div(
				htmx.ID("body"),
				htmx.HxSwapOob(conv.String(true)),
				htmx.Div(
					htmx.Raw(l.Template.Body),
				),
			),
			buttons.Button(
				buttons.ButtonProps{},
				htmx.HxSwap("outerHTML"),
				htmx.HxGet(fmt.Sprintf(utils.EditTemplateBodyUrlFormat, l.Template.ID)),
				htmx.Text("Edit"),
			),
		),
	)
}

// Get ...
func (l *EditBodyControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.Div(
				htmx.ID("body"),
				htmx.HxSwapOob(conv.String(true)),
				htmx.FormElement(
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{},
						},
						designs.Editor(
							designs.EditorProps{
								Content: l.Template.Body,
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
								htmx.Text("Supports Markdown."),
							),
						),
					),
				),
			),
			buttons.Button(
				buttons.ButtonProps{
					Type: "submit",
				},
				htmx.HxSwap("outerHTML"),
				htmx.HxPut(fmt.Sprintf(utils.EditTemplateBodyUrlFormat, l.Template.ID)),
				htmx.HxInclude("body"),
				htmx.Text("Update"),
			),
		),
	)
}
