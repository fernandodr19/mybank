package accounts

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/domain"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
	"github.com/fernandodr19/mybank-tx/pkg/instrumentation/logger"
	"google.golang.org/grpc"
)

// Client gRPC of accounts
type Client struct {
	client AccountsServiceClient
}

// NewClient returns a gRPC client
func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: NewAccountsServiceClient(conn),
	}

}

// Deposit(ctx context.Context, accID vos.AccountID, amount vos.Money) error
// Withdrawal(ctx context.Context, accID vos.AccountID, amount vos.Money) error
// ReserveCreditLimit(ctx context.Context, accID vos.AccountID, amount vos.Money) error

// Deposit requests a deposit to the accounts server
func (c Client) Deposit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "accounts.Client.Deposit"
	reply, err := c.client.Deposit(ctx, &Request{
		AccountID: accID.String(),
		Amount:    amount.Int64(),
	})
	if err != nil {
		return domain.Error(operation, err)
	}
	logger.Default().WithField("errorCode", reply.ErrorCode)
	return nil
}

// Withdrawal requests a withdrawal to the accounts server
func (c Client) Withdrawal(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "accounts.Client.Withdrawal"
	reply, err := c.client.Withdrawal(ctx, &Request{
		AccountID: accID.String(),
		Amount:    amount.Int64(),
	})
	if err != nil {
		return domain.Error(operation, err)
	}
	logger.Default().WithField("errorCode", reply.ErrorCode)
	return nil
}

// ReserveCreditLimit requests a credit limit reserval to the accounts server
func (c Client) ReserveCreditLimit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "accounts.Client.ReserveCreditLimit"
	reply, err := c.client.ReserveCreditLimit(ctx, &Request{
		AccountID: accID.String(),
		Amount:    amount.Int64(),
	})
	if err != nil {
		return domain.Error(operation, err)
	}
	logger.Default().WithField("errorCode", reply.ErrorCode)
	return nil
}
