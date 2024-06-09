package dashboard

import (
	"context"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
)

// ShowDashboardController ...
type ShowDashboardController struct {
	user  adapters.GothUser
	store ports.Datastore
	htmx.DefaultController
}

// NewShowDashboardController ...
func NewShowDashboardController(store ports.Datastore) *ShowDashboardController {
	return &ShowDashboardController{
		user:  adapters.GothUser{},
		store: store,
	}
}

// Prepare ...
func (d *ShowDashboardController) Prepare() error {
	return d.store.ReadTx(d.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetUser(ctx, &d.user)
	})
}

// Get ...
func (d *ShowDashboardController) Get() error {
	return d.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: d.user,
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
