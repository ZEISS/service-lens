package designs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/forms"
)

// EditorProps ...
type EditorProps struct {
	ClassNames htmx.ClassNames
	Content    string
}

// Editor ...
func Editor(props EditorProps) htmx.Node {
	return htmx.Div(
		alpine.XData(`{
value: '',
init() {
let editor = new SimpleMDE({
forceSync: true,
element: this.$refs.editor,
status: false,
previewRender: function(plainText, preview) {
htmx.ajax('POST', '/preview', {values: {body: plainText}, target: '.editor-preview', swap: 'innerHTML'})
	return "Loading...";
}
})

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
			htmx.Text(props.Content),
		),
	)
}
