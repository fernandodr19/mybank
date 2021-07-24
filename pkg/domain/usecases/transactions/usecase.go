package transactions

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/domain/entities"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
)

// AccountsClient responsible for communicating with accounts server
type AccountsClient interface {
	Deposit(ctx context.Context, accID vos.AccountID, amount vos.Money) error
	Withdrawal(ctx context.Context, accID vos.AccountID, amount vos.Money) error
	ReserveCreditLimit(ctx context.Context, accID vos.AccountID, amount vos.Money) error
}

// Repository of transactions
type Repository interface {
	SaveTransaction(context.Context, entities.Transaction) (vos.TransactionID, error)
}

// Usecase of transactions
type Usecase struct {
	transactionsRepo Repository
	accountsClient   AccountsClient
}

// NewUsecase builds a tx usecase
func NewUsecase(txRepo Repository, accClient AccountsClient) *Usecase {
	return &Usecase{
		transactionsRepo: txRepo,
		accountsClient:   accClient,
	}
}
