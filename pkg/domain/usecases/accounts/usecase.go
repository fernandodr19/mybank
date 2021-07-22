package accounts

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain/entities"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
)

// Usecase of accounts
type Usecase struct {
}

// Repository of accounts
type Repository interface {
	GetAccountByID(context.Context, vos.AccountID) (entities.Account, error)
	CreateAccount(context.Context, vos.Document) (vos.AccountID, error)
}

// newUsecase returns an account usecase
func NewUsecase() *Usecase {
	return &Usecase{}
}
