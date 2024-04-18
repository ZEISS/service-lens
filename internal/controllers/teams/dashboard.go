package teams

import (
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// TeamDashboardController ...
type TeamDashboardController struct {
	db                  ports.Repository
	totalCountWorkloads int
	totalCountLenses    int
	totalCountProfiles  int

	htmx.UnimplementedController
}

// NewTeamDashboardController ...
func NewTeamDashboardController(db ports.Repository) *TeamDashboardController {
	return &TeamDashboardController{
		db: db,
	}
}

// Prepare ...
func (t *TeamDashboardController) Prepare() error {
	if err := t.BindValues(utils.User(t.db), utils.Team(t.db)); err != nil {
		return err
	}

	team := t.Values(utils.ValuesKeyTeam).(*authz.Team)

	totalCountWorkloads, err := t.db.TotalCountWorkloads(t.Context(), team.Slug)
	if err != nil {
		return err
	}
	t.totalCountWorkloads = totalCountWorkloads

	totalCountLenses, err := t.db.TotalCountLenses(t.Context(), team.Slug)
	if err != nil {
		return err
	}
	t.totalCountLenses = totalCountLenses

	totalCountProfiles, err := t.db.TotalCountProfiles(t.Context(), team.Slug)
	if err != nil {
		return err
	}
	t.totalCountProfiles = totalCountProfiles

	return nil
}

// Error ...
func (t *TeamDashboardController) Error(err error) error {
	return err
}

// Get ...
func (t *TeamDashboardController) Get() error {
	return t.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: t.Values(utils.ValuesKeyUser).(*authz.User),
					Team: t.Values(utils.ValuesKeyTeam).(*authz.Team),
				},
				components.Wrap(
					components.WrapProps{},
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Overview"),
							),
							stats.Stats(
								stats.StatsProps{},
								stats.Stat(
									stats.StatProps{},
									stats.Figure(
										stats.FigureProps{},
										icons.BriefcaseOutline(
											icons.IconProps{},
										),
									),
									stats.Title(
										stats.TitleProps{},
										htmx.Text("Total Workloads"),
									),
									stats.Value(
										stats.ValueProps{},
										htmx.Text(utils.IntStr(t.totalCountWorkloads)),
									),
									stats.Description(
										stats.DescriptionProps{},
										htmx.Text("Total number of workloads in this team"),
									),
								),
								stats.Stat(
									stats.StatProps{},
									stats.Figure(
										stats.FigureProps{},
										icons.MagnifyingGlassOutline(
											icons.IconProps{},
										),
									),
									stats.Title(
										stats.TitleProps{},
										htmx.Text("Total Lenses"),
									),
									stats.Value(
										stats.ValueProps{},
										htmx.Text(utils.IntStr(t.totalCountLenses)),
									),
									stats.Description(
										stats.DescriptionProps{},
										htmx.Text("Total number of lenses in this team"),
									),
								),
								stats.Stat(
									stats.StatProps{},
									stats.Title(
										stats.TitleProps{},
										htmx.Text("Total Profiles"),
									),
									stats.Value(
										stats.ValueProps{},
										htmx.Text(utils.IntStr(t.totalCountProfiles)),
									),
									stats.Description(
										stats.DescriptionProps{},
										htmx.Text("Total number of profiles in this team"),
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
