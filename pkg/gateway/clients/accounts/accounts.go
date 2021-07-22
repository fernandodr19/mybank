package accounts

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain/entities"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
)

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

func (c Client) GetAccountDetails(ctx context.Context, accID vos.AccountID) (entities.Account, error) {
	return entities.Account{}, nil
}
