package designs

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
	return toasts.RenderToasts(l.Ctx(), toasts.Error(err.Error()))
}

// Delete ...
func (l *TagControllerImpl) Delete() error {
	var params struct {
		DesignID uuid.UUID `json:"design_id" params:"id"`
		TagID    uuid.UUID `json:"tag_id" params:"tag_id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.RemoveTagDesign(ctx, params.DesignID, &models.Tag{ID: params.TagID})
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
		DesignID uuid.UUID `json:"designID" params:"id"`
		Name     string    `json:"name" forms:"name"`
		Value    string    `json:"value" forms:"value"`
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
		return tx.AddTagDesign(ctx, params.DesignID, &tag)
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
						DesignID: params.DesignID,
						Tag:      tag,
					},
				),
			),
		),
	)
}
