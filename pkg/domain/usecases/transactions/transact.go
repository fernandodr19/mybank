package transactions

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/domain"
	"github.com/fernandodr19/mybank-tx/pkg/domain/entities"
	"github.com/fernandodr19/mybank-tx/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
	"github.com/google/uuid"
)

// Transact executes a transaction
func (u TransactionsUsecase) Transact(ctx context.Context, accID vos.AccountID, op operations.Operation, amount vos.Money) (vos.TransactionID, error) {
	const operation = "transactions.TransactionUsecase.Transact"

	//validate acc id
	_, err := uuid.Parse(accID.String())
	if err != nil {
		return "", ErrInvalidAccID
	}

	// validate amount
	if amount < 0 {
		return "", ErrInvalidAmount
	}

	switch op {
	case operations.Debit:
		err = u.handleDebit(ctx, accID, amount)
	case operations.Credit:
		err = u.handleCredit(ctx, accID, amount)
	case operations.Withdrawal:
		err = u.handleWithdrawal(ctx, accID, amount)
	case operations.Payment:
		err = u.handlePayment(ctx, accID, amount)
	default:
		return "", domain.Error(operation, operations.ErrInvalidOperation)
	}

	if err != nil {
		return "", domain.Error(operation, err)
	}

	txID, err := u.transactionsRepo.SaveTransaction(ctx, entities.Transaction{
		AccountID: accID,
		Operaion:  op,
		Amount:    amount,
	})
	if err != nil {
		return "", domain.Error(operation, err)
	}

	return txID, nil
}

func (u TransactionsUsecase) handleDebit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handleDebit"
	err := u.accountsClient.Withdrawal(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}

func (u TransactionsUsecase) handleCredit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handleCredit"
	err := u.accountsClient.ReserveCreditLimit(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}

func (u TransactionsUsecase) handleWithdrawal(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handleWithdrawal"
	err := u.accountsClient.Withdrawal(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}

func (u TransactionsUsecase) handlePayment(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handlePayment"
	err := u.accountsClient.Deposit(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}
