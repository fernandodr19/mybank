package app

import (
	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
)

// App contains application's usecases
type App struct {
	Transactions *transactions.TransactionsUsecase
}

// BuildApp builds application struct with its necessary usecases
func BuildApp(txRepo transactions.Repository, accClient transactions.AccountsClient) (*App, error) {
	return &App{
		Transactions: transactions.NewUsecase(txRepo, accClient),
	}, nil
}
