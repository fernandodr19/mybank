package app

import (
	"github.com/fernandodr19/mybank/pkg/domain/usecases/accounts"
	"github.com/fernandodr19/mybank/pkg/domain/usecases/transactions"
)

// App contains application's usecases
type App struct {
	Accounts    *accounts.Usecase
	Transaction *transactions.Usecase
}

func BuildApp() (*App, error) {
	return &App{}, nil
}
