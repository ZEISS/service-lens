package designs

import (
	"bytes"
	"context"

	"github.com/google/uuid"
	markdown "github.com/teekennedy/goldmark-markdown"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/service-lens/internal/builders/renderers"
	"github.com/zeiss/service-lens/internal/builders/transformers"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
)

var _ = htmx.Controller(&SearchWorkflowsControllerImpl{})

// TaskControllerImpl ...
type TaskControllerImpl struct {
	design models.Design
	task   int
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTaskController ...
func NewTaskController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TaskControllerImpl {
	return &TaskControllerImpl{
		store: store,
	}
}

// Error ...
func (l *TaskControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *TaskControllerImpl) Prepare() error {
	var params struct {
		DesignID uuid.UUID `json:"id" params:"id" alidate:"required"`
		Task     int       `json:"task" form:"task" validate:"required"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}
	l.design.ID = params.DesignID

	err = l.BindBody(&params)
	if err != nil {
		return err
	}

	l.task = conv.Int(params.Task)

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &l.design)
	})
}

// Post ...
func (l *TaskControllerImpl) Post() error {
	taskTransformer := transformers.TaskTransformer{
		Task: l.task,
	}
	prioritizedTransformer := util.Prioritized(&taskTransformer, 0)

	md := markdown.NewRenderer()
	gm := goldmark.New(
		goldmark.WithRenderer(md),
		goldmark.WithParserOptions(parser.WithASTTransformers(prioritizedTransformer)),
		goldmark.WithExtensions(extension.GFM),
	)

	gm.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(renderers.NewTaskCheckBoxHTMLRenderer(), 2),
	))

	buf := bytes.Buffer{}
	err := gm.Convert([]byte(l.design.Body), &buf)
	if err != nil {
		return err
	}

	l.design.Body = buf.String()

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateDesign(ctx, &l.design)
	})
	if err != nil {
		return err
	}

	return l.Render(
		htmx.Fragment(
			forms.Checkbox(
				forms.CheckboxProps{
					Disabled: true,
					Checked:  true,
				},
			),
		),
	)
}
