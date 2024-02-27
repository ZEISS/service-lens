package components

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
)

// AccountSwitcherProps ...
type AccountSwitcherProps struct {
	ctx *fiber.Ctx
}

// WithContext returns a new PageProps with the given context.
func (p AccountSwitcherProps) WithContext(ctx *fiber.Ctx) AccountSwitcherProps {
	p.ctx = ctx

	return p
}

// AccountSwitcher ...
func AccountSwitcher(p AccountSwitcherProps, children ...htmx.Node) htmx.Node {
	return dropdowns.Dropdown(
		dropdowns.DropdownProps{},
		dropdowns.DropdownButton(
			dropdowns.DropdownButtonProps{
				ClassNames: htmx.ClassNames{
					"btn":         true,
					"btn-sm":      true,
					"btn-outline": true,
				},
			},
			htmx.Text("CIT-CA"),
			icons.ChevronUpDown(icons.IconProps{}),
		),
		dropdowns.DropdownMenuItems(
			dropdowns.DropdownMenuItemsProps{},
			dropdowns.DropdownMenuItem(
				dropdowns.DropdownMenuItemProps{},
				htmx.A(
					htmx.Text("Item 1"),
				),
			),
			dropdowns.DropdownMenuItem(
				dropdowns.DropdownMenuItemProps{},
				htmx.A(
					htmx.Text("Item 2"),
				),
			),
			dropdowns.DropdownMenuItem(
				dropdowns.DropdownMenuItemProps{},
				htmx.A(
					htmx.Attribute("href", "/teams/new"),
					htmx.Text("Create Team"),
				),
			),
		),
	)
}
