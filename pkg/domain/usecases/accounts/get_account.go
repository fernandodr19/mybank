package accounts

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain/entities"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
)

// GetAccountByID retrieves an account based on a given ID
func (u Usecase) GetAccountByID(ctx context.Context, accID vos.AccountID) (entities.Account, error) {
	return entities.Account{}, nil
}
