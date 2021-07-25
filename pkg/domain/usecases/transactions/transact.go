package transactions

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/domain"
	"github.com/fernandodr19/mybank-tx/pkg/domain/entities"
	"github.com/fernandodr19/mybank-tx/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
	"github.com/fernandodr19/mybank-tx/pkg/instrumentation/logger"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Transact executes a transaction
func (u Usecase) Transact(ctx context.Context, accID vos.AccountID, op operations.Operation, amount vos.Money) (vos.TransactionID, error) {
	const operation = "transactions.TransactionUsecase.Transact"

	log := logger.FromCtx(ctx).WithFields(logrus.Fields{
		"accID":     accID,
		"operation": op.String(),
		"amout":     amount.Int(),
	})

	log.Infoln("processing transaction")

	//validate acc id
	_, err := uuid.Parse(accID.String())
	if err != nil {
		return "", ErrInvalidAccID
	}

	// validate amount
	if amount <= 0 {
		return "", ErrInvalidAmount
	}

	switch op {
	case operations.Debit:
		err = u.handleDebit(ctx, accID, amount)
		amount *= -1
	case operations.Credit:
		err = u.handleCredit(ctx, accID, amount)
		amount *= -1
	case operations.Withdrawal:
		err = u.handleWithdrawal(ctx, accID, amount)
		amount *= -1
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

	log.WithField("txID", txID).Infoln("transaction successfully processed")

	return txID, nil
}

func (u Usecase) handleDebit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handleDebit"
	err := u.accountsClient.Withdrawal(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}

func (u Usecase) handleCredit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handleCredit"
	err := u.accountsClient.ReserveCreditLimit(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}

func (u Usecase) handleWithdrawal(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handleWithdrawal"
	err := u.accountsClient.Withdrawal(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}

func (u Usecase) handlePayment(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "transactions.TransactionUsecase.handlePayment"
	err := u.accountsClient.Deposit(ctx, accID, amount)
	if err != nil {
		return domain.Error(operation, err)
	}
	return nil
}
