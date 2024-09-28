package lenses

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// NewLensControllerImpl ...
type NewLensControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensController ...
func NewLensController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewLensControllerImpl {
	return &NewLensControllerImpl{store: store}
}

// Error ...
func (l *NewLensControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Post ...
func (l *NewLensControllerImpl) Post() error {
	var lens models.Lens

	spec, err := l.Ctx().FormFile("spec")
	if err != nil {
		return err
	}
	file, err := spec.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	err = lens.UnmarshalJSON(buf.Bytes())
	if err != nil {
		return err
	}

	lens.IsDraft = true // first draft the

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateLens(ctx, &lens)
	})
	if err != nil {
		return err
	}

	return l.Redirect(fmt.Sprintf(utils.ShowLensUrlFormat, lens.ID))
}
