package workloads

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

var validate *validator.Validate

// CreateWorkloadControllerImpl ...
type CreateWorkloadControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCreateWorkloadController ...
func NewCreateWorkloadController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateWorkloadControllerImpl {
	return &CreateWorkloadControllerImpl{store: store}
}

// Error ...
func (l *CreateWorkloadControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Post ...
func (l *CreateWorkloadControllerImpl) Post() error {
	var body struct {
		Name          string    `json:"name" form:"name" validate:"required"`
		Description   string    `json:"description" form:"description" validate:"required"`
		EnvironmentID uuid.UUID `json:"environment_id" form:"environment_id" validate:"required"`
		LensID        uuid.UUID `json:"lens_id" form:"lens_id" validate:"required"`
		ProfileID     uuid.UUID `json:"profile_id" form:"profile_id" validate:"required"`
		ReviewOwner   string    `json:"review_owner" form:"review_owner" validate:"required"`
	}

	err := l.BindBody(&body)
	if err != nil {
		return err
	}

	workload := models.Workload{
		Name:          body.Name,
		Description:   body.Description,
		EnvironmentID: body.EnvironmentID,
		Lenses:        []*models.Lens{{ID: body.LensID}},
		ProfileID:     body.ProfileID,
		ReviewOwner:   body.ReviewOwner,
	}

	validate = validator.New()

	err = validate.Struct(&workload)
	if err != nil {
		return err
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateWorkload(ctx, &workload)
	})
	if err != nil {
		return err
	}

	return l.Redirect(fmt.Sprintf(utils.ShowWorkloadUrlFormat, workload.ID))
}
