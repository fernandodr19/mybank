package app

import (
	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
)

// App contains application's usecases
type App struct {
	Transactions *transactions.Usecase
}

func BuildApp(txUsecase *transactions.Usecase) (*App, error) {
	return &App{
		Transactions: txUsecase,
	}, nil
}
