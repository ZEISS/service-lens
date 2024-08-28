package designs

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// CommentsControllerImpl ...
type CommentsControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCommentsController ...
func NewCommentsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CommentsControllerImpl {
	return &CommentsControllerImpl{store: store}
}

// Error ...
func (l *CommentsControllerImpl) Error(err error) error {
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

// Delete ...
func (l *CommentsControllerImpl) Delete() error {
	var params struct {
		DesignID  uuid.UUID `json:"design_id" params:"id"`
		CommentID uuid.UUID `json:"Comment_id" params:"Comment_id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteDesignComment(ctx, &models.DesignComment{ID: params.CommentID, DesignID: params.DesignID})
	})
}
