package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

var _ ports.Lenses = (*Lenses)(nil)

// Lenses ...
type Lenses struct {
	lenses ports.Lenses
}

// NewLensesController ...
func NewLensesController(lenses ports.Lenses) *Lenses {
	return &Lenses{lenses}
}

// AddLens ...
func (l *Lenses) AddLens(ctx context.Context, lens *models.Lens) (*models.Lens, error) {
	return l.lenses.AddLens(ctx, lens)
}

// GetLensByID ...
func (l *Lenses) GetLensByID(ctx context.Context, id uuid.UUID) (*models.Lens, error) {
	return l.lenses.GetLensByID(ctx, id)
}
