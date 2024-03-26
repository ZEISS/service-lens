package teams

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/service-lens/internal/ports"
)

// TeamDashboardController ...
type TeamDashboardController struct {
	db ports.Repository

	htmx.UnimplementedController
}

// NewTeamDashboardController ...
func NewTeamDashboardController(db ports.Repository) *TeamDashboardController {
	return &TeamDashboardController{db, htmx.UnimplementedController{}}
}
