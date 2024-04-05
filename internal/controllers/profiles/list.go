package profiles

import (
	"fmt"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// ProfileListControllerParams ...
type ProfileListControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultProfileListControllerParams ...
func NewDefaultProfileListControllerParams() *ProfileListControllerParams {
	return &ProfileListControllerParams{}
}

// ProfileListControllerQuery ...
type ProfileListControllerQuery struct {
	Limit  int `json:"limit" xml:"limit" form:"limit"`
	Offset int `json:"offset" xml:"offset" form:"offset"`
}

// NewDefaultProfileListControllerQuery ...
func NewDefaultProfileListControllerQuery() *ProfileListControllerQuery {
	return &ProfileListControllerQuery{
		Limit:  10,
		Offset: 0,
	}
}

// ProfileListController ...
type ProfileListController struct {
	db       ports.Repository
	profiles []*models.Profile
	team     *authz.Team

	params *ProfileListControllerParams
	query  *ProfileListControllerQuery

	htmx.UnimplementedController
}

// NewProfileListController ...
func NewProfileListController(db ports.Repository) *ProfileListController {
	return &ProfileListController{
		db: db,
	}
}

// Prepare ...
func (w *ProfileListController) Prepare() error {
	hx := w.Hx()

	w.team = hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)

	w.params = NewDefaultProfileListControllerParams()
	if err := hx.Ctx().ParamsParser(w.params); err != nil {
		return err
	}

	w.query = NewDefaultProfileListControllerQuery()
	if err := hx.Ctx().QueryParser(w.query); err != nil {
		fmt.Println("error parsing query", err)
		return err
	}

	profiles, err := w.db.ListProfiles(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: w.query.Limit, Offset: w.query.Offset})
	if err != nil {
		return err
	}
	w.profiles = profiles

	return nil
}

// Get ...
func (w *ProfileListController) Get() error {
	hx := w.Hx()

	if hx.IsHxRequest() {
		hx.ReplaceURL(fmt.Sprintf("%s?limit=%d&offset=%d", hx.Ctx().Path(), w.query.Limit, w.query.Offset))

		return hx.RenderComp(
			ProfileListTableComponent(
				ProfileListTableProps{
					Profiles: w.profiles,
					Team:     w.team,
					Offset:   w.query.Offset,
					Limit:    w.query.Limit,
				},
			),
		)
	}

	return hx.RenderComp(
		components.Page(
			hx,
			components.PageProps{},
			components.Layout(
				hx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						ProfileListTableComponent(
							ProfileListTableProps{
								Profiles: w.profiles,
								Team:     w.team,
								Offset:   w.query.Offset,
								Limit:    w.query.Limit,
							},
						),
					),
				),
			),
		),
	)
}

// ProfileListTablePaginationProps ...
type ProfileListTablePaginationProps struct {
	Limit  int
	Offset int
	Total  int
	Target string
	Team   *authz.Team
}

// ProfileListTablePaginationComponent ...
func ProfileListTablePaginationComponent(props ProfileListTablePaginationProps, children ...htmx.Node) htmx.Node {
	return tables.Pagination(
		tables.PaginationProps{
			Limit:  props.Limit,
			Offset: props.Offset,
			Target: props.Target,
		},
		tables.Prev(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/%s/workloads", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
			},
		),
		tables.Select(
			tables.SelectProps{
				Limit:  props.Limit,
				Offset: props.Offset,
				Limits: tables.DefaultLimits,
				Target: props.Target,
			},
		),
		tables.Next(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/%s/workloads", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
			},
		),
	)
}

// ProfileListTableProps ...
type ProfileListTableProps struct {
	Profiles []*models.Profile
	Team     *authz.Team
	Offset   int
	Limit    int
}

// ProfileListTableComponent ...
func ProfileListTableComponent(props ProfileListTableProps, children ...htmx.Node) htmx.Node {
	return tables.Table[*models.Profile](
		tables.TableProps[*models.Profile]{
			ID: "profiles-tables",
			Columns: []tables.ColumnDef[*models.Profile]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps[*models.Profile]) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps[*models.Profile], row *models.Profile) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps[*models.Profile]) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps[*models.Profile], row *models.Profile) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf("/%s/profiles/%s", props.Team.Slug, row.ID.String()),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps[*models.Profile]) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps[*models.Profile], row *models.Profile) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-square": true,
									},
								},

								htmx.HxDelete(fmt.Sprintf("/%s/profiles/%s", props.Team.Slug, row.ID.String())),
								htmx.HxTarget("closest <tr />"),
								htmx.HxConfirm("Are you sure you want to delete this profile?"),
								icons.TrashOutline(
									icons.IconProps{},
								),
							),
						)
					},
				},
			},
			Rows: tables.NewRows(props.Profiles),
			Pagination: ProfileListTablePaginationComponent(
				ProfileListTablePaginationProps{
					Limit:  props.Limit,
					Offset: props.Offset,
					Total:  len(props.Profiles),
					Target: "profiles-tables",
					Team:   props.Team,
				},
			),
		},
	)
}
