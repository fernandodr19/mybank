package app

import (
	"github.com/fernandodr19/mybank/pkg/domain/usecases/transactions"
)

// App contains application's usecases
type App struct {
	Transaction *transactions.Usecase
}

func BuildApp(txUsecase *transactions.Usecase) (*App, error) {
	return &App{
		Transaction: txUsecase,
	}, nil
}
