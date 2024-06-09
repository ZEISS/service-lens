package components

import (
	"strings"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/menus"
)

// MainMenuProps ...
type MainMenuProps struct {
	ClassNames htmx.ClassNames
	Path       string
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
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{
						Open: strings.HasPrefix(p.Path, "/workloads"),
					},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Workloads"),
					),
					menus.MenuItem(
						menus.MenuItemProps{
							ClassNames: htmx.ClassNames{
								"hover:bg-base-300": false,
							},
						},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/workloads/new",
								Active: p.Path == "/workloads/new",
							},
							htmx.Text("New workload"),
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
								// Href:   fmt.Sprintf("/teams/%s/workloads/list", p.Team.Slug),
								// Active: p.Path == fmt.Sprintf("/teams/%s/workloads/list", p.Team.Slug),
							},
							htmx.Text("List workload"),
						),
					),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{
						Open: false,
					},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Lenses"),
					),
					menus.MenuItem(
						menus.MenuItemProps{
							ClassNames: htmx.ClassNames{
								"hover:bg-base-300": false,
							},
						},
						menus.MenuLink(
							menus.MenuLinkProps{
								// Href:   fmt.Sprintf("/teams/%s/lenses/new", p.Team.Slug),
								// Active: p.Path == fmt.Sprintf("/teams/%s/lenses/new", p.Team.Slug),
							},
							htmx.Text("New Lens"),
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
								// Href:   fmt.Sprintf("/teams/%s/lenses/list", p.Team.Slug),
								// Active: p.Path == fmt.Sprintf("/teams/%s/lenses/list", p.Team.Slug),
							},
							htmx.Text("List Lens"),
						),
					),
				),
			),

			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{
						Open: strings.HasPrefix(p.Path, "/profiles"),
					},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Profiles"),
					),
					menus.MenuItem(
						menus.MenuItemProps{
							ClassNames: htmx.ClassNames{
								"hover:bg-base-300": false,
							},
						},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/profiles/new",
								Active: p.Path == "/profiles/new",
							},
							htmx.Text("New Profile"),
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
								Active: p.Path == "/profiles",
							},
							htmx.Text("List Profile"),
						),
					),
				),
			),

			menus.MenuItem(
				menus.MenuItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{
						Open: strings.HasPrefix(p.Path, "/environments"),
					},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Environments"),
					),
					menus.MenuItem(
						menus.MenuItemProps{
							ClassNames: htmx.ClassNames{
								"hover:bg-base-300": false,
							},
						},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/environments/new",
								Active: p.Path == "/environments/new",
							},
							htmx.Text("New Environment"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/environments",
								Active: p.Path == "/environments",
							},
							htmx.Text("List Environment"),
						),
					),
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
								Href:   "/site/settings",
								Active: p.Path == "/site/settings",
							},
							htmx.Text("Settings"),
						),
					),
				),
			),
		),
	)
}
