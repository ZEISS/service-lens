package components

import (
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"

	htmx "github.com/zeiss/fiber-htmx"
)

// AddTagsProps ...
type AddTagsProps struct {
	ClassNames htmx.ClassNames
}

// AddTags ...
func AddTags(props AddTagsProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
				},
				props.ClassNames,
			),
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
								alpine.XBind("name", "`tags.${index}.name`"),
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
								alpine.XBind("name", "`tags.${index}.value`"),
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
					buttons.Button(
						buttons.ButtonProps{
							Type: "button",
						},
						alpine.XOn("click", "addTag()"),
						htmx.Text("Add Tag"),
					),
				),
			),
		),
	)
}
