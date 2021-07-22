package transactions

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain/entities"
)

// Usecase of transactions
type Usecase struct {
}

// Repository of accounts
type Repository interface {
	SaveTransaction(context.Context, *entities.Transaction) error
}

func NewUsecase() *Usecase {
	return &Usecase{}
}
