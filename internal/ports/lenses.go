package ports

import "context"

// Lenses ...
type Lenses interface {
	AddLens(ctx context.Context) error
}
