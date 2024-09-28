package lenses

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/components/lenses"
	"github.com/zeiss/service-lens/internal/ports"
)

// LensPublishControllerImpl ...
type LensPublishControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensPublishController ...
func NewLensPublishController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *LensPublishControllerImpl {
	return &LensPublishControllerImpl{
		store: store,
	}
}

// Error ...
func (c *LensPublishControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Delete ...
func (c *LensPublishControllerImpl) Delete() error {
	return c.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var params struct {
					LensID uuid.UUID `json:"id" params:"id"`
				}

				errorx.Panic(c.BindParams(&params))
				errorx.Panic(c.store.ReadWriteTx(c.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
					return tx.UnpublishLens(ctx, params.LensID)
				}))

				return htmx.Fragment(
					lenses.LensesPublishButton(
						lenses.LensesPublishButtonProps{
							ID:      params.LensID,
							IsDraft: true,
						},
					),
					lenses.LensesStatus(
						lenses.LensesStatusProps{
							IsDraft: true,
						},
						htmx.HxSwapOob("outerHTML"),
					),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		),
	)
}

// Post ...
func (c *LensPublishControllerImpl) Post() error {
	return c.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var params struct {
					LensID uuid.UUID `json:"id" params:"id"`
				}

				errorx.Panic(c.BindParams(&params))
				errorx.Panic(c.store.ReadWriteTx(c.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
					return tx.PublishLens(ctx, params.LensID)
				}))

				return htmx.Fragment(
					lenses.LensesPublishButton(
						lenses.LensesPublishButtonProps{
							ID:      params.LensID,
							IsDraft: false,
						},
					),
					lenses.LensesStatus(
						lenses.LensesStatusProps{
							IsDraft: false,
						},
						htmx.HxSwapOob("outerHTML"),
					),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		),
	)
}
