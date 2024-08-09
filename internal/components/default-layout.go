package components

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
)

// DefaultLayoutProps ...
type DefaultLayoutProps struct {
	ClassNames htmx.ClassNames
	Path       string
	Title      string
	User       adapters.GothUser
}

// DefaultLayout ...
func DefaultLayout(props DefaultLayoutProps, children ...htmx.Node) htmx.Node {
	return Page(
		PageProps{
			Title: props.Title,
		},
		Layout(
			LayoutProps{
				Path: props.Path,
				User: props.User,
			},
			Wrap(
				WrapProps{},
				children...,
			),
		),
	)
}
