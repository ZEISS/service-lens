package components

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
)

// TagInputProps is a component that allows users to input tags.
type TagInputProps struct {
	ClassNames htmx.ClassNames
}

// TagInput ...
func TagInput(props TagInputProps) htmx.Node {
	return htmx.Group(
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("tags"),
			htmx.ID("tags-input"),
			htmx.HyperScript(`on newtag(tag)
  append`+"`,$tag`"+`to @value
  make a <li/> then put tag into it
  put it at the start of #shown-tags
end

on deltag(tag)
  call (@value of me).split(',')
  make a Set from it then set s to it
  call s.delete(tag) then call s.delete('')
  call Array.from(s) then set my @value to it.join(',')
  for el in #shown-tags.children
    if innerHTML of el == tag
      remove el
    end
  end
end

`),
		),
		htmx.Ul(
			htmx.ID("shown-tags"),
		),
		forms.FormControl(
			forms.FormControlProps{
				ClassNames: htmx.ClassNames{
					"py-4": true,
				},
			},
			forms.FormControlLabel(
				forms.FormControlLabelProps{},
				forms.FormControlLabelText(
					forms.FormControlLabelTextProps{
						ClassNames: htmx.ClassNames{
							"-my-4": true,
						},
					},
					htmx.Text("Tag"),
				),
			),
			forms.TextInputBordered(
				forms.TextInputProps{
					Name: "tags",
				},
				htmx.HyperScript(`on keyup[key is 'Enter']
  halt the event
  set taginput to me.value
  send newtag(tag: taginput) to #tags-input
  set me.value to ''
end`,
				),
			),
		),
	)
}

// <form>
//   <label for="item">Item: </label><input name="item">
//   <input type="hidden" name="tags" id="tags-input"
//   _="
// on newtag(tag)
//        append `,$tag` to @value
//        make a <li/> then put tag into it
//        put it at the start of #shown-tags
//      end

//      on deltag(tag)
//        call (@value of me).split(',')
//        make a Set from it then set s to it
//        call s.delete(tag) then call s.delete('')
//        call Array.from(s) then set my @value to it.join(',')
//        for el in #shown-tags.children
//          if innerHTML of el == tag
//            remove el
//          end
//        end
//      end">
//   <ul id="shown-tags"
//   _="on click send deltag(tag: innerHTML of target) to #tags-input"></ul>
// </form>
// <form id="tagform"
//  _="on submit
//       halt the event
//       set taginput to (first <input/> in me).value
//       send newtag(tag: taginput) to #tags-input
//       set value of first <input/> in me to ''">
//   <label for="tags">Tags: </label><input name="tags"> (Hit enter to add a tag)
// </form>
