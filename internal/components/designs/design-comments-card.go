package designs

import (
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignCommentsCardProps ...
type DesignCommentsCardProps struct {
	User       adapters.GothUser
	ClassNames htmx.ClassNames
	Design     models.Design
}

// DesignCommentsCard ...
func DesignCommentsCard(props DesignCommentsCardProps) htmx.Node {
	return htmx.Fragment(
		htmx.Div(
			htmx.ID("comments"),
			htmx.Group(htmx.ForEach(tables.RowsPtr(props.Design.Comments), func(c *models.DesignComment, choiceIdx int) htmx.Node {
				return DesignComment(
					DesignCommentProps{
						Comment: cast.Value(c),
						User:    props.User,
						Design:  props.Design,
					},
				)
			})...),
		),
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.CreateDesignCommentUrlFormat, props.Design.ID)),
			htmx.HxTarget("#comments"),
			htmx.HxSwap("beforeend"),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						avatars.AvatarRoundSmall(
							avatars.AvatarProps{},
							htmx.Img(
								htmx.Attribute("src", cast.Value(props.User.Image)),
							),
						),
						htmx.Text("Add a comment"),
					),
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{},
						},
						htmx.StyleElement(htmx.Raw(
							`.CodeMirror, .CodeMirror-scroll {
	min-height: 200px;
}`,
						)),
						htmx.Div(
							alpine.XData(`{
            value: 'Start typing...',
            init() {
                let editor = new SimpleMDE({
                  element: this.$refs.editor,
				  status: false,
                  previewRender: function(plainText, preview) {
                    htmx.ajax('POST', '/preview', {values: {body: plainText}, target: '.editor-preview', swap: 'innerHTML'})

                    return "Loading...";
                  }
                })
                editor.value(this.value)
                editor.codemirror.on('change', () => {
                    this.value = editor.value()
                })
            },
        }`,
							),
							forms.TextareaBordered(
								forms.TextareaProps{
									ClassNames: htmx.ClassNames{
										"h-[50vh]": true,
									},
									Name: "comment",
								},
								alpine.XRef("editor"),
							),
						),
						forms.FormControlLabel(
							forms.FormControlLabelProps{},
							forms.FormControlLabelText(
								forms.FormControlLabelTextProps{
									ClassNames: htmx.ClassNames{
										"text-neutral-500": true,
									},
								},
								htmx.Text("Supports Markdown."),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Attribute("type", "submit"),
							htmx.Text("Comment"),
						),
					),
				),
			),
		),
	)
}
