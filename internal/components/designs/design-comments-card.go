package designs

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignCommentsCardProps ...
type DesignCommentsCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
}

// DesignCommentsCard ...
func DesignCommentsCard(props DesignCommentsCardProps) htmx.Node {
	return htmx.Fragment(
		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.Merge(
					htmx.ClassNames{
						"my-2": true,
						"mx-2": true,
					},
				),
			},
			cards.Body(
				cards.BodyProps{},
				htmx.Div(
					htmx.ID("comments"),
					htmx.Group(htmx.ForEach(tables.RowsPtr(props.Design.Comments), func(c *models.DesignComment, choiceIdx int) htmx.Node {
						return cards.CardBordered(
							cards.CardProps{
								ClassNames: htmx.ClassNames{
									"my-4": true,
								},
							},
							cards.Body(
								cards.BodyProps{},
								cards.Title(
									cards.TitleProps{},
									htmx.Text(
										c.CreatedAt.Format("2006-01-02 15:04:05"),
									),
									dropdowns.Dropdown(
										dropdowns.DropdownProps{},
										dropdowns.DropdownButton(
											dropdowns.DropdownButtonProps{
												ClassNames: htmx.ClassNames{
													"btn":    true,
													"btn-sm": true,
												},
											},
											icons.ChevronUpDownOutline(icons.IconProps{}),
										),
										dropdowns.DropdownMenuItems(
											dropdowns.DropdownMenuItemsProps{},
											dropdowns.DropdownMenuItem(
												dropdowns.DropdownMenuItemProps{},
												htmx.A(
													htmx.ClassNames{
														"btn":    true,
														"btn-sm": true,
													},
													htmx.Attribute("href", ""),
													htmx.Text("Edit"),
												),
											),
										),
									),
								),
								htmx.Text(c.Comment),
								cards.Actions(
									cards.ActionsProps{},
									avatars.AvatarRoundSmall(
										avatars.AvatarProps{},
										htmx.Img(
											htmx.Attribute("src", utils.PtrStr(c.Author.Image)),
										),
									),
								),
							),
						)
					})...),
				),
				htmx.FormElement(
					htmx.HxPost(fmt.Sprintf(utils.CreateDesignCommentUrlFormat, props.Design.ID)),
					htmx.HxTarget("#comments"),
					htmx.HxSwap("beforeend"),
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{},
								},
								htmx.Div(
									alpine.XData(`{
            value: 'Start typing...',
            init() {
                let editor = new SimpleMDE({
                  element: this.$refs.editor,
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
			),
		),
	)
}
