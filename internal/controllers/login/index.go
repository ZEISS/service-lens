package login

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// LoginIndexController ...
type LoginIndexController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewLoginIndexController ...
func NewLoginIndexController(db ports.Repository) *LoginIndexController {
	return &LoginIndexController{db, htmx.UnimplementedController{}}
}

// Get ...
func (l *LoginIndexController) Get() error {
	return l.Hx().RenderComp(
		components.Page(
			l.Hx(),
			components.PageProps{},
			components.Wrap(
				components.WrapProps{},
				htmx.A(
					htmx.Attribute("href", "/login/github"),
					htmx.Text("Login with GitHub"),
				),
			),
		),
	)
}
