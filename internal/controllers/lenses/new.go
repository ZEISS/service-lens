package lenses

import (
	"bytes"
	"fmt"
	"io"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/progress"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
)

// LensNewController ...
type LensNewController struct {
	db ports.Repository

	htmx.DefaultController
}

// NewLensNewController ...
func NewLensNewController(db ports.Repository) *LensNewController {
	return &LensNewController{
		db: db,
	}
}

// Prepare ...
func (l *LensNewController) Prepare() error {
	if err := l.BindValues(utils.User(l.db), utils.Team(l.db)); err != nil {
		return err
	}

	return nil
}

// Post ...
func (l *LensNewController) Post() error {
	hx := l.Hx()

	spec, err := hx.Ctx().FormFile("spec")
	if err != nil {
		return err
	}
	file, err := spec.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	team := htmx.Locals[*authz.Team](l.DefaultCtx(), utils.ValuesKeyTeam)

	lens := &models.Lens{}
	err = lens.UnmarshalJSON(buf.Bytes())
	if err != nil {
		return err
	}
	lens.Team = *team

	lens, err = l.db.AddLens(hx.Ctx().Context(), lens)
	if err != nil {
		return err
	}

	hx.Redirect(fmt.Sprintf("/%s/lenses/%s", team.Slug, lens.ID))

	return nil
}

// Get ...
func (l *LensNewController) Get() error {
	return l.Hx().RenderComp(
		components.Page(
			l.DefaultCtx(),
			components.PageProps{},
			htmx.DataAttribute("theme", "light"),
			components.Layout(
				l.DefaultCtx(),
				components.LayoutProps{},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("New Lens"),
						),
						htmx.FormElement(
							htmx.ID("new-lens-form"),
							htmx.HxEncoding("multipart/form-data"),
							htmx.HxPost(fmt.Sprintf("/%s/lenses/new", htmx.Locals[*authz.Team](l.DefaultCtx(), utils.ValuesKeyTeam).Slug)),
							htmx.Attribute("_", "on htmx:xhr:progress(loaded, total) set #new-lens-progress.value to (loaded/total)*100'"),
							htmx.Div(
								forms.FileInputBordered(
									forms.FileInputProps{},
									htmx.Attribute("name", "spec"),
								),
							),
							progress.Progress(
								progress.ProgressProps{
									ClassNames: htmx.ClassNames{
										"block": true,
										"my-4":  true,
									},
								},
								htmx.ID("new-lens-progress"),
								htmx.Value("0"),
								htmx.Max("100"),
							),
							buttons.OutlinePrimary(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"my-4": true,
									},
								},
								htmx.Text("Create Lens"),
								htmx.HxDisabledElt("this"),
							),
						),
					),
				),
			),
		),
	)
}
