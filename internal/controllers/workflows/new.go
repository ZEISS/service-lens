package workflows

import (
	"context"
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
)

var validate *validator.Validate

// NewWorkflowControllerImpl ...
type NewWorkflowControllerImpl struct {
	Name        string `json:"name" form:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" form:"description" validate:"required,min=3,max=255"`

	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkflowController ...
func NewWorkflowController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewWorkflowControllerImpl {
	return &NewWorkflowControllerImpl{store: store}
}

// Error ...
func (t *NewWorkflowControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (t *NewWorkflowControllerImpl) Prepare() error {
	err := t.BindBody(t)
	if err != nil {
		return err
	}

	validate = validator.New()
	return validate.Struct(t)
}

// Post ...
func (t *NewWorkflowControllerImpl) Post() error {
	workflow := models.Workflow{
		Name:        t.Name,
		Description: t.Description,
	}

	err := t.store.ReadWriteTx(t.Context(), func(ctx context.Context, w ports.ReadWriteTx) error {
		return w.CreateWorkflow(ctx, &workflow)
	})
	if err != nil {
		return err
	}

	return t.Redirect(fmt.Sprintf(utils.ShowWorkflowUrlFormat, workflow.ID))
}
