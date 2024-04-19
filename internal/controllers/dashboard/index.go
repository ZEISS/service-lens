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

// ShowDashboardController ...
type ShowDashboardController struct {
	db ports.Repository

	htmx.DefaultController
}

// NewShowDashboardController ...
func NewShowDashboardController(db ports.Repository) *ShowDashboardController {
	return &ShowDashboardController{
		db: db,
	}
}

// Prepare ...
func (d *ShowDashboardController) Prepare() error {
	if err := d.BindValues(utils.User(d.db)); err != nil {
		return err
	}

	return nil
}

// Get ...
func (d *ShowDashboardController) Get() error {
	return d.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: htmx.AsValue[*authz.User](d.Values(utils.ValuesKeyUser)),
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
