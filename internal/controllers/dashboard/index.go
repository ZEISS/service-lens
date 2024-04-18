package dashboard

import (
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// HelloWorld ...
func HelloWorld(children ...htmx.Node) htmx.Node {
	return htmx.Element("hello-world", children...)
}

// DashboardIndexController ...
type DashboardIndexController struct {
	db ports.Repository

	htmx.DefaultController
}

// NewDashboardIndexController ...
func NewDashboardController(db ports.Repository) *DashboardIndexController {
	return &DashboardIndexController{
		db: db,
	}
}

// Prepare ...
func (d *DashboardIndexController) Prepare() error {
	if err := d.BindValues(utils.User(d.db)); err != nil {
		return err
	}

	return nil
}

// Error ...
func (d *DashboardIndexController) Error(err error) error {
	return d.Hx().RenderComp(
		htmx.Text(err.Error()),
	)
}

// Get ...
func (d *DashboardIndexController) Get() error {
	return d.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: d.Values(utils.ValuesKeyUser).(*authz.User),
				},
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
