package dashboard

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// HelloWorld ...
func HelloWorld(children ...htmx.Node) htmx.Node {
	return htmx.Element("hello-world", children...)
}

// DashboardIndexController ...
type DashboardIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewDashboardIndexController ...
func NewDashboardController(db ports.Repository) *DashboardIndexController {
	return &DashboardIndexController{
		db: db,
	}
}

// Prepare ...
func (d *DashboardIndexController) Prepare() error {
	return nil
}

// Get ...
func (d *DashboardIndexController) Get() error {
	return d.Hx().RenderComp(
		components.Page(
			d.Hx(),
			components.PageProps{},
			components.Layout(
				d.Hx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.Form(
						// htmx.Attribute("data-target", "hello-world.form"),
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
