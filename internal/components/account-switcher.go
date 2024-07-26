package components

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
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
		),
	)
}
