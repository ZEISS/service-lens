package handlers

import (
	"github.com/gofiber/fiber/v2"
	goth "github.com/zeiss/fiber-goth"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/collapsible"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/joins"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	reload "github.com/zeiss/fiber-reload"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/service-lens/internal/components"
)

type SettingsHandler struct{}

func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{}
}

func (h *SettingsHandler) ListSettings(c *fiber.Ctx) (htmx.Node, error) {
	return components.DefaultLayout(
		components.DefaultLayoutProps{
			Path:        c.Path(),
			User:        errorx.Ignore(goth.SessionFromContext(c)).User,
			Development: reload.IsDevelopment(c.UserContext()),
		},
		func() htmx.Node {
			return htmx.Fragment(
				cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.Merge(
							htmx.ClassNames{
								tailwind.M2: true,
							},
						),
					},
					htmx.Form(
						htmx.HxPut(""),
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Settings"),
							),
							joins.JoinVertical(
								joins.JoinProps{
									ClassNames: htmx.ClassNames{
										tailwind.WFull: true,
									},
								},
								collapsible.CollapseArrow(
									collapsible.CollapseProps{
										ClassNames: htmx.ClassNames{
											"join-item": true,
										},
									},
									collapsible.CollapseRadio(
										collapsible.CollapseRadioProps{
											Name: "authentication_settings",
										},
									),
									collapsible.CollapseTitle(
										collapsible.CollapseTitleProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Microsoft Entra ID"),
									),
									collapsible.CollapseContent(
										collapsible.CollapseContentProps{},
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{
													"flex":            true,
													"justify-between": true,
													"flex-row":        true,
												},
											},
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{},
													htmx.Text("Enable"),
												),
											),
											forms.Toggle(
												forms.ToggleProps{
													Name:  "entra_id_enabled",
													Value: "true",
												},
											),
										),
									),
								),
								collapsible.CollapseArrow(
									collapsible.CollapseProps{
										ClassNames: htmx.ClassNames{
											"join-item": true,
										},
									},
									collapsible.CollapseRadio(
										collapsible.CollapseRadioProps{
											Name: "authentication_settings",
										},
									),
									collapsible.CollapseTitle(
										collapsible.CollapseTitleProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("GitHub"),
									),
									collapsible.CollapseContent(
										collapsible.CollapseContentProps{},
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{
													"flex":            true,
													"justify-between": true,
													"flex-row":        true,
												},
											},
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{},
													htmx.Text("Enable"),
												),
											),
											forms.Toggle(
												forms.ToggleProps{
													Name:  "github_enabled",
													Value: "true",
												},
											),
										),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{},
									htmx.Attribute("type", "submit"),
									htmx.Text("Save"),
								),
							),
						),
					),
				),
			)
		},
	), nil
}
