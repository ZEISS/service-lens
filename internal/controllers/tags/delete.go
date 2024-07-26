package tags

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// TagDeleteControllerImpl ...
type TagDeleteControllerImpl struct {
	ID    uuid.UUID `param:"id"`
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTagDeleteController ...
func NewTagDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TagDeleteControllerImpl {
	return &TagDeleteControllerImpl{store: store}
}

// Prepare ...
func (p *TagDeleteControllerImpl) Prepare() error {
	return p.BindParams(p)
}

// Delete ...
func (p *TagDeleteControllerImpl) Delete() error {
	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTag(ctx, &models.Tag{ID: p.ID})
	})
}
