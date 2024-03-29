package workloads

import (
	"fmt"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
)

// WorkloadNewController ...
type WorkloadNewController struct {
	db   ports.Repository
	team *authz.Team

	htmx.UnimplementedController
}

// NewWorkloadsNewController ...
func NewWorkloadsNewController(db ports.Repository) *WorkloadNewController {
	return &WorkloadNewController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadNewController) Prepare() error {
	team := w.Hx().Values(resolvers.ValuesKeyTeam).(*authz.Team)
	w.team = team

	return nil
}

// Post ...
func (w *WorkloadNewController) Post() error {
	hx := w.Hx()

	profileId, err := uuid.Parse(hx.Ctx().FormValue("profile"))
	if err != nil {
		return err
	}

	lensId, err := uuid.Parse(hx.Ctx().FormValue("lens"))
	if err != nil {
		return err
	}

	workload := &models.Workload{
		Name:        hx.Ctx().FormValue("name"),
		Description: hx.Ctx().FormValue("description"),
		Team:        *w.team,
		ProfileID:   profileId,
		Lenses:      []*models.Lens{{ID: lensId}},
	}

	err = w.db.StoreWorkload(hx.Ctx().Context(), workload)
	if err != nil {
		return err
	}

	hx.Redirect(fmt.Sprintf("/%s/workloads/%s", w.team.Slug, workload.ID))

	return nil
}

// Get ...
func (w *WorkloadNewController) Get() error {
	hx := w.Hx()

	profiles, err := w.db.ListProfiles(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return err
	}

	profilesItems := make([]htmx.Node, len(profiles))
	for i, profile := range profiles {
		profilesItems[i] = htmx.Option(
			htmx.Attribute("value", profile.ID.String()),
			htmx.Text(profile.Name),
		)
	}

	lenses, err := w.db.ListLenses(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return err
	}

	lensesItems := make([]htmx.Node, len(lenses))
	for i, lens := range lenses {
		lensesItems[i] = htmx.Option(
			htmx.Attribute("value", lens.ID.String()),
			htmx.Text(lens.Name),
		)
	}

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				htmx.FormElement(
					htmx.HxPost(""),
					htmx.Label(
						htmx.ClassNames{
							"form-control": true,
							"w-full":       true,
							"max-w-lg":     true,
							"mb-4":         true,
						},
						htmx.Div(
							htmx.ClassNames{
								"label": true,
							},
							htmx.Span(
								htmx.ClassNames{
									"label-text": true,
								},
							),
						),
						htmx.Input(
							htmx.Attribute("type", "text"),
							htmx.Attribute("name", "name"),
							htmx.Attribute("placeholder", "Name ..."),
							htmx.ClassNames{
								"input":          true,
								"input-bordered": true,
								"w-full":         true,
								"max-w-lg":       true,
							},
						),
					),
					htmx.Label(
						htmx.ClassNames{
							"form-control": true,
							"w-full":       true,
							"max-w-lg":     true,
						},
						htmx.Div(
							htmx.ClassNames{
								"label":   true,
								"sr-only": true,
							},
						),
						htmx.Input(
							htmx.Attribute("type", "text"),
							htmx.Attribute("name", "description"),
							htmx.Attribute("placeholder", "Description ..."),
							htmx.ClassNames{
								"input":          true,
								"input-bordered": true,
								"w-full":         true,
								"max-w-lg":       true,
							},
						),
					),
					htmx.Select(
						htmx.ClassNames{
							"select":   true,
							"max-w-xs": true,
							"block":    true,
						},
						htmx.Attribute("name", "profile"),
						htmx.Group(profilesItems...),
					),
					htmx.Select(
						htmx.ClassNames{
							"select":   true,
							"max-w-xs": true,
							"block":    true,
						},
						htmx.Attribute("name", "lens"),
						htmx.Group(lensesItems...),
					),
					buttons.Primary(
						buttons.ButtonProps{},
						htmx.Attribute("type", "submit"),
						htmx.Text("Create Workload"),
					),
				),
			),
		),
	)
}
