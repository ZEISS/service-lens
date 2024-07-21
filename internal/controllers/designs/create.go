package designs

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

var validate *validator.Validate

type CreateDesignBody struct {
	Title string `json:"title" form:"title" validate:"required,min=3,max=2048"`
	Body  string `json:"body" form:"body" validate:"required"`
}

// CreateDesignControllerImpl ...
type CreateDesignControllerImpl struct {
	body  CreateDesignBody
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCreateDesignController ...
func NewCreateDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateDesignControllerImpl {
	return &CreateDesignControllerImpl{store: store}
}

// Prepare ...
func (l *CreateDesignControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.Ctx().BodyParser(&l.body)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.body)
	if err != nil {
		return err
	}

	return nil
}

// Error ...
func (l *CreateDesignControllerImpl) Error(err error) error {
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

// Post ...
func (l *CreateDesignControllerImpl) Post() error {
	design := models.Design{
		Title: l.body.Title,
		Body:  l.body.Body,
	}

	l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateDesign(ctx, &design)
	})

	return l.Redirect(fmt.Sprintf(utils.ShowDesigUrlFormat, design.ID))
}
