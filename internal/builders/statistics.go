package builders

import (
	"fmt"

	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/service-lens/internal/models"

	"github.com/expr-lang/expr"
)

// Stats ...
type Stats struct {
	// High ...
	High bool
	// Medium ...
	Medium bool
	// Low ...
	Low bool
	// NotApplicable ...
	NotApplicable bool
	// NotAnswered ...
	NotAnswered bool
}

// StatKind ...
type StatKind int

// StatKindNames ...
var StatKindNames = map[StatKind]string{}

// NewStatKind ...
func NewStatKind(name string) StatKind {
	return StatKind(len(StatKindNames))
}

// QuestionStatsKind ...
var QuestionStatsKind = NewStatKind("question")

// QuestionStats ...
type QuestionStats struct{}

// Kind ...
func (q *QuestionStats) Kind() StatKind {
	return QuestionStatsKind
}

// StatisticsBuilder ...
type StatisticsBuilder struct {
	lens    models.Lens
	answers []models.WorkloadLensQuestionAnswer

	cache map[StatKind]Stats
}

// StatisticsBuilderOpt ...
type StatisticsBuilderOpt func(*StatisticsBuilder)

// NewStatisticsBuilder ...
func NewStatisticsBuilder(lens models.Lens, answers []models.WorkloadLensQuestionAnswer, opts ...StatisticsBuilderOpt) *StatisticsBuilder {
	s := &StatisticsBuilder{lens, answers, map[StatKind]Stats{}}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Build ...
func (s *StatisticsBuilder) Build() error {
	idxAnswers := make(map[int]models.WorkloadLensQuestionAnswer)

	for _, a := range s.answers {
		idxAnswers[a.Question.ID] = a
	}

	for _, q := range s.lens.Pillars {
		for _, q := range q.Questions {
			stats := Stats{}

			env := map[string]bool{
				"default": true,
			}

			for _, c := range q.Choices {
				env[string(c.Ref)] = false
			}

			a, ok := idxAnswers[q.ID]
			if !ok {
				break
			}

			for _, c := range a.Choices {
				fmt.Println(c.Ref)
				env[cast.String(c.Ref)] = true
			}

			for _, r := range q.Risks {
				rule := r.Condition

				fmt.Println(rule)

				program, err := expr.Compile(rule, expr.Env(env))
				if err != nil {
					return err
				}

				output, err := expr.Run(program, env)
				if err != nil {
					return err
				}
				v, ok := output.(bool)
				if !ok {
					return fmt.Errorf("expected bool, got %T", v)
				}

				switch r.Risk {
				case "HIGH_RISK":
					stats.High = v
				case "MEDIUM_RISK":
					stats.Medium = v
				case "LOW_RISK":
					stats.Low = v
				}
			}

			fmt.Println(stats)
		}
	}

	fmt.Println(s.cache)

	return nil
}
