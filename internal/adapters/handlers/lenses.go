package handlers

import (
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

type lensesHandler struct {
	lc ports.Lenses
}

// NewLensesHandler returns a new LensesHandler.
func NewLensesHandler(lc ports.Lenses) *lensesHandler {
	return &lensesHandler{lc}
}

// Index is the handler for the index page.
func (p *lensesHandler) Index() fiber.Handler {
	return htmx.NewCompHandler(
		components.Page(
			components.PageProps{
				Children: []htmx.Node{
					htmx.P(htmx.Text("Lenses")),
				},
			},
		),
	)
}

// NewLens is the handler for the new lens page.
func (p *lensesHandler) NewLens() htmx.HtmxHandlerFunc {
	return func(hx *htmx.Htmx) error {
		profile := &models.Profile{
			Name:        hx.Ctx().FormValue("name"),
			Description: hx.Ctx().FormValue("description"),
		}

		err := p.lc.AddLens(hx.Ctx().Context())
		if err != nil {
			return err
		}

		hx.Redirect("/profiles/" + profile.ID.String())

		return nil
	}
}

// New is the handler for the new lens page.
func (p *lensesHandler) New() fiber.Handler {
	return htmx.NewCompHandler(
		components.Page(
			components.PageProps{
				Children: []htmx.Node{
					htmx.FormElement(
						htmx.ID("new-lens-form"),
						htmx.HxPost("/lenses/new"),
						htmx.HxEncoding("multipart/form-data"),
						htmx.LabElement(
							htmx.ClassNames{
								"form-control": true,
								"w-full":       true,
								"max-w-xs":     true,
							},
							htmx.Input(
								htmx.Attribute("type", "file"),
								htmx.ClassNames{
									"file-input":          true,
									"file-input-bordered": true,
									"w-full":              true,
									"max-w-xs":            true,
								},
							),
						),
						htmx.Button(
							htmx.ClassNames{
								"btn":         true,
								"btn-default": true,
								"my-4":        true,
							},
							htmx.Attribute("type", "submit"),
							htmx.Text("Create Profile"),
						),
					),
				},
			},
		),
	)
}
