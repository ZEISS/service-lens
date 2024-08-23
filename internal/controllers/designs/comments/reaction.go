package comments

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// ReactionCommentControllerImpl ...
type ReactionCommentControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewReactionCommentController ...
func NewReactionCommentController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ReactionCommentControllerImpl {
	return &ReactionCommentControllerImpl{store: store}
}

// Error ...
func (l *ReactionCommentControllerImpl) Error(err error) error {
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
func (l *ReactionCommentControllerImpl) Delete() error {
	var params struct {
		ID         uuid.UUID `json:"id" params:"id"`
		CommentID  uuid.UUID `json:"comment_id" params:"comment_id"`
		ReactionID int       `json:"reaction_id" params:"reaction_id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	reaction := models.Reaction{
		ID:        params.ReactionID,
		ReactorID: l.Session().User.ID,
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteReaction(ctx, &reaction)
	})
}

// Prepare ...
func (l *ReactionCommentControllerImpl) Prepare() error {
	return nil
}

// Post ...
func (l *ReactionCommentControllerImpl) Post() error {
	var params struct {
		ID        uuid.UUID `json:"id" params:"id"`
		CommentID uuid.UUID `json:"comment_id" params:"comment_id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	var body struct {
		Reaction string `json:"emoji" form:"emoji" validate:"required"`
	}

	err = l.BindBody(&body)
	if err != nil {
		return err
	}

	reaction := models.Reaction{
		ReactableID:   params.CommentID,
		ReactableType: models.ReactableTypeDesignComment,
		Value:         body.Reaction,
		ReactorID:     l.Session().User.ID,
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateReaction(ctx, &reaction)
	})
}
