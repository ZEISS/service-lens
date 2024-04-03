package workloads

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
	"github.com/zeiss/fiber-htmx/components/tables"
)

// WorkloadListControllerParams ...
type WorkloadListControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultWorkloadListControllerParams ...
func NewDefaultWorkloadListControllerParams() *WorkloadListControllerParams {
	return &WorkloadListControllerParams{}
}

// WorkloadListControllerQuery ...
type WorkloadListControllerQuery struct {
	Limit  int `json:"limit" xml:"limit" form:"limit"`
	Offset int `json:"offset" xml:"offset" form:"offset"`
}

// NewDefaultWorkloadListControllerQuery ...
func NewDefaultWorkloadListControllerQuery() *WorkloadListControllerQuery {
	return &WorkloadListControllerQuery{
		Limit:  10,
		Offset: 0,
	}
}

// WorkloadListController ...
type WorkloadListController struct {
	db        ports.Repository
	workloads []*models.Workload
	team      *authz.Team

	params *WorkloadListControllerParams
	query  *WorkloadListControllerQuery

	htmx.UnimplementedController
}

// NewWorkloadListController ...
func NewWorkloadListController(db ports.Repository) *WorkloadListController {
	return &WorkloadListController{
		db: db,
	}
}

// Prepare ...
func (w *WorkloadListController) Prepare() error {
	hx := w.Hx()

	w.team = hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)

	w.params = NewDefaultWorkloadListControllerParams()
	if err := hx.Ctx().ParamsParser(w.params); err != nil {
		return err
	}

	w.query = NewDefaultWorkloadListControllerQuery()
	if err := hx.Ctx().QueryParser(w.query); err != nil {
		fmt.Println("error parsing query", err)
		return err
	}

	workloads, err := w.db.ListWorkloads(hx.Context().Context(), w.team.Slug, &models.Pagination{Limit: w.query.Limit, Offset: w.query.Offset})
	if err != nil {
		return err
	}
	w.workloads = workloads

	return nil
}

// Get ...
func (w *WorkloadListController) Get() error {
	hx := w.Hx()

	if hx.IsHxRequest() {
		hx.ReplaceURL(fmt.Sprintf("%s?limit=%d&offset=%d", hx.Ctx().Path(), w.query.Limit, w.query.Offset))

		return hx.RenderComp(
			WorkloadListTableComponent(
				WorkloadListTableProps{
					Workloads: w.workloads,
					Team:      w.team,
					Offset:    w.query.Offset,
					Limit:     w.query.Limit,
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
						WorkloadListTableComponent(
							WorkloadListTableProps{
								Workloads: w.workloads,
								Team:      w.team,
								Offset:    w.query.Offset,
								Limit:     w.query.Limit,
							},
						),
					),
				),
			),
		),
	)
}

// WorkloadListTablePaginationProps ...
type WorkloadListTablePaginationProps struct {
	Limit  int
	Offset int
	Total  int
	Target string
	Team   *authz.Team
}

// WorkloadListTablePaginationComponent ...
func WorkloadListTablePaginationComponent(props WorkloadListTablePaginationProps, children ...htmx.Node) htmx.Node {
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

// WorkloadListTableProps ...
type WorkloadListTableProps struct {
	Workloads []*models.Workload
	Team      *authz.Team
	Offset    int
	Limit     int
}

// WorkloadListTableComponent ...
func WorkloadListTableComponent(props WorkloadListTableProps, children ...htmx.Node) htmx.Node {
	return tables.Table[*models.Workload](
		tables.TableProps[*models.Workload]{
			ID: "workloads-table",
			Columns: []tables.ColumnDef[*models.Workload]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps[*models.Workload]) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps[*models.Workload], row *models.Workload) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps[*models.Workload]) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps[*models.Workload], row *models.Workload) htmx.Node {
						return htmx.Td(
							htmx.Text(row.Name),
						)
					},
				},
				{
					Header: func(p tables.TableProps[*models.Workload]) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps[*models.Workload], row *models.Workload) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-square": true,
									},
								},

								htmx.HxDelete(fmt.Sprintf("/%s/workloads/%s", props.Team.Slug, row.ID.String())),
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
			Rows: tables.NewRows(props.Workloads),
			Pagination: WorkloadListTablePaginationComponent(
				WorkloadListTablePaginationProps{
					Limit:  props.Limit,
					Offset: props.Offset,
					Total:  len(props.Workloads),
					Target: "#workloads-table",
					Team:   props.Team,
				},
			),
		},
	)
}
