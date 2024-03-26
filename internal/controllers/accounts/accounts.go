package accounts

import (
	"github.com/zeiss/service-lens/internal/ports"
)

// Accounts ...
type Accounts struct {
	db ports.Repository
}

// NewAccountsController ...
func NewAccountsController(db ports.Repository) *Accounts {
	return &Accounts{db}
}

// Show ...
func (a *Accounts) Show() {}
