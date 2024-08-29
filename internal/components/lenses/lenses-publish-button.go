package lenses

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/service-lens/internal/utils"
)

// LensesPublishButtonProps ...
type LensesPublishButtonProps struct {
	// ClassNames ...
	ClassNames htmx.ClassNames
	// ID ...
	ID uuid.UUID
	// IsDraft ...
	IsDraft bool
}

// LensesPublishButton ...
func LensesPublishButton(props LensesPublishButtonProps) htmx.Node {
	return htmx.IfElse(
		props.IsDraft,
		buttons.Button(
			buttons.ButtonProps{},
			htmx.HxConfirm("Are you sure you want to publish this lens?"),
			htmx.HxPost(fmt.Sprintf(utils.PublishLensUrlFormat, props.ID)),
			htmx.HxSwap("outerHTML"),
			htmx.Text("Publish"),
		),
		buttons.Button(
			buttons.ButtonProps{},
			htmx.HxConfirm("Are you sure you want to unpublish this lens?"),
			htmx.HxDelete(fmt.Sprintf(utils.UnpublishLensUrlFormat, props.ID)),
			htmx.HxSwap("outerHTML"),
			htmx.Text("Unpublish"),
		),
	)
}
