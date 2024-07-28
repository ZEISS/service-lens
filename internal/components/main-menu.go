package components

import (
	"strings"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/service-lens/internal/utils"
)

// MainMenuProps ...
type MainMenuProps struct {
	ClassNames htmx.ClassNames
	Path       string
	Team       string
}

// MainMenu ...
func MainMenu(p MainMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Nav(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		menus.Menu(
			menus.MenuProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
			},
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/",
						Active: p.Path == "/",
					},
					htmx.Text("Dashboard"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/designs",
						Active: strings.HasPrefix(p.Path, "/designs"),
					},
					htmx.Text("Designs"),
				),
			),

			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/workloads",
						Active: strings.HasPrefix(p.Path, "/workloads"),
					},
					htmx.Text("Workloads"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/lenses",
						Active: strings.HasPrefix(p.Path, "/lenses"),
					},
					htmx.Text("Lenses"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/profiles",
						Active: strings.HasPrefix(p.Path, "/profiles"),
					},
					htmx.Text("Profiles"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/environments",
						Active: strings.HasPrefix(p.Path, "/environments"),
					},
					htmx.Text("Environments"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   utils.ListTagsUrlFormat,
						Active: strings.HasPrefix(p.Path, "/tags"),
					},
					htmx.Text("Tags"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{
						Open: strings.HasPrefix(p.Path, "/site"),
					},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Administration"),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/site/teams/new",
								Active: p.Path == "/site/teams/new",
							},
							htmx.Text("New Team"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/site/teams",
								Active: p.Path == "/site/teams",
							},
							htmx.Text("List Teams"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/settings",
								Active: p.Path == "/settings",
							},
							htmx.Text("Settings"),
						),
					),
				),
			),
		),
	)
}
