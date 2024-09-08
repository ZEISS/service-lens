package workloads

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/builders"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/components/workloads"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadLensController ...
type WorkloadLensController struct {
	workload models.Workload
	lens     models.Lens
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	risks    *builders.RiskAnalyzerBuilder
	htmx.DefaultController
}

// NewWorkloadLensController ...
func NewWorkloadLensController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadLensController {
	return &WorkloadLensController{
		risks: builders.NewRiskAnalyzerBuilder(),
		store: store,
	}
}

// Prepare ...
func (w *WorkloadLensController) Prepare() error {
	err := w.BindParams(&w.workload)
	if err != nil {
		return err
	}

	err = w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		if err := tx.GetWorkload(ctx, &w.workload); err != nil {
			return err
		}

		id, err := uuid.Parse(w.Ctx().Params("lens"))
		if err != nil {
			return err
		}
		w.lens.ID = id

		return tx.GetLens(ctx, &w.lens)
	})
	if err != nil {
		return err
	}

	answers := make([]models.WorkloadLensQuestionAnswer, 0)

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListLensAnswers(ctx, w.lens.ID, &answers)
	})
}

// Get ...
func (w *WorkloadLensController) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       w.lens.Name,
				Path:        w.Path(),
				User:        w.Session().User,
				Development: w.IsDevelopment(),
			},
			func() htmx.Node {
				totalHighRisk := 0
				totalMediumRisk := 0

				for _, a := range w.workload.Answers {
					if a.Risk != nil {
						switch a.Risk.Risk {
						case "HIGH_RISK":
							totalHighRisk++
						case "MEDIUM_RISK":
							totalMediumRisk++
						}
					}
				}

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
								htmx.H1(
									htmx.Text(w.lens.Name),
								),
								htmx.P(
									htmx.Text(w.lens.Description),
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
											w.lens.CreatedAt.Format("2006-01-02 15:04:05"),
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
											w.lens.UpdatedAt.Format("2006-01-02 15:04:05"),
										),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								links.Button(
									links.LinkProps{
										Href: fmt.Sprintf("/workloads/%s/lenses/%s/edit", w.workload.ID, w.lens.ID),
									},
									htmx.Text("Start Review"),
								),
							),
						),
					),
					stats.Stats(
						stats.StatsProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2:     true,
								tailwind.Shadow: false,
							},
						},
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Overall Questions Answered"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text(fmt.Sprintf("%d", len(w.workload.Answers))),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text(fmt.Sprintf("Total of %d questions", w.lens.TotalQuestions())),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Overall High Risks"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text(fmt.Sprintf("%d", totalHighRisk)),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text(fmt.Sprintf("Total of %d questions", w.lens.TotalQuestions())),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Overall Medium Risks"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text(fmt.Sprintf("%d", totalMediumRisk)),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text(fmt.Sprintf("Total of %d questions", w.lens.TotalQuestions())),
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
								htmx.Text("Pillars"),
							),
							workloads.LensPillarTable(
								workloads.LensPillarTableProps{
									Lens: &w.lens,
								},
							),
						),
					),
				)
			},
		),
	)
}
