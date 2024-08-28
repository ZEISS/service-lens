package workflows

import (
	"context"
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alerts"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/loading"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/workflows"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
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
				Title:       "Error",
				Path:        p.Path(),
				User:        p.Session().User,
				Development: p.IsDevelopment(),
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
				Title:       p.workflow.Name,
				Path:        p.Path(),
				User:        p.Session().User,
				Development: p.IsDevelopment(),
				Head: []htmx.Node{
					htmx.Script(
						htmx.Src("https://cdn.jsdelivr.net/npm/sortablejs@latest/Sortable.min.js"),
					),
				},
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
						// Todo: Move this script to a separate file
						htmx.Script(htmx.Raw(`
						htmx.onLoad(function(content) {
    var sortables = content.querySelectorAll(".sortable");
    for (var i = 0; i < sortables.length; i++) {
      var sortable = sortables[i];
      var sortableInstance = new Sortable(sortable, {
          animation: 150,
          ghostClass: 'blue-background-class',

          filter: ".htmx-indicator",
          onMove: function (evt) {
            return evt.related.className.indexOf('htmx-indicator') === -1;
          },

          onEnd: function (evt) {
            this.option("disabled", true);
          }
      });

      sortable.addEventListener("htmx:afterSwap", function() {
        sortableInstance.option("disabled", false);
      });
    }
})
						`)),
						cards.Body(
							cards.BodyProps{},
							cards.Actions(
								cards.ActionsProps{},
								workflows.NewStepModal(
									workflows.NewStepModalProps{
										Workflow: p.workflow,
									},
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.OnClick("new_step_modal.showModal()"),
									htmx.Text("Add Step"),
								),
							),
							htmx.FormElement(
								htmx.ClassNames{
									"sortable": true,
								},
								htmx.HxPut(fmt.Sprintf(utils.UpdateWorkflowStepUrlFormat, p.workflow.ID)),
								htmx.HxTrigger("end"),
								htmx.HxSwap("none"),
								htmx.ID("steps"),
								loading.Spinner(
									loading.SpinnerProps{
										ClassNames: htmx.ClassNames{
											"htmx-indicator": true,
										},
									},
								),
								htmx.Group(
									htmx.ForEach(p.workflow.GetStates(), func(state models.WorkflowState, idx int) htmx.Node {
										return workflows.WorkflowStep(
											workflows.WorkflowStepProps{
												ClassNames: htmx.ClassNames{
													tailwind.My2: true,
												},
												State:      state,
												WorkflowID: p.workflow.ID,
											},
										)
									})...,
								),
							),
						),
					),
				)
			},
		),
	)
}
