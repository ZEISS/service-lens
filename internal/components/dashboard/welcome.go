package dashboard

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// WelcomeCardProps ...
type WelcomeCardProps struct {
	// ClassNames ...
	ClassNames htmx.ClassNames
	// User ...
	User adapters.GothUser
}

// WelcomeCard ...
func WelcomeCard(props WelcomeCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
				},
			),
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text("Welcome"),
			),
		),
	)
}
