package profiles

import (
	"fmt"

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
	profiles *models.Pagination[*models.Profile]
	ctx      htmx.Ctx

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

	ctx, err := htmx.NewDefaultContext(w.Hx().Ctx(), utils.Team(w.Hx().Ctx(), w.db), utils.User(w.Hx().Ctx(), w.db))
	if err != nil {
		return err
	}
	w.ctx = ctx

	w.params = NewDefaultProfileListControllerParams()
	if err := hx.Ctx().ParamsParser(w.params); err != nil {
		return err
	}

	w.query = NewDefaultProfileListControllerQuery()
	if err := hx.Ctx().QueryParser(w.query); err != nil {
		return err
	}

	pagination := models.NewPagination[*models.Profile]()
	if err := hx.Ctx().QueryParser(&pagination); err != nil {
		return err
	}

	team := htmx.Locals[*authz.Team](w.ctx, utils.ValuesKeyTeam)

	profiles, err := w.db.ListProfiles(hx.Context().Context(), team.Slug, pagination)
	if err != nil {
		return err
	}
	w.profiles = profiles

	return nil
}

// Get ...
func (w *ProfileListController) Get() error {
	team := htmx.Locals[*authz.Team](w.ctx, utils.ValuesKeyTeam)

	return w.Hx().RenderComp(
		components.Page(
			w.ctx,
			components.PageProps{},
			components.Layout(
				w.ctx,
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
					htmx.Div(
						htmx.ClassNames{
							"overflow-x-auto": true,
						},
						ProfileListTableComponent(
							ProfileListTableProps{
								Profiles: w.profiles.Rows,
								Team:     team,
								Offset:   w.query.Offset,
								Limit:    w.query.Limit,
								Total:    int(w.profiles.TotalRows),
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
			URL:    fmt.Sprintf("/%s/profiles/list", props.Team.Slug),
			Limit:  props.Limit,
			Offset: props.Offset,
			Target: props.Target,
			Total:  props.Total,
		},
		tables.Prev(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/%s/profiles/list", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Select(
			tables.SelectProps{
				URL:    fmt.Sprintf("/%s/profiles/list", props.Team.Slug),
				Limit:  props.Limit,
				Offset: props.Offset,
				Limits: tables.DefaultLimits,
				Target: props.Target,
				Total:  props.Total,
			},
		),
		tables.Next(
			tables.PaginationProps{
				URL:    fmt.Sprintf("/%s/profiles/list", props.Team.Slug),
				Offset: props.Offset,
				Limit:  props.Limit,
				Target: props.Target,
				Total:  props.Total,
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
	Total    int
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
					Total:  props.Total,
					Target: "profiles-tables",
					Team:   props.Team,
				},
			),
		},
	)
}
