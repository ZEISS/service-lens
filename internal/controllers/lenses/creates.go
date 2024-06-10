package lenses

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

var validate *validator.Validate

// CreateLensControllerImpl ...
type CreateLensControllerImpl struct {
	lens  models.Lens
	store ports.Datastore
	htmx.DefaultController
}

// NewCreateLensController ...
func NewCreateLensController(store ports.Datastore) *CreateLensControllerImpl {
	return &CreateLensControllerImpl{
		store: store,
	}
}

// Prepare ...
func (l *CreateLensControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindQuery(&l.lens)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.lens)
	if err != nil {
		return err
	}

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
	err = l.lens.UnmarshalJSON(buf.Bytes())
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateLens(ctx, &l.lens)
	})
}

// Post ...
func (l *CreateLensControllerImpl) Post() error {
	return l.Redirect(fmt.Sprintf("/lenses/%s", l.lens.ID))
}
