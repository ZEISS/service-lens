package controllers

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/breadcrumbs"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
)

// Lenses ...
type Lenses struct {
	db ports.Repository
}

// NewLensesController ...
func NewLensesController(db ports.Repository) *Lenses {
	return &Lenses{db}
}

// Store ...
func (l *Lenses) Store(hx *htmx.Htmx) error {
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

// New ...
func (l *Lenses) New(c *fiber.Ctx) (htmx.Node, error) {
	return components.Page(
		components.PageProps{},
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
	), nil
}

// Show ...
func (l *Lenses) Show(c *fiber.Ctx) (htmx.Node, error) {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return nil, err
	}

	lens, err := l.db.GetLensByID(c.Context(), id)
	if err != nil {
		return nil, err
	}

	return components.Page(
		components.PageProps{
			Children: []htmx.Node{
				htmx.FormElement(
					htmx.HxPost("/lenses"),
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
								htmx.Text("What is your name?"),
							),
						),
						htmx.Input(
							htmx.Attribute("type", "text"),
							htmx.Attribute("name", "name"),
							htmx.Attribute("placeholder", "Name ..."),
							htmx.Attribute("value", lens.Name),
							htmx.Attribute("readonly", "true"),
							htmx.Attribute("disabled", "true"),
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
							htmx.Attribute("value", lens.Description),
							htmx.Attribute("readonly", "true"),
							htmx.Attribute("disabled", "true"),
							htmx.ClassNames{
								"input":          true,
								"input-bordered": true,
								"w-full":         true,
								"max-w-lg":       true,
							},
						),
					),
					htmx.Div(
						htmx.ClassNames{
							"divider": true,
						},
					),
					htmx.Div(
						htmx.ClassNames{
							"flex":     true,
							"flex-col": true,
							"py-2":     true,
						},
						htmx.H4(
							htmx.ClassNames{
								"text-gray-500": true,
							},
							htmx.Text("Last updated"),
						),
						htmx.H3(
							htmx.Text(lens.UpdatedAt.Format("2006-01-02 15:04:05")),
						),
					),
					htmx.Div(
						htmx.ClassNames{
							"flex":     true,
							"flex-col": true,
							"py-2":     true,
						},
						htmx.H4(
							htmx.ClassNames{
								"text-gray-500": true,
							},
							htmx.Text("Created at"),
						),
						htmx.H3(
							htmx.Text(
								lens.CreatedAt.Format("2006-01-02 15:04:05"),
							),
						),
					),
				),
			},
		},
	), nil
}

// List ...
func (l *Lenses) List(c *fiber.Ctx) (htmx.Node, error) {
	lenses, err := l.db.ListLenses(c.Context(), &models.Pagination{Limit: 10, Offset: 0})
	if err != nil {
		return nil, err
	}

	items := make([]htmx.Node, len(lenses))
	for i, lens := range lenses {
		tags := make([]htmx.Node, len(lens.Tags))
		for j, tag := range lens.Tags {
			tags[j] = htmx.Span(
				htmx.ClassNames{
					"badge":         true,
					"badge-primary": true,
					"badge-outline": true,
				},
				htmx.Text(tag.Name),
			)
		}

		items[i] = htmx.Tr(
			htmx.Th(
				htmx.Label(
					htmx.Input(
						htmx.ClassNames{
							"checkbox": true,
						},
						htmx.Attribute("type", "checkbox"),
						htmx.Attribute("name", "lens"),
						htmx.Attribute("value", lens.ID.String()),
					),
				),
			),
			htmx.Th(htmx.Text(lens.ID.String())),
			htmx.Td(htmx.Text(lens.Name)),
			htmx.Td(htmx.Text(lens.Description)),
			htmx.Td(tags...),
		)
	}

	return components.Page(
		components.PageProps{}.WithContext(c),
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
			components.SubNavActions(
				components.SubNavActionsProps{},
				links.Link(
					links.LinkProps{
						Href: "/lenses/new",
						ClassNames: htmx.ClassNames{
							"btn":         true,
							"btn-outline": true,
							"btn-xs":      true,
							"link-hover":  true,
						},
					},
					htmx.Text("Create Lens"),
				),
			),
		),
		components.Wrap(
			components.WrapProps{},
			htmx.Div(
				htmx.ClassNames{"overflow-x-auto": true},
				htmx.Table(
					htmx.ClassNames{"table": true},
					htmx.THead(
						htmx.Tr(
							htmx.Th(
								htmx.Label(
									htmx.Input(
										htmx.ClassNames{
											"checkbox": true,
										},
										htmx.Attribute("type", "checkbox"),
										htmx.Attribute("name", "all"),
									),
								),
							),
							htmx.Th(htmx.Text("ID")),
							htmx.Th(htmx.Text("Name")),
							htmx.Th(htmx.Text("Description")),
						),
					),
					htmx.TBody(
						items...,
					),
				),
				htmx.Div(
					htmx.ClassNames{},
					htmx.Select(
						htmx.ClassNames{
							"select":   true,
							"max-w-xs": true,
						},
						htmx.Option(
							htmx.Text("10"),
							htmx.Attribute("value", "10"),
						),
						htmx.Option(
							htmx.Text("20"),
							htmx.Attribute("value", "20"),
						),
						htmx.Option(
							htmx.Text("30"),
							htmx.Attribute("value", "30"),
						),
					),
				),
			),
		),
	), nil
}
