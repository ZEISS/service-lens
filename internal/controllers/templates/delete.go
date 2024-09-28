package templates

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// DeleteTemplateControllerImpl ...
type DeleteTemplateControllerImpl struct {
	ID    uuid.UUID `param:"id"`
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewDeleteTemplateController ...
func NewDeleteTemplateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *DeleteTemplateControllerImpl {
	return &DeleteTemplateControllerImpl{store: store}
}

// Prepare ...
func (l *DeleteTemplateControllerImpl) Prepare() error {
	return l.BindParams(l)
}

// Error ...
func (l *DeleteTemplateControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Post ...
func (l *DeleteTemplateControllerImpl) Delete() error {
	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTemplate(ctx, &models.Template{ID: l.ID})
	})
}
