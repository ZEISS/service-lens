package workloads

import (
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadPillarController ...
type WorkloadPillarController struct {
	db   ports.Repository
	team *authz.Team
	lens *models.Lens

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

	team := hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)
	w.team = team

	lensID, err := uuid.Parse(hx.Context().Params("lens"))
	if err != nil {
		return err
	}

	lens, err := w.db.GetLensByID(hx.Context().Context(), team.Slug, lensID)
	if err != nil {
		return err
	}
	w.lens = lens

	return nil
}

// Get ...
func (w *WorkloadPillarController) Get() error {
	hx := w.Hx()

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(),
				),
			),
		),
	)
}
