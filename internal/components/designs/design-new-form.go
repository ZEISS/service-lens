package designs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/loading"
	"github.com/zeiss/fiber-htmx/components/tabs"
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
					tabs.TabsBoxed(
						tabs.TabsProps{},
						tabs.Tab(
							tabs.TabProps{
								Active: true,
								Label:  "Write",
								Name:   "editor",
								ClassNames: htmx.ClassNames{
									tailwind.P10: false,
									tailwind.P2:  true,
								},
							},
							forms.TextareaBordered(
								forms.TextareaProps{
									ClassNames: htmx.ClassNames{
										tailwind.WFull: true,
										tailwind.H96:   true,
									},
									Name:        "body",
									Placeholder: "Write your design here...",
								},
								htmx.Class("mt-3 d-block width-full"),
								htmx.ID("body"),
								htmx.HxPost("/preview"),
								htmx.HxTarget("#preview"),
								htmx.HxSwap("innerHTML"),
								htmx.Attribute("hx-sync", "closest form:abort"),
								htmx.ContentEditable("false"),
								htmx.Text(props.Template),
							),
						),
						tabs.Tab(
							tabs.TabProps{
								Label: "Preview",
								Name:  "editor",
								ClassNames: htmx.ClassNames{
									tailwind.P10: false,
									tailwind.P4:  true,
								},
							},
							htmx.ID("preview"),
							htmx.Text("Whoops! You have not written anything yet."),
						),
						components.MarkdownToolbar(
							htmx.ClassNames{
								tailwind.Flex:           true,
								tailwind.JustifyBetween: true,
							},
							htmx.For("body"),
							htmx.Role("toolbar"),
							components.MarkdownBold(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("0"),
								icons.BoldOutline(
									icons.IconProps{},
								),
							),
							components.MarkdownItalic(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.ItalicOutline(
									icons.IconProps{},
								),
							),
							components.MarkdownQuote(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.ArrowUturnLeftOutline(
									icons.IconProps{},
								),
							),
							components.MarkdownCode(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.ItalicCodeBracketOutline(
									icons.IconProps{},
								),
							),
							components.MarkdownLink(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.LinkOutline(
									icons.IconProps{},
								),
							),
							components.MarkdownImage(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.ImageOutline(
									icons.IconProps{},
								),
							),
							components.MarkdownUnorderedList(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.ListBulletOutline(
									icons.IconProps{},
								),
							),
							components.MarkdownOrderedList(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.NumberedListOutline(
									icons.IconProps{},
								),
							),
							// components.MarkdownTaskList(
							// 	htmx.Class("btn btn-ghost btn-sm"),
							// 	htmx.TabIndex("-1"),
							// 	htmx.Text("task-list"),
							// ),
							components.MarkdownMention(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.AtSymbolOutline(
									icons.IconProps{},
								),
							),
							// components.MarkdownRef(
							// 	htmx.Class("btn btn-sm"),
							// 	htmx.TabIndex("-1"),
							// 	icons.
							// ),
							components.MarkdownStrikethrough(
								htmx.Class("btn btn-ghost btn-sm"),
								htmx.TabIndex("-1"),
								icons.StrikethroughOutline(
									icons.IconProps{},
								),
							),
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
								Name:        "workflow_id",
								Placeholder: "Search a workflow ...",
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
							htmx.Text("Optional - Select a workflow to associate with this design."),
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
