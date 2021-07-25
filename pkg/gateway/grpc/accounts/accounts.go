package accounts

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/domain"
	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ transactions.AccountsClient = &Client{}

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

// Deposit requests a deposit to the accounts server
func (c Client) Deposit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "accounts.Client.Deposit"
	_, err := c.client.Deposit(ctx, &Request{
		AccountID: accID.String(),
		Amount:    amount.Int64(),
	})
	if err != nil {
		return parseServerErr(operation, err)
	}
	return nil
}

// Withdrawal requests a withdrawal to the accounts server
func (c Client) Withdrawal(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "accounts.Client.Withdrawal"
	_, err := c.client.Withdrawal(ctx, &Request{
		AccountID: accID.String(),
		Amount:    amount.Int64(),
	})
	if err != nil {
		return parseServerErr(operation, err)
	}
	return nil
}

// ReserveCreditLimit requests a credit limit reserval to the accounts server
func (c Client) ReserveCreditLimit(ctx context.Context, accID vos.AccountID, amount vos.Money) error {
	const operation = "accounts.Client.ReserveCreditLimit"
	_, err := c.client.ReserveCreditLimit(ctx, &Request{
		AccountID: accID.String(),
		Amount:    amount.Int64(),
	})
	if err != nil {
		return parseServerErr(operation, err)
	}
	return nil
}

func parseServerErr(operation string, err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return domain.Error(operation, err)
	}
	//nolint
	switch st.Code() {
	case codes.NotFound:
		return transactions.ErrAccountNotFound
	case codes.InvalidArgument:
		switch st.Message() {
		case "err::insufficient_balance":
			return transactions.ErrInsufficientBalance
		case "err::insufficient_credit":
			return transactions.ErrInsufficientCredit
		case "err::invalid_amount":
			return transactions.ErrInvalidAmount
		}
	}

	return domain.Error(operation, err)
}
