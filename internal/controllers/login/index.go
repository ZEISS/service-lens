package login

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/ports"
)

// LoginIndexController ...
type LoginIndexController struct {
	db  ports.Repository
	ctx htmx.Ctx

	htmx.DefaultController
}

// NewLoginIndexController ...
func NewLoginIndexController(db ports.Repository) *LoginIndexController {
	return &LoginIndexController{
		db: db,
	}
}

// Prepare ...
func (l *LoginIndexController) Prepare() error {
	ctx, err := htmx.NewDefaultContext(l.Hx().Ctx())
	if err != nil {
		return err
	}
	l.ctx = ctx

	return nil
}

// Get ...
func (l *LoginIndexController) Get() error {
	return l.Hx().RenderComp(
		components.Page(
			l.ctx,
			components.PageProps{},
			components.Wrap(
				components.WrapProps{},
				htmx.Section(
					htmx.Merge(
						htmx.ClassNames{
							"bg-gray-50 ":      true,
							"dark:bg-gray-900": true,
						},
					),
				),
				htmx.Div(
					htmx.Merge(
						htmx.ClassNames{
							"flex":           true,
							"flex-col":       true,
							"items-center":   true,
							"justify-center": true,
							"px-6":           true,
							"py-8":           true,
							"mx-auto":        true,
							"md:h-screen":    true,
							"lg:py-0":        true,
						},
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"w-96":     true,
								"max-w-lg": true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Sign in to your account"),
							),
							htmx.Div(
								htmx.ClassNames{
									"mt-4": true,
								},
								links.Button(
									links.LinkProps{
										ClassNames: htmx.ClassNames{
											"w-full":      true,
											"btn-primary": true,
											"btn-outline": true,
										},
										Href: "/login/github",
									},
									htmx.Text("Login on GitHub"),
								),
							),
							dividers.Divider(
								dividers.DividerProps{},
								htmx.Text("OR"),
							),
							htmx.Form(
								htmx.HxPost("/login"),
								forms.FormControl(
									forms.FormControlProps{
										ClassNames: htmx.ClassNames{
											"py-4": true,
										},
									},
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:        "username",
											Placeholder: "indy@jones.com",
										},
									),
								),
								forms.FormControl(
									forms.FormControlProps{},
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:        "password",
											Placeholder: "supersecret",
										},
										htmx.Type("password"),
									),
								),
								cards.Actions(
									cards.ActionsProps{
										ClassNames: htmx.ClassNames{
											"py-4":  true,
											"-mb-4": true,
										},
									},
									buttons.OutlinePrimary(
										buttons.ButtonProps{
											ClassNames: htmx.ClassNames{
												"w-full": true,
											},
										},
										htmx.Text("Login"),
									),
								),
							),
						),
					),
				),
			),
		),
	)

	//	<section class="bg-gray-50 dark:bg-gray-900">
	//	<div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
	//	    <a href="#" class="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
	//	        <img class="w-8 h-8 mr-2" src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/logo.svg" alt="logo">
	//	        Flowbite
	//	    </a>
	//	    <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
	//	        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
	//	            <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
	//	                Sign in to your account
	//	            </h1>
	//	            <form class="space-y-4 md:space-y-6" action="#">
	//	                <div>
	//	                    <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your email</label>
	//	                    <input type="email" name="email" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@company.com" required="">
	//	                </div>
	//	                <div>
	//	                    <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
	//	                    <input type="password" name="password" id="password" placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required="">
	//	                </div>
	//	                <div class="flex items-center justify-between">
	//	                    <div class="flex items-start">
	//	                        <div class="flex items-center h-5">
	//	                          <input id="remember" aria-describedby="remember" type="checkbox" class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800" required="">
	//	                        </div>
	//	                        <div class="ml-3 text-sm">
	//	                          <label for="remember" class="text-gray-500 dark:text-gray-300">Remember me</label>
	//	                        </div>
	//	                    </div>
	//	                    <a href="#" class="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500">Forgot password?</a>
	//	                </div>
	//	                <button type="submit" class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign in</button>
	//	                <p class="text-sm font-light text-gray-500 dark:text-gray-400">
	//	                    Don’t have an account yet? <a href="#" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</a>
	//	                </p>
	//	            </form>
	//	        </div>
	//	    </div>
	//	</div>
	//
	// </section>
}
