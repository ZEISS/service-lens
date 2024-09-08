package workloads

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/service-lens/internal/models"
)

// WorkloadsRisksCardProps ...
type WorkloadsRisksCardProps struct {
	// Workload ...
	Workload models.Workload
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
			stats.Description(
				stats.DescriptionProps{},
				htmx.Text(fmt.Sprintf("Total of %d lenses", props.Workload.TotalLenses())),
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
			stats.Description(
				stats.DescriptionProps{},
				htmx.Text(fmt.Sprintf("Total of %d lenses", props.Workload.TotalLenses())),
			),
		),
	)
}
