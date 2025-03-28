package components

import (
	"strings"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/menus"
)

// UserMenuProps ...
type UserMenuProps struct {
	ClassNames htmx.ClassNames
	Path       string
}

// UserMenu ...
func UserMenu(p UserMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Nav(
		htmx.Merge(
			htmx.ClassNames{},
		),
		menus.Menu(
			menus.MenuProps{
				ClassNames: htmx.ClassNames{
					"w-full":      true,
					"bg-base-200": false,
				},
			},
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/settings",
						Active: strings.HasPrefix(p.Path, "/settings"),
					},
					htmx.Text("Settings"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/me",
						Active: strings.HasPrefix(p.Path, "/me"),
					},
					htmx.Text("Profile"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href: "/logout",
					},
					htmx.Text("Logout"),
				),
			),
		),
	)
}
