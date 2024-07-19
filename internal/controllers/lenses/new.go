package lenses

import (
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/progress"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// LensNewController ...
type LensNewControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensNewController ...
func NewLensController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *LensNewControllerImpl {
	return &LensNewControllerImpl{
		store: store,
	}
}

// // Post ...
// func (l *LensNewControllerImpl) Post() error {
// 	spec, err := l.Ctx().FormFile("spec")
// 	if err != nil {
// 		return err
// 	}
// 	file, err := spec.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	buf := bytes.NewBuffer(nil)
// 	if _, err := io.Copy(buf, file); err != nil {
// 		return err
// 	}

// 	team := l.Values(utils.ValuesKeyTeam).(*authz.Team)

// 	lens := &models.Lens{}
// 	err = lens.UnmarshalJSON(buf.Bytes())
// 	if err != nil {
// 		return err
// 	}
// 	lens.Team = *team

// 	lens, err = l.db.AddLens(l.Context(), lens)
// 	if err != nil {
// 		return err
// 	}

// 	l.Hx().Redirect(fmt.Sprintf("/teams/%s/lenses/%s/index", team.Slug, lens.ID))

// 	return nil
// }

// Get ...
func (l *LensNewControllerImpl) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{},
			htmx.DataAttribute("theme", "light"),
			components.Layout(
				components.LayoutProps{
					Path: l.Path(),
				},
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
