package components

import (
	"fmt"
	"strings"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/menus"
)

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	Children []htmx.Node
	Team     *authz.Team
	User     adapters.GothUser
	Path     string
}

// WrapProps ...
type WrapProps struct {
	ClassNames htmx.ClassNames
}

// Wrap ...
func Wrap(p WrapProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Layout is a whole document to output.
func Layout(p LayoutProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		drawers.Drawer(
			drawers.DrawerProps{
				ID: "global-drawer",
				ClassNames: htmx.ClassNames{
					"lg:drawer-open": true,
				},
			},
			drawers.DrawerContent(
				drawers.DrawerContentProps{
					ID: "drawer",
				},
				SubNav(
					SubNavProps{
						ClassNames: htmx.ClassNames{
							"lg:hidden": true,
						},
					},
					drawers.DrawerOpenButton(
						drawers.DrawerOpenProps{
							ID: "global-drawer",
							ClassNames: htmx.ClassNames{
								"lg:hidden":   true,
								"btn-md":      true,
								"btn-square":  true,
								"btn-outline": true,
								"btn-primary": false,
							},
						},
						icons.Bars3Outline(
							icons.IconProps{},
						),
					),
				),
				Wrap(
					WrapProps{
						ClassNames: htmx.ClassNames{
							"m-6": true,
						},
					},
					htmx.Group(children...),
				),
			),
			drawers.DrawerSide(
				drawers.DrawerSideProps{
					ID: "drawer",
				},
				drawers.DrawerSideMenu(
					drawers.DrawerSideMenuProps{},
					AccountSwitcher(
						AccountSwitcherProps{
							User: p.User,
						},
					),
					MainMenu(
						MainMenuProps{
							Team: p.Team,
						},
					),
					dividers.Divider(
						dividers.DividerProps{
							ClassNames: htmx.ClassNames{
								"my-0": true,
							},
						},
					),
					UserMenu(
						UserMenuProps{
							Path: p.Path,
						},
					),
				),
			),
		),
	)
}

// MainMenuProps ...
type MainMenuProps struct {
	ClassNames htmx.ClassNames
	Team       *authz.Team
	Path       string
}

// MainMenu ...
func MainMenu(p MainMenuProps, children ...htmx.Node) htmx.Node {
	if p.Team == nil {
		p.Team = &authz.Team{}
	}

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
						Href:   fmt.Sprintf("/%s", p.Team.Slug),
						Active: p.Path == fmt.Sprintf("/%s", p.Team.Slug),
					},
					htmx.Text("Dashboard"),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/workloads", p.Team.Slug)),
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
									Href:   fmt.Sprintf("/teams/%s/workloads/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/workloads/new", p.Team.Slug),
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
									Href:   fmt.Sprintf("/teams/%s/workloads/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/workloads/list", p.Team.Slug),
								},
								htmx.Text("List workload"),
							),
						),
					),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/lenses", p.Team.Slug)),
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
									Href:   fmt.Sprintf("/teams/%s/lenses/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/lenses/new", p.Team.Slug),
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
									Href:   fmt.Sprintf("/teams/%s/lenses/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/lenses/list", p.Team.Slug),
								},
								htmx.Text("List Lens"),
							),
						),
					),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/profiles", p.Team.Slug)),
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
									Href:   fmt.Sprintf("/teams/%s/profiles/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/profiles/new", p.Team.Slug),
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
									Href:   fmt.Sprintf("/teams/%s/profiles/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/profiles/list", p.Team.Slug),
								},
								htmx.Text("List Profile"),
							),
						),
					),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/environments", p.Team.Slug)),
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
									Href:   fmt.Sprintf("/teams/%s/environments/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/environments/new", p.Team.Slug),
								},
								htmx.Text("New Environment"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/environments/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/environments/list", p.Team.Slug),
								},
								htmx.Text("List Environment"),
							),
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

// UserMenuProps ...
type UserMenuProps struct {
	ClassNames htmx.ClassNames
	Path       string
}

// UserMenu ...
func UserMenu(p UserMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
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
						Href:   "/me",
						Active: p.Path == "/me",
					},
					htmx.Text("Profile"),
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
