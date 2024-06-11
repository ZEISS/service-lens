package workloads

import (
	"context"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/workloads"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
)

// WorkloadShowControllerImpl ...
type WorkloadShowControllerImpl struct {
	workload models.Workload
	store    ports.Datastore
	htmx.DefaultController
}

// NewWorkloadShowController ...
func NewWorkloadShowController(store ports.Datastore) *WorkloadShowControllerImpl {
	return &WorkloadShowControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *WorkloadShowControllerImpl) Prepare() error {
	err := w.BindParams(&w.workload)
	if err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetWorkload(ctx, &w.workload)
	})
}

// Get ...
func (w *WorkloadShowControllerImpl) Get() error {
	return w.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: w.Path(),
				},
				components.Wrap(
					components.WrapProps{
						ClassNames: htmx.ClassNames{
							"py-4": true,
						},
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
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "ID",
										Data:  w.workload.ID.String(),
									},
								),
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "Name",
										Data:  w.workload.Name,
									},
								),
								components.CardDataBlock(
									&components.CardDataBlockProps{
										Title: "Description",
										Data:  w.workload.Description,
									},
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
											w.workload.CreatedAt.Format("2006-01-02 15:04:05"),
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
											w.workload.UpdatedAt.Format("2006-01-02 15:04:05"),
										),
									),
								),
							),
						),
					),
				),
				components.Wrap(
					components.WrapProps{
						ClassNames: htmx.ClassNames{
							"py-4": true,
						},
					},
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Lenses"),
							),
							workloads.LensesTable(
								workloads.LensesTableProps{
									Workload: &w.workload,
								},
							),
							// htmx.Div(
							// 	htmx.ClassNames{
							// 		"overflow-x-auto": true,
							// 	},

							// 	htmx.Table(
							// 		htmx.ClassNames{
							// 			"table": true,
							// 		},
							// 		htmx.THead(
							// 			htmx.Tr(
							// 				htmx.Th(htmx.Text("ID")),
							// 				htmx.Th(htmx.Text("Lens")),
							// 			),
							// 		),
							// 		// htmx.TBody(
							// 		// 	htmx.Group(lenses...),
							// 		// ),
							// 	),
							// ),
						),
					),
				),
				components.Wrap(
					components.WrapProps{
						ClassNames: htmx.ClassNames{
							"py-4": true,
						},
					},
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Profile"),
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
									htmx.Text(
										w.workload.Profile.Name,
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
									htmx.Text("Description"),
								),
								htmx.H3(
									htmx.Text(
										w.workload.Profile.Description,
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
										w.workload.Profile.UpdatedAt.Format("2006-01-02 15:04:05"),
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
