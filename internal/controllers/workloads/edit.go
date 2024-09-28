package workloads

import (
	"context"
	"fmt"

	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/fiber-htmx/components/toasts"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"
)

// WorkloadEditControllerImpl ...
type WorkloadEditControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadEditController ...
func NewWorkloadEditController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadEditControllerImpl {
	return &WorkloadEditControllerImpl{store: store}
}

// Error ...
func (p *WorkloadEditControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Post ...
func (p *WorkloadEditControllerImpl) Post() error {
	var body struct {
		Name        string `json:"name" form:"name" validate:"required,min=3,max=128,alphanum"`
		Description string `json:"description" form:"description" validate:"omitempty,min=3,max=1024"`
	}

	var params struct {
		ID uuid.UUID `json:"id" form:"id" validate:"required"`
	}
	errorx.Panic(p.BindParams(&params))

	errorx.Panic(p.BindBody(&body))
	workload := models.Workload{}
	workload.ID = params.ID
	workload.Name = body.Name
	workload.Description = body.Description

	errorx.Panic(p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateWorkload(ctx, &workload)
	}))

	return p.Render(

		cards.CardBordered(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					tailwind.M2: true,
				},
			},
			htmx.HxTarget("this"),
			htmx.HxSwap("outerHTML"),
			cards.Body(
				cards.BodyProps{},
				cards.Title(
					cards.TitleProps{},
					htmx.Text("Overview"),
				),
				htmx.Div(
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
							htmx.Text("Name"),
						),
						htmx.H3(
							htmx.Text(workload.Name),
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
							htmx.Text("Description"),
						),
						htmx.H3(
							htmx.Text(workload.Description),
						),
					),
				),
				cards.Actions(
					cards.ActionsProps{},
					buttons.Button(
						buttons.ButtonProps{
							Type: "button",
						},
						htmx.Text("Edit"),
						htmx.HxGet(fmt.Sprintf(utils.EditWorkloadUrlFormat, workload.ID)),
						htmx.HxSwap("outerHTML"),
					),
					buttons.Button(
						buttons.ButtonProps{},
						htmx.HxDelete(""),
						htmx.HxConfirm("Are you sure you want to delete this workload?"),
						htmx.Text("Delete"),
					),
				),
			),
		),
	)
}

// Get ...
func (p *WorkloadEditControllerImpl) Get() error {
	return p.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var params struct {
					ID uuid.UUID `json:"id" form:"id" validate:"required"`
				}

				errorx.Panic(p.BindParams(&params))
				var workload models.Workload
				workload.ID = params.ID

				errorx.Panic(p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetWorkload(ctx, &workload)
				}))

				return htmx.FormElement(
					htmx.HxPost(fmt.Sprintf(utils.EditWorkloadUrlFormat, workload.ID)),
					htmx.HxTarget("this"),
					htmx.HxSwap("outerHTML"),
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
								htmx.Text("Overview"),
							),
							htmx.Div(
								forms.FormControl(
									forms.FormControlProps{},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{},
											htmx.Text("Name"),
										),
									),
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:        "name",
											Placeholder: "Shop System, Payment Gateway, etc.",
											Value:       workload.Name,
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
											htmx.Text("The name must be from 3 to 128 characters. At least 3 characters must be non-whitespace, only alphanumeric characters are allowed."),
										),
									),
								),
								forms.FormControl(
									forms.FormControlProps{},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{},
											htmx.Text("Description"),
										),
									),
									forms.TextareaBordered(
										forms.TextareaProps{
											Name:        "description",
											Placeholder: "This is a shop system that processes payments.",
										},
										htmx.Text(workload.Description),
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
												},
											},
											htmx.Text("This is optional. The description must be from 3 to 1024 characters."),
										),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								links.Button(
									links.LinkProps{
										ClassNames: htmx.ClassNames{
											"btn": true,
										},
										Href: fmt.Sprintf(utils.ShowWorkloadUrlFormat, workload.ID),
									},
									htmx.Text("Cancel"),
								),
								buttons.Button(
									buttons.ButtonProps{
										Type: "button",
									},
									htmx.Text("Save"),
									htmx.HxPost(fmt.Sprintf(utils.EditWorkloadUrlFormat, workload.ID)),
									htmx.HxSwap("outerHTML"),
								),
							),
						),
					),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		),
	)
}
