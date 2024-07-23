package designs

import (
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// NewDesignControllerImpl ...
type NewDesignControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewDesignController ...
func NewDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewDesignControllerImpl {
	return &NewDesignControllerImpl{store: store}
}

// Get ...
func (l *NewDesignControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: l.Path(),
			},
			designs.DesignNewForm(
				designs.DesignNewFormProps{},
			),
		),
	)
}
