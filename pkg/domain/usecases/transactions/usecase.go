package transactions

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/domain/entities"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
)

type AccountsClient interface {
	Deposit(ctx context.Context, accID vos.AccountID, amount vos.Money) error
	Withdrawal(ctx context.Context, accID vos.AccountID, amount vos.Money) error
	ReserveCreditLimit(ctx context.Context, accID vos.AccountID, amount vos.Money) error
}

// Repository of transactions
type Repository interface {
	SaveTransaction(context.Context, entities.Transaction) (vos.TransactionID, error)
}

type TransactionsUsecase struct {
	transactionsRepo Repository
	accountsClient   AccountsClient
}

func NewUsecase(txRepo Repository, accClient AccountsClient) *TransactionsUsecase {
	return &TransactionsUsecase{
		transactionsRepo: txRepo,
		accountsClient:   accClient,
	}
}
