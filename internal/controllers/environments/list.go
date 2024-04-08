package environments

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

// EnvironmentListControllerParams ...
type EnvironmentListControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultEnvironmentListControllerParams ...
func NewDefaultEnvironmentListControllerParams() *EnvironmentListControllerParams {
	return &EnvironmentListControllerParams{}
}

// EnvironmentListControllerQuery ...
type EnvironmentListControllerQuery struct {
	Limit  int `json:"limit" xml:"limit" form:"limit"`
	Offset int `json:"offset" xml:"offset" form:"offset"`
}

// NewDefaultEnvironmentListControllerQuery ...
func NewDefaultEnvironmentListControllerQuery() *EnvironmentListControllerQuery {
	return &EnvironmentListControllerQuery{
		Limit:  10,
		Offset: 0,
	}
}

// EnvironmentListController ...
type EnvironmentListController struct {
	db           ports.Repository
	Environments *models.Pagination[*models.Environment]
	team         *authz.Team

	params *EnvironmentListControllerParams
	query  *EnvironmentListControllerQuery

	htmx.UnimplementedController
}

// NewEnvironmentListController ...
func NewEnvironmentListController(db ports.Repository) *EnvironmentListController {
	return &EnvironmentListController{
		db: db,
	}
}

// Prepare ...
func (w *EnvironmentListController) Prepare() error {
	hx := w.Hx()

	w.team = hx.Values(resolvers.ValuesKeyTeam).(*authz.Team)

	w.params = NewDefaultEnvironmentListControllerParams()
	if err := hx.Ctx().ParamsParser(w.params); err != nil {
		return err
	}

	w.query = NewDefaultEnvironmentListControllerQuery()
	if err := hx.Ctx().QueryParser(w.query); err != nil {
		return err
	}

	pagination := models.NewPagination[*models.Environment]()
	if err := hx.Ctx().QueryParser(&pagination); err != nil {
		return err
	}

	Environments, err := w.db.ListEnvironment(hx.Context().Context(), w.team.Slug, pagination)
	if err != nil {
		return err
	}
	w.Environments = Environments

	return nil
}

// Get ...
func (w *EnvironmentListController) Get() error {
	fmt.Println(w.Environments.TotalRows)
	return w.Hx().RenderComp(
		components.Page(
			w.Hx(),
			components.PageProps{},
			components.Layout(
				w.Hx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						EnvironmentListTableComponent(
							EnvironmentListTableProps{
								Environments: w.Environments.Rows,
								Team:         w.team,
								Offset:       w.query.Offset,
								Limit:        w.query.Limit,
								Total:        int(w.Environments.TotalRows),
							},
						),
					),
				),
			),
		),
	)
}

// EnvironmentListTablePaginationProps ...
type EnvironmentListTablePaginationProps struct {
	Limit  int
	Offset int
	Total  int
	Target string
	Team   *authz.Team
}

// EnvironmentListTablePaginationComponent ...
func EnvironmentListTablePaginationComponent(props EnvironmentListTablePaginationProps, children ...htmx.Node) htmx.Node {
	return tables.Pagination(
		tables.PaginationProps{
			URL:    fmt.Sprintf("/%s/Environments/list", props.Team.Slug),
			Limit:  props.Limit,
			Offset: props.Offset,
			Target: props.Target,
			Total:  props.Total,
		},
		tables.Prev(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/%s/Environments/list", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Select(
			tables.SelectProps{
				URL:    fmt.Sprintf("/%s/Environments/list", props.Team.Slug),
				Limit:  props.Limit,
				Offset: props.Offset,
				Limits: tables.DefaultLimits,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Next(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/%s/Environments/list", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
			},
		),
	)
}

// EnvironmentListTableProps ...
type EnvironmentListTableProps struct {
	Environments []*models.Environment
	Team         *authz.Team
	Offset       int
	Limit        int
	Total        int
}

// EnvironmentListTableComponent ...
func EnvironmentListTableComponent(props EnvironmentListTableProps, children ...htmx.Node) htmx.Node {
	return tables.Table[*models.Environment](
		tables.TableProps[*models.Environment]{
			ID: "Environments-tables",
			Columns: []tables.ColumnDef[*models.Environment]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps[*models.Environment]) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps[*models.Environment], row *models.Environment) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps[*models.Environment]) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps[*models.Environment], row *models.Environment) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf("/%s/Environments/%s", props.Team.Slug, row.ID.String()),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps[*models.Environment]) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps[*models.Environment], row *models.Environment) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-square": true,
									},
								},

								htmx.HxDelete(fmt.Sprintf("/%s/Environments/%s", props.Team.Slug, row.ID.String())),
								htmx.HxTarget("closest <tr />"),
								htmx.HxConfirm("Are you sure you want to delete this Environment?"),
								icons.TrashOutline(
									icons.IconProps{},
								),
							),
						)
					},
				},
			},
			Rows: tables.NewRows(props.Environments),
			Pagination: EnvironmentListTablePaginationComponent(
				EnvironmentListTablePaginationProps{
					Limit:  props.Limit,
					Offset: props.Offset,
					Total:  props.Total,
					Target: "Environments-tables",
					Team:   props.Team,
				},
			),
		},
	)
}
