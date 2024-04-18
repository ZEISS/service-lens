package teams

import (
	"fmt"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// TeamListControllerParams ...
type TeamListControllerParams struct {
	ID uuid.UUID `json:"id" xml:"id" form:"id"`
}

// NewDefaultTeamListControllerParams ...
func NewDefaultTeamListControllerParams() *TeamListControllerParams {
	return &TeamListControllerParams{}
}

// TeamListControllerQuery ...
type TeamListControllerQuery struct {
	Limit  int `json:"limit" xml:"limit" form:"limit"`
	Offset int `json:"offset" xml:"offset" form:"offset"`
}

// NewDefaultTeamListControllerQuery ...
func NewDefaultTeamListControllerQuery() *TeamListControllerQuery {
	return &TeamListControllerQuery{
		Limit:  10,
		Offset: 0,
	}
}

// TeamListController ...
type TeamListController struct {
	db     ports.Repository
	params *TeamListControllerParams
	query  *TeamListControllerQuery
	teams  *models.Pagination[*authz.Team]

	htmx.UnimplementedController
}

// NewTeamListController ...
func NewTeamListController(db ports.Repository) *TeamListController {
	return &TeamListController{
		db: db,
	}
}

// Prepare ...
func (t *TeamListController) Prepare() error {
	if err := t.BindValues(utils.User(t.db), utils.Team(t.db)); err != nil {
		return err
	}

	params := NewDefaultTeamListControllerParams()
	if err := t.BindParams(params); err != nil {
		return err
	}
	t.params = params

	query := NewDefaultTeamListControllerQuery()
	if err := t.BindQuery(query); err != nil {
		return err
	}
	t.query = query

	pagination := models.NewPagination[*authz.Team]()
	if err := t.BindQuery(&pagination); err != nil {
		return err
	}

	teams, err := t.db.ListTeams(t.Context(), pagination)
	if err != nil {
		return err
	}
	t.teams = teams

	return nil
}

// Get ...
func (w *TeamListController) Get() error {
	if w.Hx().HxRequest() {
		w.Hx().ReplaceURL(fmt.Sprintf("%s?limit=%d&offset=%d", w.Ctx().Path(), w.query.Limit, w.query.Offset))

		return w.Hx().RenderComp(
			TeamListTableComponent(
				TeamListTableProps{
					Teams:  w.teams.Rows,
					Offset: w.teams.GetOffset(),
					Limit:  w.teams.GetLimit(),
					Total:  int(w.teams.TotalRows),
				},
			),
		)
	}

	return w.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					User: w.Values(utils.ValuesKeyUser).(*authz.User),
					Team: w.Values(utils.ValuesKeyTeam).(*authz.Team),
				},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						TeamListTableComponent(
							TeamListTableProps{
								Teams:  w.teams.Rows,
								Offset: w.teams.GetOffset(),
								Limit:  w.teams.GetLimit(),
								Total:  int(w.teams.TotalRows),
							},
						),
					),
				),
			),
		),
	)
}

// TeamListTablePaginationProps ...
type TeamListTablePaginationProps struct {
	Limit  int
	Offset int
	Total  int
	Target string
}

// TeamListTablePaginationComponent ...
func TeamListTablePaginationComponent(props TeamListTablePaginationProps, children ...htmx.Node) htmx.Node {
	return tables.Pagination(
		tables.PaginationProps{
			Limit:  props.Limit,
			Offset: props.Offset,
			Target: props.Target,
		},
		tables.Prev(
			tables.PaginationProps{
				URL:    "/site/teams",
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Select(
			tables.SelectProps{
				Limit:  props.Limit,
				Offset: props.Offset,
				Limits: tables.DefaultLimits,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Next(
			tables.PaginationProps{
				URL:    "/site/teams",
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
			},
		),
	)
}

// TeamListTableProps ...
type TeamListTableProps struct {
	Teams  []*authz.Team
	Offset int
	Limit  int
	Total  int
}

// TeamListTableComponent ...
func TeamListTableComponent(props TeamListTableProps, children ...htmx.Node) htmx.Node {
	return tables.Table[*authz.Team](
		tables.TableProps[*authz.Team]{
			ID: "workloads-table",
			Columns: []tables.ColumnDef[*authz.Team]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps[*authz.Team]) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps[*authz.Team], row *authz.Team) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps[*authz.Team]) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps[*authz.Team], row *authz.Team) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf("/site/teams/%s", row.ID.String()),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps[*authz.Team]) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps[*authz.Team], row *authz.Team) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-square": true,
									},
								},

								htmx.HxDelete(fmt.Sprintf("/site/teams/%s", row.ID.String())),
								htmx.HxTarget("closest <tr />"),
								htmx.HxConfirm("Are you sure you want to delete this workload?"),
								icons.TrashOutline(
									icons.IconProps{},
								),
							),
						)
					},
				},
			},
			Rows: tables.NewRows(props.Teams),
			Pagination: TeamListTablePaginationComponent(
				TeamListTablePaginationProps{
					Limit:  props.Limit,
					Offset: props.Offset,
					Total:  props.Total,
					Target: "#workloads-table",
				},
			),
		},
	)
}
