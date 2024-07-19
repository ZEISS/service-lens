package partials

import (
	"context"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// EnvironmentPartialListControllerImpl ...
type EnvironmentParialListControllerImpl struct {
	environments tables.Results[models.Environment]
	store        seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.UnimplementedController
}

// NewEnvironmentPartialListController ...
func NewEnvironmentPartialListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EnvironmentParialListControllerImpl {
	return &EnvironmentParialListControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *EnvironmentParialListControllerImpl) Prepare() error {
	err := w.BindQuery(&w.environments)
	if err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListEnvironments(ctx, &w.environments)
	})
}

// Get ...
func (w *EnvironmentParialListControllerImpl) Get() error {
	return w.Render(
		htmx.Fragment(
			htmx.ForEach(w.environments.GetRows(), func(e *models.Environment, choiceIdx int) htmx.Node {
				return dropdowns.DropdownMenuItem(
					dropdowns.DropdownMenuItemProps{},
					htmx.A(
						htmx.Text(e.Name),
						htmx.DataAttribute("environment", e.ID.String()),
						htmx.HyperScript(`on click set (previous <input/>).value to my @data-environment then put my innerHTML into #environments-button`),
					),
				)
			})...,
		),
	)
}
