package workloads

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/builders"
	"github.com/zeiss/service-lens/internal/models"
)

// WorkloadsRisksCardProps ...
type WorkloadsRisksCardProps struct {
	// Workload ...
	Workload models.Workload
	// Lens ...
	Lens models.Lens
	// Risks ...
	Risks *builders.RiskAnalyzerBuilder
}

// WorkloadsRisksCard ...
func WorkloadsRisksCard(props WorkloadsRisksCardProps) htmx.Node {
	return stats.Stats(
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
				htmx.Text(fmt.Sprintf("%d", props.Workload.TotalAnswers())),
			),
			stats.Description(
				stats.DescriptionProps{},
				htmx.Text(fmt.Sprintf("Total of %d questions", props.Workload.TotalQuestions())),
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
				htmx.Text(fmt.Sprintf("%d", props.Workload.TotalHighRisks())),
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
				htmx.Text(fmt.Sprintf("%d", props.Workload.TotalMediumRisks())),
			),
		),
		// stats.Stat(
		// 	stats.StatProps{},
		// 	stats.Title(
		// 		stats.TitleProps{},
		// 		htmx.Text("Total Medium Risks"),
		// 	),
		// 	stats.Value(
		// 		stats.ValueProps{},
		// 		htmx.Text(conv.String(props.Risks.TotalMediumRisks())),
		// 	),
		// ),
		// stats.Stat(
		// 	stats.StatProps{},
		// 	stats.Title(
		// 		stats.TitleProps{},
		// 		htmx.Text("Total Low Risks"),
		// 	),
		// 	stats.Value(
		// 		stats.ValueProps{},
		// 		htmx.Text(conv.String(props.Risks.TotalLowRisks())),
		// 	),
		// ),
		// stats.Stat(
		// 	stats.StatProps{},
		// 	stats.Title(
		// 		stats.TitleProps{},
		// 		htmx.Text("Total No Risks"),
		// 	),
		// 	stats.Value(
		// 		stats.ValueProps{},
		// 		htmx.Text(conv.String(props.Risks.TotalNoRisk())),
		// 	),
		// ),
		// stats.Stat(
		// 	stats.StatProps{},
		// 	stats.Title(
		// 		stats.TitleProps{},
		// 		htmx.Text("Total Unanswered Risks"),
		// 	),
		// 	stats.Value(
		// 		stats.ValueProps{},
		// 		htmx.Text(conv.String(props.Risks.TotalNotAnswered())),
		// 	),
		// ),
		// stats.Stat(
		// 	stats.StatProps{},
		// 	stats.Title(
		// 		stats.TitleProps{},
		// 		htmx.Text("Total Not Applicable Risks"),
		// 	),
		// 	stats.Value(
		// 		stats.ValueProps{},
		// 		htmx.Text(conv.String(props.Risks.TotalNotApplicable())),
		// 	),
		// ),
	)
}
