package login

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
)

func NewLogin() htmx.Node {
	return htmx.Fragment(
		htmx.Section(
			htmx.Merge(
				htmx.ClassNames{
					"bg-gray-50":       true,
					"dark:bg-gray-900": true,
				},
			),
		),
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"flex":           true,
					"flex-col":       true,
					"items-center":   true,
					"justify-center": true,
					"px-6":           true,
					"py-8":           true,
					"mx-auto":        true,
					"md:h-screen":    true,
					"lg:py-0":        true,
				},
			),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						"w-96":     true,
						"max-w-lg": true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text("Sign in to your account"),
					),
					htmx.Div(
						htmx.ClassNames{
							"mt-4": true,
						},
						links.Button(
							links.LinkProps{
								ClassNames: htmx.ClassNames{
									"w-full":      true,
									"btn-outline": true,
								},
								Href: "/login/entraid",
							},
							htmx.Text("Login on Microsoft Entra ID"),
						),
					),
					htmx.Div(
						htmx.ClassNames{
							"mt-4": true,
						},
						links.Button(
							links.LinkProps{
								ClassNames: htmx.ClassNames{
									"w-full":      true,
									"btn-outline": true,
								},
								Href: "/login/github",
							},
							htmx.Text("Login on GitHub"),
						),
					),
					dividers.Divider(
						dividers.DividerProps{},
						htmx.Text("OR"),
					),
					htmx.Form(
						htmx.HxPost("/login"),
						forms.FormControl(
							forms.FormControlProps{
								ClassNames: htmx.ClassNames{
									"py-4": true,
								},
							},
							forms.TextInputBordered(
								forms.TextInputProps{
									Name:        "username",
									Placeholder: "indy@jones.com",
								},
							),
						),
						forms.FormControl(
							forms.FormControlProps{},
							forms.TextInputBordered(
								forms.TextInputProps{
									Name:        "password",
									Placeholder: "supersecret",
								},
								htmx.Type("password"),
							),
						),
						cards.Actions(
							cards.ActionsProps{
								ClassNames: htmx.ClassNames{
									"py-4":  true,
									"-mb-4": true,
								},
							},
							buttons.Outline(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"w-full": true,
									},
								},
								htmx.Text("Login"),
							),
						),
					),
				),
			),
		),
	)
}
