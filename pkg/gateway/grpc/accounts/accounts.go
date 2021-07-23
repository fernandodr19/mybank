package accounts

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain"
	"github.com/fernandodr19/mybank/pkg/domain/entities"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
	"github.com/fernandodr19/mybank/pkg/instrumentation/logger"
	"google.golang.org/grpc"
)

type Client struct {
	client AccountsServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: NewAccountsServiceClient(conn),
	}

}

func (c Client) GetAccountDetails(ctx context.Context, accID vos.AccountID) (entities.Account, error) {
	const operation = "accounts.Client.GetAccountDetails"
	reply, err := c.client.Deposit(ctx, &DepositRequest{
		AccountID: "acc 123",
		Amount:    "amount 999",
	})
	if err != nil {
		return entities.Account{}, domain.Error(operation, err)
	}
	logger.Default().WithField("errorCode", reply.ErrorCode)
	return entities.Account{}, nil
}
