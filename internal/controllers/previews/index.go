package previews

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/ports"
	"go.abhg.dev/goldmark/mermaid"

	htmx "github.com/zeiss/fiber-htmx"
)

// PreviewControllerImpl ...
type PreviewControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewPreviewController ...
func NewPreviewController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *PreviewControllerImpl {
	return &PreviewControllerImpl{
		store: store,
	}
}

// Post ...
func (m *PreviewControllerImpl) Post() error {
	var form struct {
		Body string `json:"body" form:"body"`
	}

	err := m.BindBody(&form)
	if err != nil {
		return fmt.Errorf("bind body: %w", err)
	}

	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			extension.GFM,
			emoji.Emoji,
			&mermaid.Extender{},
		),
	)

	var b bytes.Buffer
	err = markdown.Convert([]byte(form.Body), &b)
	if err != nil {
		return err
	}

	return m.Render(htmx.Raw(b.String()))
}
