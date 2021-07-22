package transactions

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain/entities"
	"github.com/fernandodr19/mybank/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
)

//go:generate moq -skip-ensure -stub -out mocks.gen.go . Usecase:TransactionsMockUsecase

// Usecase of transactions
type Usecase interface {
	Transact(ctx context.Context, accID vos.AccountID, op operations.Operation, amount vos.Money) (vos.TransactionID, error)
}

// Repository of transactions
type Repository interface {
	SaveTransaction(context.Context, *entities.Transaction) error
}

type TransactionsUsecase struct {
}

func NewUsecase() *TransactionsUsecase {
	return &TransactionsUsecase{}
}
