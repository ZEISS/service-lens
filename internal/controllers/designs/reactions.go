package designs

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/service-lens/internal/components/designs"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// ReactionControllerImpl ...
type ReactionControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewReactionController ...
func NewReactionController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ReactionControllerImpl {
	return &ReactionControllerImpl{store: store}
}

// Error ...
func (l *ReactionControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Delete ...
func (l *ReactionControllerImpl) Delete() error {
	var params struct {
		ID         uuid.UUID `json:"id" params:"id"`
		ReactionID int       `json:"reaction_id" params:"reaction_id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	reaction := models.Reaction{
		ID:            params.ReactionID,
		ReactorID:     l.Session().User.ID,
		ReactableID:   params.ID,
		ReactableType: models.ReactableTypeDesign,
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteReaction(ctx, &reaction)
	})
	if err != nil {
		return err
	}

	design := models.Design{
		ID: params.ID,
	}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &design)
	})
	if err != nil {
		return err
	}

	return l.Render(
		htmx.Fragment(
			designs.DesignReactions(
				designs.DesignReactionsProps{
					User:   l.Session().User,
					Design: design,
				},
			),
		),
	)
}

// Post ...
func (l *ReactionControllerImpl) Post() error {
	var params struct {
		ID uuid.UUID `json:"id" params:"id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	var body struct {
		Reaction string `json:"emoji" form:"emoji" validate:"required"`
	}

	err = l.BindBody(&body)
	if err != nil {
		return err
	}

	reaction := models.Reaction{
		ReactableID:   params.ID,
		ReactableType: models.ReactableTypeDesign,
		Value:         body.Reaction,
		ReactorID:     l.Session().User.ID,
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateReaction(ctx, &reaction)
	})
	if err != nil {
		return err
	}

	design := models.Design{
		ID: params.ID,
	}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &design)
	})
	if err != nil {
		return err
	}

	return l.Render(
		htmx.Fragment(
			designs.DesignReactions(
				designs.DesignReactionsProps{
					User:   l.Session().User,
					Design: design,
				},
			),
		),
	)
}
