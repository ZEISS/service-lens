package teams

import (
	"fmt"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
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
	teams  []*authz.Team

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
	hx := t.Hx()

	params := NewDefaultTeamListControllerParams()
	if err := hx.Context().ParamsParser(params); err != nil {
		return err
	}
	t.params = params

	query := NewDefaultTeamListControllerQuery()
	if err := hx.Context().QueryParser(query); err != nil {
		return err
	}
	t.query = query

	teams, err := t.db.ListTeams(hx.Context().Context(), &models.Pagination{Limit: query.Limit, Offset: query.Offset})
	if err != nil {
		return err
	}
	t.teams = teams

	return nil
}

// Get ...
func (t *TeamListController) Get() error {
	hx := t.Hx()

	table := tables.Table[*authz.Team](
		tables.TableProps[*authz.Team]{
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
							htmx.Text(row.Name),
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
								htmx.HxConfirm("Are you sure you want to delete this team?"),
								icons.TrashOutline(
									icons.IconProps{},
								),
							),
						)
					},
				},
			},
			Rows: tables.NewRows(t.teams),
		},
		htmx.ID("data-table"),
	)

	// if hx.IsHxRequest() {
	// 	hx.ReplaceURL(fmt.Sprintf("%s?limit=%d,offset=%d", hx.Ctx().Path(), limit, offset))

	// 	return hx.RenderComp(table)
	// }

	return hx.RenderComp(
		components.Layout(
			hx,
			components.LayoutProps{},
			components.Wrap(
				components.WrapProps{},
				htmx.Div(
					htmx.ClassNames{
						"overflow-x-auto": true,
					},
					table,
					htmx.Div(
						htmx.ClassNames{
							"bg-base-100": true,
							"p-4":         true,
						},
						tables.Pagination(
							tables.PaginationProps{
								Limit:  t.query.Limit,
								Offset: t.query.Offset,
							},
							tables.Prev(
								tables.PaginationProps{
									URL:    "/api/data",
									Offset: t.query.Offset,
									Limit:  t.query.Limit,
								},
							),
							tables.Next(
								tables.PaginationProps{
									URL: "/api/data",
								},
							),
						),
					),
				),
			),
		),
	)
}
