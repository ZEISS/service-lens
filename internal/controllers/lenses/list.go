package lenses

import (
	"fmt"
	"strconv"

	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// LensListControllerParams ...
type LensListControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultLensListControllerParams ...
func NewDefaultLensListControllerParams() *LensListControllerParams {
	return &LensListControllerParams{}
}

// LensListControllerQuery ...
type LensListControllerQuery struct {
	Limit  int `json:"limit" xml:"limit" form:"limit"`
	Offset int `json:"offset" xml:"offset" form:"offset"`
}

// NewDefaultLensListControllerQuery ...
func NewDefaultLensListControllerQuery() *LensListControllerQuery {
	return &LensListControllerQuery{
		Limit:  10,
		Offset: 0,
	}
}

// LensListController ...
type LensListController struct {
	db     ports.Repository
	lenses *models.Pagination[*models.Lens]

	params *LensListControllerParams
	query  *LensListControllerQuery

	htmx.UnimplementedController
}

// NewLensListController ...
func NewLensListController(db ports.Repository) *LensListController {
	return &LensListController{
		db: db,
	}
}

// Prepare ...
func (w *LensListController) Prepare() error {
	hx := w.Hx()

	if err := w.BindValues(utils.User(w.db), utils.Team(w.db)); err != nil {
		return err
	}

	team := htmx.Locals[*authz.Team](w.DefaultCtx(), utils.ValuesKeyTeam)

	w.params = NewDefaultLensListControllerParams()
	if err := hx.Ctx().ParamsParser(w.params); err != nil {
		return err
	}

	w.query = NewDefaultLensListControllerQuery()
	if err := hx.Ctx().QueryParser(w.query); err != nil {
		return err
	}

	pagination := models.NewPagination[*models.Lens]()
	if err := hx.Ctx().QueryParser(&pagination); err != nil {
		return err
	}

	lenses, err := w.db.ListLenses(hx.Context().Context(), team.Slug, pagination)
	if err != nil {
		return err
	}
	w.lenses = lenses

	return nil
}

// Get ...
func (w *LensListController) Get() error {
	team := htmx.Locals[*authz.Team](w.DefaultCtx(), utils.ValuesKeyTeam)

	return w.Hx().RenderComp(
		components.Page(
			w.DefaultCtx(),
			components.PageProps{},
			components.Layout(
				w.DefaultCtx(),
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						LensListTableComponent(
							LensListTableProps{
								Lenses: w.lenses.Rows,
								Team:   team,
								Offset: w.lenses.GetOffset(),
								Limit:  w.lenses.GetLimit(),
								Total:  int(w.lenses.TotalRows),
							},
						),
					),
				),
			),
		),
	)
}

// LensListTablePaginationProps ...
type LensListTablePaginationProps struct {
	Limit  int
	Offset int
	Total  int
	Target string
	Team   *authz.Team
}

// LensListTablePaginationComponent ...
func LensListTablePaginationComponent(props LensListTablePaginationProps, children ...htmx.Node) htmx.Node {
	return tables.Pagination(
		tables.PaginationProps{
			URL:    fmt.Sprintf("/teams/%s/lenses/list", props.Team.Slug),
			Limit:  props.Limit,
			Offset: props.Offset,
			Target: props.Target,
			Total:  props.Total,
		},
		tables.Prev(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/teams/%s/lenses/list", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Select(
			tables.SelectProps{
				URL:    fmt.Sprintf("/teams%s/lenses/list", props.Team.Slug),
				Limit:  props.Limit,
				Offset: props.Offset,
				Limits: tables.DefaultLimits,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Next(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/teams/%s/lenses/list", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
			},
		),
	)
}

// LensListTableProps ...
type LensListTableProps struct {
	Lenses []*models.Lens
	Team   *authz.Team
	Offset int
	Limit  int
	Total  int
}

// LensListTableComponent ...
func LensListTableComponent(props LensListTableProps, children ...htmx.Node) htmx.Node {
	return tables.Table[*models.Lens](
		tables.TableProps[*models.Lens]{
			ID: "lenses-table",
			Columns: []tables.ColumnDef[*models.Lens]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps[*models.Lens]) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps[*models.Lens], row *models.Lens) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps[*models.Lens]) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps[*models.Lens], row *models.Lens) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf("/%s/lenses/%s", props.Team.Slug, row.ID.String()),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					ID:          "version",
					AccessorKey: "version",
					Header: func(p tables.TableProps[*models.Lens]) htmx.Node {
						return htmx.Th(htmx.Text("Version"))
					},
					Cell: func(p tables.TableProps[*models.Lens], row *models.Lens) htmx.Node {
						return htmx.Td(
							htmx.Text(strconv.Itoa(row.Version)),
						)
					},
				},
				{
					Header: func(p tables.TableProps[*models.Lens]) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps[*models.Lens], row *models.Lens) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-square": true,
									},
								},

								htmx.HxDelete(fmt.Sprintf("/teams/%s/lenses/%s", props.Team.Slug, row.ID.String())),
								htmx.HxTarget("closest <tr />"),
								htmx.HxConfirm("Are you sure you want to delete this lens?"),
								icons.TrashOutline(
									icons.IconProps{},
								),
							),
						)
					},
				},
			},
			Rows: tables.NewRows(props.Lenses),
			Pagination: LensListTablePaginationComponent(
				LensListTablePaginationProps{
					Limit:  props.Limit,
					Offset: props.Offset,
					Total:  props.Total,
					Target: "lenses-table",
					Team:   props.Team,
				},
			),
		},
	)
}
