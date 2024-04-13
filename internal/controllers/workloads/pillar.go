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
	ctx    htmx.Ctx

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
	hx := w.Hx()

	ctx, err := htmx.NewDefaultContext(w.Hx().Ctx(), utils.Team(w.Hx().Ctx(), w.db), utils.User(w.Hx().Ctx(), w.db))
	if err != nil {
		return err
	}
	w.ctx = ctx

	lensID, err := uuid.Parse(hx.Context().Params("lens"))
	if err != nil {
		return err
	}

	pillarId, err := hx.Context().ParamsInt("pillar")
	if err != nil {
		return err
	}

	team := htmx.Locals[*authz.Team](w.ctx, utils.ValuesKeyTeam)

	pillar, err := w.db.GetPillarById(hx.Context().Context(), team.Slug, lensID, pillarId)
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
			w.ctx,
			components.PageProps{},
			components.Layout(
				w.ctx,
				components.LayoutProps{},
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
