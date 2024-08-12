package designs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// DesignNewFormProps ...
type DesignNewFormProps struct {
	ClassNames htmx.ClassNames
}

// DesignNewForm ...
func DesignNewForm(props DesignNewFormProps) htmx.Node {
	return htmx.FormElement(
		htmx.HxPost(""),
		htmx.HxTarget("this"),
		htmx.HxSwap("outerHTML"),
		htmx.ID("body"),
		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
			},
			cards.Body(
				cards.BodyProps{},
				cards.Title(
					cards.TitleProps{},
					htmx.Text("Create Design"),
				),
				forms.FormControl(
					forms.FormControlProps{
						ClassNames: htmx.ClassNames{},
					},
					forms.TextInputBordered(
						forms.TextInputProps{
							Name:        "title",
							Placeholder: "Title",
						},
					),
					forms.FormControlLabel(
						forms.FormControlLabelProps{},
						forms.FormControlLabelText(
							forms.FormControlLabelTextProps{
								ClassNames: htmx.ClassNames{
									"text-neutral-500": true,
								},
							},
							htmx.Text("The title must be from 3 to 2048 characters."),
						),
					),
				),
				forms.FormControl(
					forms.FormControlProps{
						ClassNames: htmx.ClassNames{},
					},
					htmx.Div(
						alpine.XData(`{
        value: '# Write Some Markdown...',
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
								Name: "body",
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
					buttons.Outline(
						buttons.ButtonProps{},
						htmx.Attribute("type", "submit"),
						htmx.Text("Save Design"),
					),
				),
			),
		),
		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					"my-2": true,
					"mx-2": true,
				},
			},
			cards.Body(
				cards.BodyProps{},
				cards.Title(
					cards.TitleProps{},
					htmx.Text("Tags - Optional"),
				),
				htmx.Div(
					alpine.XData(`{
            tags: [],
            addTag(tag) {
              this.tags.push({name: '', value: ''});
            },
            removeTag(index) {
              this.tags.splice(index, 1);
            }
          }`),
					htmx.Template(
						alpine.XFor("(tag, index) in tags"),
						htmx.Attribute(":key", "index"),
						htmx.Div(
							htmx.ClassNames{
								tailwind.Flex:    true,
								tailwind.SpaceX4: true,
							},
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{},
								},
								forms.TextInputBordered(
									forms.TextInputProps{},
									alpine.XModel("tag.name"),
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("Key in a tag."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{},
								},
								forms.TextInputBordered(
									forms.TextInputProps{},
									alpine.XModel("tag.value"),
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("Value in a tag."),
									),
								),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Outline(
							buttons.ButtonProps{
								Type: "button",
							},
							alpine.XOn("click", "addTag()"),
							htmx.Text("Add new tag"),
						),
					),
				),
			),
		),
	)
}
