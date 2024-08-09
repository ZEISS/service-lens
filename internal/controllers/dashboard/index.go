package dashboard

import (
	"fmt"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// ShowDashboardController ...
type ShowDashboardController struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewShowDashboardController ...
func NewShowDashboardController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowDashboardController {
	return &ShowDashboardController{
		store: store,
	}
}

// Get ...
func (d *ShowDashboardController) Get() error {
	return d.Render(
		components.Page(
			components.PageProps{
				Title: fmt.Sprintf("Dashboard - %s", d.Session().User.Name),
			},
			components.Layout(
				components.LayoutProps{
					User: d.Session().User,
					Path: d.Path(),
				},
				components.Wrap(
					components.WrapProps{},
					htmx.Form(
						htmx.Attribute("is", "chat-input"),
						htmx.Input(
							htmx.Attribute("type", "text"),
							htmx.Attribute("name", "user-message"),
						),
					),
				),
				htmx.Div(
					htmx.ID("messages"),
				),
			),
		),
	)
}
