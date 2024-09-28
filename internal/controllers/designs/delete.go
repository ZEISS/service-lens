package designs

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// DesignDeleteControllerImpl ...
type DesignDeleteControllerImpl struct {
	ID    uuid.UUID `json:"id" form:"id" param:"id" validate:"required"`
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewDesignDeleteController ...
func NewDesignDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *DesignDeleteControllerImpl {
	return &DesignDeleteControllerImpl{store: store}
}

// Error ...
func (l *DesignDeleteControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (p *DesignDeleteControllerImpl) Prepare() error {
	return p.BindParams(p)
}

// Delete ...
func (p *DesignDeleteControllerImpl) Delete() error {
	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteDesign(ctx, &models.Design{ID: p.ID})
	})
}
