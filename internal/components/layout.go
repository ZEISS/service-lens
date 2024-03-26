package components

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/service-lens/internal/resolvers"
)

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	Children []htmx.Node

	htmx.Ctx
}

// WrapProps ...
type WrapProps struct {
	ClassName map[string]bool
}

// Wrap ...
func Wrap(p WrapProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"p-6": true,
		}.Merge(p.ClassName),
		htmx.Group(children...),
	)
}

// Layout is a whole document to output.
func Layout(ctx htmx.Ctx, p LayoutProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		drawers.Drawer(
			drawers.DrawerProps{
				ID: "drawer",
				ClassNames: htmx.ClassNames{
					"lg:drawer-open": true,
				},
			},
			drawers.DrawerContent(
				drawers.DrawerContentProps{
					ID: "drawer",
				},
				SubNav(
					SubNavProps{},
					drawers.DrawerOpenButton(
						drawers.DrawerOpenProps{
							ID: "drawer",
							ClassNames: htmx.ClassNames{
								"lg:hidden": true,
							},
						},
						icons.Bars3Outline(
							icons.IconProps{},
						),
					),
				),
				Wrap(
					WrapProps{},
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
						ctx,
						AccountSwitcherProps{},
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
						UserMenuProps{},
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
	team := ctx.Values(resolvers.ValuesKeyTeam).(*authz.Team)

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
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Workloads"),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: fmt.Sprintf("/%s/workloads/new", team.Slug),
							},
							htmx.Text("New workload"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: fmt.Sprintf("/%s/workloads/list", team.Slug),
							},
							htmx.Text("List workload"),
						),
					),
				),
			),

			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Lenses"),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: fmt.Sprintf("/%s/lenses/new", team.Slug),
							},
							htmx.Text("New Lens"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: fmt.Sprintf("/%s/lenses/list", team.Slug),
							},
							htmx.Text("List Lens"),
						),
					),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Profiles"),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: fmt.Sprintf("/%s/profiles/new", team.Slug),
							},
							htmx.Text("New Profile"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: fmt.Sprintf("/%s/profiles/list", team.Slug),
							},
							htmx.Text("List Profile"),
						),
					),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Administration"),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: "/teams/list",
							},
							htmx.Text("List Teams"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href: "/site/settings/index",
							},
							htmx.Text("Site settings"),
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
						Href: "/me/index",
					},
					htmx.Text("Profiles"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href: "/site/settings/index",
					},
					htmx.Text("Site settings"),
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
