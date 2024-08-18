package search

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/pkg/slices"
)

var _ = htmx.Controller(&SearchEnvironmentsControllerImpl{})

// Search ...
type SearchEnvironmentsControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewSearchEnvironmentsController ...
func NewSearchEnvironmentsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *SearchEnvironmentsControllerImpl {
	return &SearchEnvironmentsControllerImpl{store: store}
}

// Error ...
func (l *SearchEnvironmentsControllerImpl) Error(err error) error {
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
func (l *SearchEnvironmentsControllerImpl) Post() error {
	return l.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(
				func() htmx.Node {
					results := tables.Results[models.Environment]{}
					errorx.Panic(l.BindQuery(&results))

					errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
						return tx.ListEnvironments(ctx, &results)
					}))

					return htmx.IfElse(
						!slices.Size(0, results.Rows...),
						htmx.Fragment(
							htmx.ForEach(results.GetRows(), func(e *models.Environment, idx int) htmx.Node {
								return dropdowns.DropdownMenuItem(
									dropdowns.DropdownMenuItemProps{},
									htmx.A(
										htmx.Text(e.Name),
										htmx.Value(e.ID.String()),
										alpine.XOn("click", "onOptionClick($event.target.getAttribute('value'), $event.target.innerText)"),
									),
								)
							})...,
						),
						dropdowns.DropdownMenuItem(
							dropdowns.DropdownMenuItemProps{},
							htmx.A(
								htmx.Text("No profile found"),
							),
						),
					)
				},
			),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		),
	)
}
