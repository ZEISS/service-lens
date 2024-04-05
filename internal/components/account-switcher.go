package components

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
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
	User *authz.User
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
			htmx.Text("CIT-CA"),
			icons.ChevronUpDownOutline(icons.IconProps{}),
		),
		dropdowns.DropdownMenuItems(
			dropdowns.DropdownMenuItemsProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
			},
			utils.Map(func(el authz.Team) htmx.Node {
				return dropdowns.DropdownMenuItem(
					dropdowns.DropdownMenuItemProps{},
					links.Link(
						links.LinkProps{
							ClassNames: htmx.ClassNames{
								"link": false,
							},
							Href: fmt.Sprintf("/%s", el.Slug),
						},
						htmx.Text(el.Name),
					),
				)
			}, *props.User.Teams...),
		),
	)
}
