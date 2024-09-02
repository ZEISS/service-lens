package builders

import (
	"fmt"
	"sync"

	"github.com/expr-lang/expr"
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/service-lens/internal/models"
)

// Risk ...
type Risk struct {
	// High ...
	High bool
	// Medium ...
	Medium bool
	// Low ...
	Low bool
	// NoRisk ...
	NoRisk bool
	// NotApplicable ...
	NotApplicable bool
	// NotAnswered ...
	NotAnswered bool
}

// DefaultRiskAnalyzerFunc ...
// nolint:gocyclo
var DefaultRiskAnalyzerFunc = func(s *RiskAnalyzerBuilder) error {
	for _, q := range s.Questions {
		s.Risks[q.ID] = &Risk{}

		a, ok := s.Answers[q.ID]
		if !ok {
			s.Risks[q.ID].NotAnswered = true
			continue
		}

		env := map[string]bool{
			"default": true,
		}

		for _, c := range q.Choices {
			env[string(c.Ref)] = false
		}

		for _, c := range a.Choices {
			env[cast.String(c.Ref)] = true
		}

		for _, r := range q.Risks {
			rule := r.Condition

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
				s.Risks[q.ID].High = v
			case "MEDIUM_RISK":
				s.Risks[q.ID].Medium = v
			case "LOW_RISK":
				s.Risks[q.ID].Low = v
			case "NO_RISK":
				s.Risks[q.ID].NoRisk = v
			}

			if v {
				break
			}
		}
	}

	return nil
}

// RiskAnalyzerBuilder ...
type RiskAnalyzerBuilder struct {
	Questions map[int]models.Question
	Answers   map[int]models.WorkloadLensQuestionAnswer
	Risks     map[int]*Risk
	Error     error

	errOnce sync.Once
}

// RiskAnalyzerBuilderOpt ...
type RiskAnalyzerBuilderOpt func(*RiskAnalyzerBuilder)

// NewRiskAnalyzerBuilder ...
func NewRiskAnalyzerBuilder(opts ...RiskAnalyzerBuilderOpt) *RiskAnalyzerBuilder {
	s := &RiskAnalyzerBuilder{}
	s.Questions = make(map[int]models.Question)
	s.Answers = make(map[int]models.WorkloadLensQuestionAnswer)
	s.Risks = make(map[int]*Risk)

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// AddAnswers ...
func (r *RiskAnalyzerBuilder) AddAnswers(answers ...models.WorkloadLensQuestionAnswer) *RiskAnalyzerBuilder {
	for _, a := range answers {
		r.Answers[a.Question.ID] = a
	}

	return r
}

// AddQuestions ...
func (r *RiskAnalyzerBuilder) AddQuestions(questions ...models.Question) *RiskAnalyzerBuilder {
	for _, q := range questions {
		r.Questions[q.ID] = q
	}

	return r
}

// RiskAnalyzerBuildFunc ...
type RiskAnalyzerBuildFunc func(*RiskAnalyzerBuilder) error

// Build ...
func (s *RiskAnalyzerBuilder) Build(fn RiskAnalyzerBuildFunc) *RiskAnalyzerBuilder {
	err := fn(s)
	s.errOnce.Do(func() {
		s.Error = err
	})

	return s
}

// TotalHighRisks ...
func (r *RiskAnalyzerBuilder) TotalHighRisks() int {
	var total int
	for _, risk := range r.Risks {
		if risk.High {
			total++
		}
	}

	return total
}

// TotalMediumRisks ...
func (r *RiskAnalyzerBuilder) TotalMediumRisks() int {
	var total int
	for _, risk := range r.Risks {
		if risk.Medium {
			total++
		}
	}

	return total
}

// TotalLowRisks ...
func (r *RiskAnalyzerBuilder) TotalLowRisks() int {
	var total int
	for _, risk := range r.Risks {
		if risk.Low {
			total++
		}
	}

	return total
}

// TotalNotAnswered ...
func (r *RiskAnalyzerBuilder) TotalNotAnswered() int {
	var total int
	for _, risk := range r.Risks {
		if risk.NotAnswered {
			total++
		}
	}

	return total
}

// TotalNotApplicable ...
func (r *RiskAnalyzerBuilder) TotalNotApplicable() int {
	var total int
	for _, risk := range r.Risks {
		if risk.NotApplicable {
			total++
		}
	}

	return total
}

// TotalNoRisk ...
func (r *RiskAnalyzerBuilder) TotalNoRisk() int {
	var total int
	for _, risk := range r.Risks {
		if risk.NoRisk {
			total++
		}
	}

	return total
}

// TotalQuestions ...
func (r *RiskAnalyzerBuilder) TotalQuestions() int {
	return len(r.Questions)
}

// TotalAnswers ...
func (r *RiskAnalyzerBuilder) TotalAnswers() int {
	return len(r.Answers)
}
