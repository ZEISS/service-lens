package workloads

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// TagControllerImpl ...
type TagControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTagController ...
func NewTagController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TagControllerImpl {
	return &TagControllerImpl{store: store}
}

// Error ...
func (l *TagControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Delete ...
func (l *TagControllerImpl) Delete() error {
	var params struct {
		WorkloadID uuid.UUID `json:"workload_id" params:"id"`
		TagID      uuid.UUID `json:"tag_id" params:"tag_id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.RemoveTagWorkload(ctx, params.WorkloadID, &models.Tag{ID: params.TagID})
	})
	if err != nil {
		return err
	}

	return l.Render(
		htmx.Fragment(
			htmx.Div(
				htmx.ID("tag-"+params.TagID.String()),
				htmx.HxSwapOob("outerHTML"),
			),
		),
	)
}

// Post ...
func (l *TagControllerImpl) Post() error {
	var params struct {
		WorkloadID uuid.UUID `json:"workload_id" params:"id"`
		Name       string    `json:"name" forms:"name"`
		Value      string    `json:"value" forms:"value"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	err = l.BindBody(&params)
	if err != nil {
		return err
	}

	tag := models.Tag{
		Name:  params.Name,
		Value: params.Value,
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.AddTagWorkload(ctx, params.WorkloadID, &tag)
	})
	if err != nil {
		return err
	}

	return l.Render(
		htmx.Fragment(
			htmx.Div(
				htmx.ID("tags"),
				htmx.HxSwapOob("beforeend"),
				designs.DesignTag(
					designs.DesignTagProps{
						DesignID: params.WorkloadID,
						Tag:      tag,
					},
				),
			),
		),
	)
}
