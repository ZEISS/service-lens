package components

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/resolvers"

	htmx "github.com/zeiss/fiber-htmx"
)

// UserNavProps ...
type UserNavProps struct{}

// UserNav ...
func UserNav(ctx htmx.Ctx, p UserNavProps) htmx.Node {
	user, ok := ctx.Values(resolvers.ValuesKeyUser).(*authz.User)
	if !ok {
		return htmx.Text("User not found")
	}

	users := []htmx.Node{}
	for _, u := range *user.Teams {
		users = append(users, htmx.Li(htmx.Text(u.Name)))
	}

	fmt.Println(user.Teams)

	return htmx.Div(
		htmx.ClassNames{"dropdown": true, "dropdown-end": true},
		htmx.Div(
			htmx.Attribute("tabindex", "0"),
			htmx.ClassNames{
				"btn":        true,
				"btn-ghost":  true,
				"btn-circle": true,
				"avatar":     true,
			},
			htmx.Div(
				htmx.ClassNames{
					"w-10":         true,
					"rounded-full": true,
				},
				htmx.Img(
					htmx.Attribute("alt", "Tailwind CSS Navbar component"),
					htmx.Attribute("src", "https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"),
				),
			),
		),
		htmx.Ul(
			htmx.Attribute("tabindex", "0"),
			htmx.ClassNames{
				"mt-3":             true,
				"z-[1]":            true,
				"p-2":              true,
				"shadow":           true,
				"menu":             true,
				"menu-sm":          true,
				"dropdown-content": true,
				"bg-base-100":      true,
				"rounded-box":      true,
				"w-52":             true,
			},
			htmx.Group(users...),
			htmx.Li(
				htmx.A(
					htmx.ClassNames{"justify-between": true},
					htmx.Text("Profile"),
					htmx.Span(htmx.ClassNames{"badge": true}, htmx.Text("New")),
				),
			),
			htmx.Li(
				links.Link(
					links.LinkProps{
						Href:       "/settings/list",
						ClassNames: htmx.ClassNames{"underline-none": true},
					},
					htmx.Text(
						"Settings",
					),
				),
			),
			htmx.Li(
				htmx.A(
					htmx.Attribute("href", "/logout"),
					htmx.Text("Logout"),
				),
			),
		),
	)
}
