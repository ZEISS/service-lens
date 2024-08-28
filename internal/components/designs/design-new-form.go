package designs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/loading"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/utils"
)

// DesignNewFormProps ...
type DesignNewFormProps struct {
	ClassNames htmx.ClassNames
	Template   string
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
					tailwind.M2: true,
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
        value: '',
        init() {
            let editor = new SimpleMDE({
              element: this.$refs.editor,
              previewRender: function(plainText, preview) {
                htmx.ajax('POST', '/preview', {values: {body: plainText}, target: '.editor-preview', swap: 'innerHTML'})

		            return "Loading...";
	            }
            })

           //  editor.value(this.$ref.editor.innerHTML)

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
							htmx.Text(props.Template),
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
						htmx.Text("Save Design"),
					),
				),
			),
		),
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
					htmx.Text("Workflow"),
				),
				forms.FormControl(
					forms.FormControlProps{
						ClassNames: htmx.ClassNames{},
					},
					htmx.Div(
						htmx.ClassNames{
							tailwind.Flex:           true,
							tailwind.JustifyBetween: true,
						},
						forms.Datalist(
							forms.DatalistProps{
								ID:          "workflows",
								Name:        "workflow",
								Placeholder: "Select workflows ...",
								URL:         utils.SearchWorkflowsUrlFormat,
							},
						),
						loading.Spinner(
							loading.SpinnerProps{
								ClassNames: htmx.ClassNames{
									"htmx-indicator": true,
									tailwind.M2:      true,
								},
							},
						),
					),
					forms.FormControlLabel(
						forms.FormControlLabelProps{},
						forms.FormControlLabelText(
							forms.FormControlLabelTextProps{
								ClassNames: htmx.ClassNames{
									tailwind.TextNeutral500: true,
								},
							},
							htmx.Text("The team that owns the account."),
						),
					),
				),
			),
		),
		components.AddTags(
			components.AddTagsProps{},
		),
	)
}
