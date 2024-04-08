package components

import (
	"fmt"
	"strings"

	"github.com/zeiss/service-lens/internal/resolvers"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/menus"
)

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	Children []htmx.Node

	htmx.Ctx
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
func Layout(ctx htmx.Ctx, p LayoutProps, children ...htmx.Node) htmx.Node {
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
							User: ctx.Values(resolvers.ValuesKeyUser).(*authz.User),
						},
					),
					MainMenu(
						ctx,
						MainMenuProps{},
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
							Path: ctx.Path(),
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

	htmx.Ctx
}

// MainMenu ...
func MainMenu(ctx htmx.Ctx, p MainMenuProps, children ...htmx.Node) htmx.Node {
	team, ok := ctx.Values(resolvers.ValuesKeyTeam).(*authz.Team)
	if !ok {
		team = &authz.Team{}
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
						Href:   fmt.Sprintf("/%s", team.Slug),
						Active: ctx.Path() == fmt.Sprintf("/%s", team.Slug),
					},
					htmx.Text("Dashboard"),
				),
			),
			htmx.If(
				team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(ctx.Path(), fmt.Sprintf("/%s/workloads", team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Workloads"),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/workloads/new", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/workloads/new", team.Slug),
								},
								htmx.Text("New workload"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/workloads", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/workloads", team.Slug),
								},
								htmx.Text("List workload"),
							),
						),
					),
				),
			),
			htmx.If(
				team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(ctx.Path(), fmt.Sprintf("/%s/lenses", team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Lenses"),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/lenses/new", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/lenses/new", team.Slug),
								},
								htmx.Text("New Lens"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/lenses/list", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/lenses/list", team.Slug),
								},
								htmx.Text("List Lens"),
							),
						),
					),
				),
			),
			htmx.If(
				team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(ctx.Path(), fmt.Sprintf("/%s/profiles", team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Profiles"),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/profiles/new", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/profiles/new", team.Slug),
								},
								htmx.Text("New Profile"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/profiles/list", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/profiles/list", team.Slug),
								},
								htmx.Text("List Profile"),
							),
						),
					),
				),
			),
			htmx.If(
				team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(ctx.Path(), fmt.Sprintf("/%s/environments", team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Environments"),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/environments/new", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/environments/new", team.Slug),
								},
								htmx.Text("New Environment"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/%s/environments/list", team.Slug),
									Active: ctx.Path() == fmt.Sprintf("/%s/environments/list", team.Slug),
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
						Open: strings.HasPrefix(ctx.Path(), "/site"),
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
								Active: ctx.Path() == "/site/teams/new",
							},
							htmx.Text("New Team"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/site/teams",
								Active: ctx.Path() == "/site/teams",
							},
							htmx.Text("List Teams"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/site/settings",
								Active: ctx.Path() == "/site/settings",
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

	htmx.Ctx
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
		htmx.Input(
			htmx.ClassNames{
				"toggle":           true,
				"theme-controller": true,
				"mx-4":             true,
			},
			htmx.Attribute("type", "checkbox"),
			htmx.Value("cupcake"),
		),
	)
}
