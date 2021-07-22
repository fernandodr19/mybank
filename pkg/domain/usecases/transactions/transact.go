package transactions

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
)

// Transact executes a transaction
func (u TransactionsUsecase) Transact(ctx context.Context, accID vos.AccountID, op operations.Operation, amount vos.Money) (vos.TransactionID, error) {
	return "", nil
}
