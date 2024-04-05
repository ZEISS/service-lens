package lenses

import (
	"bytes"
	"fmt"
	"io"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/progress"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/resolvers"

	htmx "github.com/zeiss/fiber-htmx"
)

// LensEditControllerParams ...
type LensEditControllerParams struct {
	ID   uuid.UUID `json:"id" xml:"id" form:"id"`
	Team string    `json:"team" xml:"team" form:"team"`
}

// NewDefaultLensEditControllerParams ...
func NewDefaultLensEditControllerParams() *LensEditControllerParams {
	return &LensEditControllerParams{}
}

// LensEditController ...
type LensEditController struct {
	db     ports.Repository
	team   *authz.Team
	params *LensEditControllerParams
	lens   *models.Lens

	htmx.UnimplementedController
}

// NewLensEditController ...
func NewLensEditController(db ports.Repository) *LensEditController {
	return &LensEditController{
		db: db,
	}
}

// Prepare ...
func (l *LensEditController) Prepare() error {
	team := l.Hx().Values(resolvers.ValuesKeyTeam).(*authz.Team)
	l.team = team

	l.params = NewDefaultLensEditControllerParams()
	if err := l.Hx().Ctx().ParamsParser(l.params); err != nil {
		return err
	}

	lens, err := l.db.GetLensByID(l.Hx().Ctx().Context(), l.params.ID)
	if err != nil {
		return err
	}
	l.lens = lens

	return nil
}

// Post ...
func (l *LensEditController) Post() error {
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

	lens := &models.Lens{}
	err = lens.UnmarshalJSON(buf.Bytes())
	if err != nil {
		return err
	}
	lens.Team = *l.team

	lens, err = l.db.AddLens(hx.Ctx().Context(), lens)
	if err != nil {
		return err
	}

	hx.Redirect(fmt.Sprintf("/%s/lenses/%s", l.team.Slug, lens.ID))

	return nil
}

// Get ...
func (l *LensEditController) Get() error {
	return l.Hx().RenderComp(
		components.Page(
			l.Hx(),
			components.PageProps{},
			htmx.DataAttribute("theme", "light"),
			components.Layout(
				l.Hx(),
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
							htmx.HxPost(fmt.Sprintf("/%s/lenses/new", l.team.Slug)),
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
