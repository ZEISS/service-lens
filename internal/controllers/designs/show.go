package designs

import (
	"context"
	"fmt"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// ShowDesignControllerImpl ...
type ShowDesignControllerImpl struct {
	Design models.Design
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewShowDesignController ...
func NewShowDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowDesignControllerImpl {
	return &ShowDesignControllerImpl{
		store: store,
	}
}

// Prepare ...
func (l *ShowDesignControllerImpl) Prepare() error {
	var params struct {
		ID string `uri:"id" validate:"required,uuid"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &l.Design)
	})
}

// Get ...
func (l *ShowDesignControllerImpl) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Overview"),
						),
						htmx.Div(
							htmx.H1(
								htmx.Text(l.Design.Title),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Created at"),
								),
								htmx.H3(
									htmx.Text(
										l.Design.CreatedAt.Format("2006-01-02 15:04:05"),
									),
								),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Updated at"),
								),
								htmx.H3(
									htmx.Text(
										l.Design.UpdatedAt.Format("2006-01-02 15:04:05"),
									),
								),
							),
						),
					),
				),
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Comments"),
						),
						htmx.Div(
							htmx.ID("comments"),
							htmx.Group(htmx.ForEach(tables.RowsPtr(l.Design.Comments), func(c *models.DesignComment, choiceIdx int) htmx.Node {
								return cards.CardBordered(
									cards.CardProps{
										ClassNames: htmx.ClassNames{
											"my-4": true,
										},
									},
									cards.Body(
										cards.BodyProps{},
										htmx.Text(c.Comment),
									),
								)
							})...),
						),

						htmx.FormElement(
							htmx.HxPost(fmt.Sprintf(utils.CreateDesignCommentUrlFormat, l.Design.ID)),
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
										forms.TextareaBordered(
											forms.TextareaProps{
												ClassNames: htmx.ClassNames{
													"h-32": true,
												},
												Name:        "comment",
												Placeholder: "Add a comment...",
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
												htmx.Text("Supports Markdown."),
											),
										),
									),
									cards.Actions(
										cards.ActionsProps{},
										buttons.Outline(
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
			),
		),
	)
}
