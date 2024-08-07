package tags

import (
	"context"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

var validate *validator.Validate

// NewTagControllerImpl ...
type NewTagControllerImpl struct {
	Name  string `json:"name" form:"name" validate:"required,min=3,max=255"`
	Value string `json:"value" form:"value" validate:"required,min=3,max=255"`
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTagController ...
func NewTagController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewTagControllerImpl {
	return &NewTagControllerImpl{store: store}
}

// Error ...
func (t *NewTagControllerImpl) Error(err error) error {
	return toasts.RenderToasts(
		t.Ctx(),
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
func (t *NewTagControllerImpl) Prepare() error {
	err := t.BindBody(t)
	if err != nil {
		return err
	}

	validate = validator.New()
	return validate.Struct(t)
}

// Post ...
func (t *NewTagControllerImpl) Post() error {
	tag := models.Tag{
		Name:  t.Name,
		Value: t.Value,
	}

	return t.store.ReadWriteTx(t.Context(), func(ctx context.Context, w ports.ReadWriteTx) error {
		return w.CreateTag(ctx, &tag)
	})
}
