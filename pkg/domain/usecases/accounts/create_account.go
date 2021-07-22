package accounts

import (
	"context"

	"github.com/fernandodr19/mybank/pkg/domain/vos"
)

// CreateAccount creates a new account for a given document number
func (u Usecase) CreateAccount(ctx context.Context, doc vos.Document) (vos.AccountID, error) {
	return "", nil
}
