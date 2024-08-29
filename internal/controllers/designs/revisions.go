package designs

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// RevisionControllerImpl ...
type RevisionControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewRevisionController ...
func NewRevisionController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *RevisionControllerImpl {
	return &RevisionControllerImpl{store: store}
}

// Get ...
func (c *RevisionControllerImpl) Get() error {
	return c.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var params struct {
					DesignID uuid.UUID `json:"design_id" params:"id"`
				}
				errorx.Panic(c.BindParams(&params))

				results := tables.Results[models.DesignRevision]{}

				errorx.Panic(c.store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListDesignRevisions(ctx, params.DesignID, &results)
				}))

				return htmx.Ul(
					htmx.ForEach(results.GetRows(), func(c *models.DesignRevision, choiceIdx int) htmx.Node {
						return forms.FormControl(
							forms.FormControlProps{},
							forms.FormControlLabel(
								forms.FormControlLabelProps{
									ClassNames: htmx.ClassNames{
										tailwind.JustifyStart: true,
										tailwind.SpaceX2:      true,
									},
								},
								forms.Radio(
									forms.RadioProps{
										Name:    "revision",
										Value:   conv.String(c.ID),
										Checked: choiceIdx == 0, // todo(katallaxie): should be a default option in the model
									},
								),
								forms.FormControlLabelText(
									forms.FormControlLabelTextProps{},
									htmx.Text(fmt.Sprintf("Created at %s", c.CreatedAt.Format("2006-01-02 15:04:05"))),
								),
							),
						)
					})...,
				)
			}),
			func(err error) htmx.Node {
				return htmx.Div(
					htmx.Text("An error occurred: " + err.Error()),
				)
			},
		),
	)
}
