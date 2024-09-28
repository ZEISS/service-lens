package designs

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/components/designs"
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
	return toasts.Error(err.Error())
}

// Post ...
func (l *CommentsControllerImpl) Post() error {
	validate = validator.New()

	var params struct {
		DesignID uuid.UUID `json:"id" params:"id" validate:"required,uuid"`
		Comment  string    `json:"comment" validate:"required"`
	}

	err := l.BindBody(&params)
	if err != nil {
		return err
	}

	err = l.BindParams(&params)
	if err != nil {
		return err
	}

	err = validate.Struct(&params)
	if err != nil {
		return err
	}

	comment := models.DesignComment{
		DesignID: params.DesignID,
		Comment:  params.Comment,
		AuthorID: l.Session().ID,
		Author:   l.Session().User,
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateDesignComment(ctx, &comment)
	})
	if err != nil {
		return err
	}

	return l.Render(
		designs.DesignComment(
			designs.DesignCommentProps{
				Comment: comment,
				User:    l.Session().User,
				Design:  comment.Design,
			},
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
