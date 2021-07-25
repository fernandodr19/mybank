package app

import (
	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/db/postgres"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/grpc/accounts"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

// App contains application's usecases
type App struct {
	Transactions *transactions.Usecase
}

// BuildApp builds application struct with its necessary usecases
func BuildApp(dbConn *pgx.Conn, grpcConn *grpc.ClientConn) *App {
	txRepo := postgres.NewTransactionRepository(dbConn)
	accClient := accounts.NewClient(grpcConn)
	return &App{
		Transactions: transactions.NewUsecase(txRepo, accClient),
	}
}
