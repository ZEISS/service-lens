package dashboard

import (
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
	db  ports.Repository
	ctx htmx.Ctx

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
	ctx, err := htmx.NewDefaultContext(d.Hx().Ctx(), utils.Team(d.Hx().Ctx(), d.db), utils.User(d.Hx().Ctx(), d.db))
	if err != nil {
		return err
	}
	d.ctx = ctx

	return nil
}

// Get ...
func (d *DashboardIndexController) Get() error {
	return d.Hx().RenderComp(
		components.Page(
			d.ctx,
			components.PageProps{},
			components.Layout(
				d.ctx,
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
