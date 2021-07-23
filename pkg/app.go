package app

import (
	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
)

// App contains application's usecases
type App struct {
	Transactions *transactions.TransactionsUsecase
}

// BuildApp builds application struct with its necessary usecases
func BuildApp(txUsecase *transactions.TransactionsUsecase) (*App, error) {
	return &App{
		Transactions: txUsecase,
	}, nil
}
