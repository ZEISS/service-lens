package controllers

import (
	"context"

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
func (l *Lenses) AddLens(ctx context.Context) error {
	return l.lenses.AddLens(ctx)
}
