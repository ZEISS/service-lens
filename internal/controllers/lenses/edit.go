package lenses

import (
	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/progress"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// LensEditControllerParams ...
type LensEditControllerParams struct {
	ID uuid.UUID `json:"id" xml:"id" form:"id"`
}

// NewDefaultLensEditControllerParams ...
func NewDefaultLensEditControllerParams() *LensEditControllerParams {
	return &LensEditControllerParams{}
}

// LensEditController ...
type LensEditController struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensEditController ...
func NewLensEditController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *LensEditController {
	return &LensEditController{
		store: store,
	}
}

// Prepare ...
func (l *LensEditController) Prepare() error {
	// if err := l.BindValues(utils.User(l.db), utils.Team(l.db)); err != nil {
	// 	return err
	// }

	// l.params = NewDefaultLensEditControllerParams()
	// if err := l.BindParams(l.params); err != nil {
	// 	return err
	// }

	// lens, err := l.db.GetLensByID(l.Context(), l.params.ID)
	// if err != nil {
	// 	return err
	// }
	// l.lens = lens

	return nil
}

// Post ...
func (l *LensEditController) Post() error {
	// spec, err := l.Ctx().FormFile("spec")
	// if err != nil {
	// 	return err
	// }
	// file, err := spec.Open()
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// buf := bytes.NewBuffer(nil)
	// if _, err := io.Copy(buf, file); err != nil {
	// 	return err
	// }

	// lens := &models.Lens{}
	// err = lens.UnmarshalJSON(buf.Bytes())
	// if err != nil {
	// 	return err
	// }
	// lens.Team = *l.team

	// lens, err = l.db.AddLens(l.Context(), lens)
	// if err != nil {
	// 	return err
	// }

	// l.Hx().Redirect(fmt.Sprintf("/%s/lenses/%s", l.team.Slug, lens.ID))

	return nil
}

// Get ...
func (l *LensEditController) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{},
			htmx.DataAttribute("theme", "light"),
			components.Layout(
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
							htmx.HxPost("/lenses/new"),
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
