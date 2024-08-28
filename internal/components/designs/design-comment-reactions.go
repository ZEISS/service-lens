package designs

import (
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/slices"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignCommentReactionsProps ...
type DesignCommentReactionsProps struct {
	User    adapters.GothUser
	Design  models.Design
	Comment models.DesignComment
}

// DesignCommentReactions ...
func DesignCommentReactions(props DesignCommentReactionsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ID(fmt.Sprintf("reaction-%s", props.Comment.ID)),
		htmx.ClassNames{
			tailwind.Flex:        true,
			tailwind.ItemsCenter: true,
		},
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.CreateDesignCommentReactionUrlFormat, props.Design.ID, props.Comment.ID)),
			components.EmojiPicker(
				components.EmojiPickerProps{},
			),
		),
		htmx.Group(
			htmx.Map(props.Comment.GetReactionsByValue(), func(reaction string, reactions []models.Reaction) htmx.Node {
				react := slices.Index(func(reaction models.Reaction) bool {
					return reaction.ReactorID == props.User.ID
				}, reactions...)
				return htmx.FormElement(
					htmx.IfElse(
						react != -1,
						htmx.HxDelete(fmt.Sprintf(utils.DeleteDesignCommentReactionUrlFormat, props.Design.ID, props.Comment.ID, reactions[react].ID)),
						htmx.HxPost(fmt.Sprintf(utils.CreateDesignCommentReactionUrlFormat, props.Design.ID, props.Comment.ID)),
					),
					htmx.Input(
						htmx.Type("hidden"),
						htmx.Name("reaction"),
						htmx.Value(reaction),
					),
					buttons.Button(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								tailwind.M1: true,
							},
						},
						htmx.Text(fmt.Sprintf("%s (%d)", reaction, (len(reactions)))),
					),
				)
			},
			)...,
		),
		htmx.Group(children...),
	)
}
