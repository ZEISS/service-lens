package components

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/resolvers"
)

// AccountSwitcherProps ...
type AccountSwitcherProps struct{}

// AccountSwitcher ...
func AccountSwitcher(ctx htmx.Ctx, p AccountSwitcherProps, children ...htmx.Node) htmx.Node {
	user, ok := ctx.Values(resolvers.ValuesKeyUser).(*authz.User)
	if !ok {
		return htmx.Text("User not found")
	}

	users := []htmx.Node{}
	for _, t := range *user.Teams {
		users = append(users, dropdowns.DropdownMenuItem(dropdowns.DropdownMenuItemProps{}, links.Link(links.LinkProps{Href: fmt.Sprintf("/%s/index", t.Slug)}, htmx.Text(t.Name))))
	}

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
			icons.ChevronUpDownOutline(icons.IconProps{}),
		),
		dropdowns.DropdownMenuItems(
			dropdowns.DropdownMenuItemsProps{},
			htmx.Group(users...),
		),
	)
}
