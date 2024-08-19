package workflows

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alerts"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// WorkflowShowControllerImpl ...
type WorkflowShowControllerImpl struct {
	workflow models.Workflow
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.UnimplementedController
}

// NewWorkflowShowController ...
func NewWorkflowShowController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkflowShowControllerImpl {
	return &WorkflowShowControllerImpl{store: store}
}

// Error ...
func (p *WorkflowShowControllerImpl) Error(err error) error {
	return p.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Error",
				Path:  p.Path(),
				User:  p.Session().User,
			},
			func() htmx.Node {
				return alerts.Error(alerts.AlertProps{}, htmx.Text(err.Error()))
			},
		),
	)
}

// Prepare ...
func (p *WorkflowShowControllerImpl) Prepare() error {
	err := p.BindParams(&p.workflow)
	if err != nil {
		return err
	}

	return p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetWorkflow(ctx, &p.workflow)
	})
}

// Get ...
func (p *WorkflowShowControllerImpl) Get() error {
	return p.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: p.workflow.Name,
				Path:  p.Path(),
				User:  p.Session().User,
				Head:  []htmx.Node{},
			},
			func() htmx.Node {
				return htmx.Fragment(
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
									htmx.Text(p.workflow.Name),
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
									htmx.Text(p.workflow.Description),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{},
									htmx.Text("Edit"),
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(""),
									htmx.HxConfirm("Are you sure you want to delete this workflow?"),
									htmx.Text("Delete"),
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
						alpine.XData(`{
            }`),
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Steps"),
								htmx.Div(
									htmx.ClassNames{
										"flex":     true,
										"flex-col": true,
										"py-2":     true,
									},
								),
								htmx.Ul(
									htmx.Attribute("x-sort", ""),
									htmx.Li(
										htmx.Attribute("x-sort:item", ""),
										htmx.Text("Step 1"),
									),
									htmx.Li(
										htmx.Attribute("x-sort:item", ""),
										htmx.Text("Step 1"),
									),
									htmx.Li(
										htmx.Attribute("x-sort:item", ""),
										htmx.Text("Step 1"),
									),
								),
							),
						),
					),
				)
			},
		),
	)
}
