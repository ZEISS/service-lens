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
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
				Head: []htmx.Node{
					htmx.Script(
						htmx.Src("https://unpkg.com/@github/markdown-toolbar-element@latest/dist/index.js"),
						htmx.Type("module"),
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
