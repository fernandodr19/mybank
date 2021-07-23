package postgres

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/domain"
	"github.com/fernandodr19/mybank-tx/pkg/domain/entities"
	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/db/postgres/sqlc"
	"github.com/jackc/pgx/v4"
)

var _ transactions.Repository = &TransactionsRepository{}

// TransactionsRepository is the repository of transactions
type TransactionsRepository struct {
	conn *pgx.Conn
	q    *sqlc.Queries
}

// NewTransactionRepository returns a transaction repository
func NewTransactionRepository(conn *pgx.Conn) *TransactionsRepository {
	return &TransactionsRepository{
		conn: conn,
		q:    sqlc.New(conn),
	}
}

// SaveTransaction saves the transaction on DB, returning its auto generated ID in case of success
func (r TransactionsRepository) SaveTransaction(ctx context.Context, tx entities.Transaction) (vos.TransactionID, error) {
	const operation = "postgres.TransactionsRepository.SaveTransaction"
	txID, err := r.q.SaveTransaction(ctx, sqlc.SaveTransactionParams{
		AccountID:     tx.AccountID.String(),
		OperationType: int32(tx.Operaion),
		Amount:        int32(tx.Amount.Int()),
	})
	if err != nil {
		return "", domain.Error(operation, err)
	}
	return vos.TransactionID(txID), nil
}
