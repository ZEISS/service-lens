package designs

import (
	"context"
	"errors"

	"github.com/google/uuid"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"gorm.io/gorm"

	htmx "github.com/zeiss/fiber-htmx"
)

// NewDesignControllerImpl ...
type NewDesignControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewDesignController ...
func NewDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewDesignControllerImpl {
	return &NewDesignControllerImpl{store: store}
}

// Get ...
func (l *NewDesignControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: l.Path(),
				User: l.Session().User,
				Head: []htmx.Node{
					htmx.Link(
						htmx.Attribute("href", "https://cdn.jsdelivr.net/simplemde/1.11/simplemde.min.css"),
						htmx.Rel("stylesheet"),
						htmx.Type("text/css"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://cdn.jsdelivr.net/simplemde/1.11/simplemde.min.js"),
						htmx.Type("text/javascript"),
					),
				},
			},
			func() htmx.Node {
				params := struct {
					Template uuid.UUID `json:"template"`
				}{}
				errorx.Ignore(params, l.BindQuery(&params))

				template := models.Template{
					ID: params.Template,
				}

				err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTemplate(ctx, &template)
				})
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					panic(err)
				}

				return designs.DesignNewForm(
					designs.DesignNewFormProps{
						Template: template.Body,
					},
				)
			},
		),
	)
}
