package workloads

import (
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
)

// WorkloadPillarController ...
type WorkloadPillarController struct {
	db     ports.Repository
	pillar *models.Pillar

	htmx.UnimplementedController
}

// NewWorkloadLensController ...
func NewWorkloadPillarController(db ports.Repository) *WorkloadPillarController {
	return &WorkloadPillarController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadPillarController) Prepare() error {
	if err := w.BindValues(utils.User(w.db), utils.Team(w.db)); err != nil {
		return err
	}

	lensID, err := uuid.Parse(w.Ctx().Params("lens"))
	if err != nil {
		return err
	}

	pillarId, err := w.Ctx().ParamsInt("pillar")
	if err != nil {
		return err
	}

	team := w.Values(utils.ValuesKeyTeam).(*authz.Team)

	pillar, err := w.db.GetPillarById(w.Context(), team.Slug, lensID, pillarId)
	if err != nil {
		return err
	}
	w.pillar = pillar

	return nil
}

// Get ...
func (w *WorkloadPillarController) Get() error {
	hx := w.Hx()

	questions := make([]htmx.Node, len(w.pillar.Questions))
	for i, question := range w.pillar.Questions {
		questions[i] = cards.Card(
			cards.CardProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
			},
			cards.Body(
				cards.BodyProps{},
				cards.Title(
					cards.TitleProps{},
					htmx.Text(question.Title),
				),
				htmx.Text(question.Description),
			),
		)
	}

	return hx.RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: w.Values(utils.ValuesKeyUser).(*authz.User),
					Team: w.Values(utils.ValuesKeyTeam).(*authz.Team),
				},
				components.Wrap(
					components.WrapProps{},
					htmx.H1(
						htmx.Text(w.pillar.Name),
					),
				),
				components.Wrap(
					components.WrapProps{},
					htmx.Group(questions...),
				),
			),
		),
	)
}
