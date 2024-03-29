package lenses

import (
	"fmt"
	"io"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"
)

// LensListController ...
type LensListController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewLensListController ...
func NewLensListController(db ports.Repository) *LensListController {
	return &LensListController{db, htmx.UnimplementedController{}}
}

// Put ...
func (l *LensListController) Put() error {
	hx := l.Hx()

	file, err := hx.Ctx().FormFile("spec")
	if err != nil {
		return err
	}

	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	bb, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	lens := &models.Lens{}
	err = lens.UnmarshalJSON(bb)
	if err != nil {
		return err
	}

	lens.Tags = []*models.Tag{
		{Name: hx.Ctx().FormValue("tag")},
	}

	lens, err = l.db.AddLens(hx.Ctx().Context(), lens)
	if err != nil {
		return err
	}

	hx.Redirect("/lenses/" + lens.ID.String())

	return nil
}

// Post ...
func (l *LensListController) Post() error {
	return l.Hx().RenderComp(
		components.Page(
			l.Hx(),
			components.PageProps{},
			components.Layout(
				l.Hx(),
				components.LayoutProps{},
				components.SubNav(
					components.SubNavProps{},
					components.SubNavBreadcrumb(
						components.SubNavBreadcrumbProps{},
						breadcrumbs.Breadcrumbs(
							breadcrumbs.BreadcrumbsProps{},
							breadcrumbs.Breadcrumb(
								breadcrumbs.BreadcrumbProps{
									Href:  "/",
									Title: "Home",
								},
							),
							breadcrumbs.Breadcrumb(
								breadcrumbs.BreadcrumbProps{
									Href:  "/lenses/list",
									Title: "Lenses",
								},
							),
						),
					),
				),
				components.Wrap(
					components.WrapProps{},
					htmx.FormElement(
						htmx.ID("new-lens-form"),
						htmx.HxPost("/lenses/new"),
						htmx.HxEncoding("multipart/form-data"),
						htmx.Label(
							htmx.ClassNames{
								"form-control": true,
								"w-full":       true,
								"max-w-xs":     true,
							},
							htmx.Input(
								htmx.Attribute("type", "file"),
								htmx.Attribute("name", "spec"),
								htmx.ClassNames{
									"file-input":          true,
									"file-input-bordered": true,
									"w-full":              true,
									"max-w-xs":            true,
								},
							),
							htmx.Input(
								htmx.Attribute("type", "text"),
								htmx.Attribute("name", "tag"),
								htmx.Attribute("placeholder", "Tag ..."),
								htmx.ClassNames{
									"input":          true,
									"input-bordered": true,
									"w-full":         true,
									"max-w-xs":       true,
								},
							),
							htmx.Progress(
								htmx.Attribute("id", "progress"),
								htmx.Value("0"),
								htmx.Max("100"),
							),
						),
						htmx.Button(
							htmx.ClassNames{
								"btn":         true,
								"btn-default": true,
								"my-4":        true,
							},
							htmx.Attribute("type", "submit"),
							htmx.Text("Create Lens"),
						),
					),
				),
			),
		),
	)
}

// Get ...
// func (l *LensesIndexController) Get() error {
// 	id, err := uuid.Parse(l.Hx().Context().Params("id"))
// 	if err != nil {
// 		return err
// 	}

// 	lens, err := l.db.GetLensByID(l.Hx().Context().Context(), id)
// 	if err != nil {
// 		return err
// 	}

// 	return l.Hx().RenderComp(
// 		components.Page(
// 			l.Hx(),
// 			components.PageProps{
// 				Children: []htmx.Node{
// 					htmx.FormElement(
// 						htmx.HxPost("/lenses"),
// 						htmx.Label(
// 							htmx.ClassNames{
// 								"form-control": true,
// 								"w-full":       true,
// 								"max-w-lg":     true,
// 								"mb-4":         true,
// 							},
// 							htmx.Div(
// 								htmx.ClassNames{
// 									"label": true,
// 								},
// 								htmx.Span(
// 									htmx.ClassNames{
// 										"label-text": true,
// 									},
// 									htmx.Text("What is your name?"),
// 								),
// 							),
// 							htmx.Input(
// 								htmx.Attribute("type", "text"),
// 								htmx.Attribute("name", "name"),
// 								htmx.Attribute("placeholder", "Name ..."),
// 								htmx.Attribute("value", lens.Name),
// 								htmx.Attribute("readonly", "true"),
// 								htmx.Attribute("disabled", "true"),
// 								htmx.ClassNames{
// 									"input":          true,
// 									"input-bordered": true,
// 									"w-full":         true,
// 									"max-w-lg":       true,
// 								},
// 							),
// 						),
// 						htmx.Label(
// 							htmx.ClassNames{
// 								"form-control": true,
// 								"w-full":       true,
// 								"max-w-lg":     true,
// 							},
// 							htmx.Div(
// 								htmx.ClassNames{
// 									"label":   true,
// 									"sr-only": true,
// 								},
// 							),
// 							htmx.Input(
// 								htmx.Attribute("type", "text"),
// 								htmx.Attribute("name", "description"),
// 								htmx.Attribute("placeholder", "Description ..."),
// 								htmx.Attribute("value", lens.Description),
// 								htmx.Attribute("readonly", "true"),
// 								htmx.Attribute("disabled", "true"),
// 								htmx.ClassNames{
// 									"input":          true,
// 									"input-bordered": true,
// 									"w-full":         true,
// 									"max-w-lg":       true,
// 								},
// 							),
// 						),
// 						htmx.Div(
// 							htmx.ClassNames{
// 								"divider": true,
// 							},
// 						),
// 						htmx.Div(
// 							htmx.ClassNames{
// 								"flex":     true,
// 								"flex-col": true,
// 								"py-2":     true,
// 							},
// 							htmx.H4(
// 								htmx.ClassNames{
// 									"text-gray-500": true,
// 								},
// 								htmx.Text("Last updated"),
// 							),
// 							htmx.H3(
// 								htmx.Text(lens.UpdatedAt.Format("2006-01-02 15:04:05")),
// 							),
// 						),
// 						htmx.Div(
// 							htmx.ClassNames{
// 								"flex":     true,
// 								"flex-col": true,
// 								"py-2":     true,
// 							},
// 							htmx.H4(
// 								htmx.ClassNames{
// 									"text-gray-500": true,
// 								},
// 								htmx.Text("Created at"),
// 							),
// 							htmx.H3(
// 								htmx.Text(
// 									lens.CreatedAt.Format("2006-01-02 15:04:05"),
// 								),
// 							),
// 						),
// 					),
// 				},
// 			},
// 		),
// 	)
// }

// Get ...
func (p *LensListController) Get() error {
	hx := p.Hx()

	limit, offset := tables.PaginationPropsFromContext(hx.Ctx())

	team, err := utils.TeamFromContext(hx.Ctx())
	if err != nil {
		return err
	}

	lenses, err := p.db.ListLenses(hx.Ctx().Context(), team, &models.Pagination{Limit: limit, Offset: offset})
	if err != nil {
		return err
	}

	table := tables.Table[*models.Lens](
		tables.TableProps[*models.Lens]{
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
							htmx.Text(row.Name),
						)
					},
				},
			},
			Rows: tables.NewRows(lenses),
		},
		htmx.ID("data-table"),
	)

	if hx.IsHxRequest() {
		hx.ReplaceURL(fmt.Sprintf("%s?limit=%d,offset=%d", hx.Ctx().Path(), limit, offset))

		return hx.RenderComp(table)
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
						table,
						htmx.Div(
							htmx.ClassNames{
								"bg-base-100": true,
								"p-4":         true,
							},
							tables.Pagination(
								tables.PaginationProps{
									Limit:  limit,
									Offset: offset,
								},
								tables.Prev(
									tables.PaginationProps{
										URL:    "/api/data",
										Offset: offset,
										Limit:  limit,
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
		),
	)
}
