package templates

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

// CreateTemplateBody ...
type CreateTemplateBody struct {
	Name        string `json:"name" form:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" form:"description" validate:"required,min=3,max=2048"`
	Body        string `json:"body" form:"body" validate:"required"`
}

// CreateTemplateControllerImpl ...
type CreateTemplateControllerImpl struct {
	body  CreateTemplateBody
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCreateTemplateController ...
func NewCreateTemplateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateTemplateControllerImpl {
	return &CreateTemplateControllerImpl{store: store}
}

// Prepare ...
func (l *CreateTemplateControllerImpl) Prepare() error {
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
func (l *CreateTemplateControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Post ...
func (l *CreateTemplateControllerImpl) Post() error {
	template := models.Template{
		Name:        l.body.Name,
		Description: l.body.Description,
		Body:        l.body.Body,
	}

	err := l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateTemplate(ctx, &template)
	})
	if err != nil {
		return err
	}

	return l.Redirect(fmt.Sprintf(utils.ShowTemplateUrlFormat, template.ID))
}
