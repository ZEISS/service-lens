package designs

import (
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/slices"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignReactionsProps ...
type DesignReactionsProps struct {
	User   adapters.GothUser
	Design models.Design
}

// DesignReactions ...
func DesignReactions(props DesignReactionsProps) htmx.Node {
	return htmx.Div(
		htmx.ID("design-reactions"),
		htmx.HxSwapOob(conv.String(true)),
		htmx.ClassNames{
			tailwind.Flex:        true,
			tailwind.ItemsCenter: true,
		},
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.CreateDesignReactionUrlFormat, props.Design.ID)),
			components.EmojiPicker(
				components.EmojiPickerProps{},
			),
		),
		htmx.Group(
			htmx.Map(props.Design.GetReactionsByValue(), func(reaction string, reactions []models.Reaction) htmx.Node {
				react := slices.Index(func(reaction models.Reaction) bool {
					return reaction.ReactorID == props.User.ID
				}, reactions...)
				return htmx.FormElement(
					htmx.IfElse(
						react != -1,
						htmx.HxDelete(fmt.Sprintf(utils.DeleteDesignReactionUrlFormat, props.Design.ID, reactions[react].ID)),
						htmx.HxPost(fmt.Sprintf(utils.CreateDesignReactionUrlFormat, props.Design.ID)),
					),
					htmx.Input(
						htmx.Type("hidden"),
						htmx.Name("reaction"),
						htmx.Value(reaction),
					),
					buttons.Button(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								tailwind.Mx1: true,
							},
						},
						htmx.Text(fmt.Sprintf("%s (%d)", reaction, (len(reactions)))),
					),
				)
			},
			)...,
		),
	)
}
