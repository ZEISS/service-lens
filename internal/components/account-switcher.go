package components

import (
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/utils"
)

// AccountSwitcherProps ...
type AccountSwitcherProps struct {
	// ClassNames ...
	ClassNames htmx.ClassNames
	// User ...
	User adapters.GothUser
}

// AccountSwitcher ...
func AccountSwitcher(props AccountSwitcherProps, children ...htmx.Node) htmx.Node {
	return dropdowns.Dropdown(
		dropdowns.DropdownProps{},
		dropdowns.DropdownButton(
			dropdowns.DropdownButtonProps{
				ClassNames: htmx.ClassNames{
					"btn":             true,
					"btn-sm":          true,
					"btn-outline":     true,
					"w-full":          true,
					"justify-between": true,
				},
			},
			htmx.Text("ZEISS"),
			icons.ChevronUpDownOutline(icons.IconProps{}),
		),
		dropdowns.DropdownMenuItems(
			dropdowns.DropdownMenuItemsProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
			},
			utils.Map(func(el adapters.GothTeam) htmx.Node {
				return dropdowns.DropdownMenuItem(
					dropdowns.DropdownMenuItemProps{},
					links.Link(
						links.LinkProps{
							ClassNames: htmx.ClassNames{
								"link": false,
							},
							Href: fmt.Sprintf("/teams/%s", el.Slug),
						},
						htmx.Text(el.Name),
					),
				)
			}, *props.User.Teams...),
			dividers.Divider(
				dividers.DividerProps{
					ClassNames: htmx.ClassNames{},
				},
			),
			dropdowns.DropdownMenuItem(
				dropdowns.DropdownMenuItemProps{},
				links.Link(
					links.LinkProps{
						ClassNames: htmx.ClassNames{
							"link": false,
						},
						Href: "/teams/new",
					},
					htmx.Text("Create team"),
				),
			),
		),
	)
}
