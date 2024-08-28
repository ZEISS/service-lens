package comments

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

var validate *validator.Validate

// CreateDesignCommentControllerImpl ...
type CreateDesignCommentControllerImpl struct {
	Comment models.DesignComment
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCreateDesignCommentController ...
func NewCreateDesignCommentController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateDesignCommentControllerImpl {
	return &CreateDesignCommentControllerImpl{store: store}
}

// Error ...
func (l *CreateDesignCommentControllerImpl) Error(err error) error {
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

// Prepare ...
func (l *CreateDesignCommentControllerImpl) Prepare() error {
	validate = validator.New()

	var request struct {
		ID      uuid.UUID `uri:"id" validate:"required,uuid"`
		Comment string    `json:"comment" validate:"required"`
	}

	err := l.BindBody(&request)
	if err != nil {
		return err
	}

	err = l.BindParams(&request)
	if err != nil {
		return err
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	l.Comment = models.DesignComment{
		DesignID: request.ID,
		Comment:  request.Comment,
		AuthorID: l.Session().ID,
		Author:   l.Session().User,
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateDesignComment(ctx, &l.Comment)
	})
}

// Post ...
func (l *CreateDesignCommentControllerImpl) Post() error {
	return l.Render(
		designs.DesignComment(
			designs.DesignCommentProps{
				Comment: l.Comment,
				User:    l.Session().User,
				Design:  l.Comment.Design,
			},
		),
	)
}
