package entities

import (
	"time"

	"github.com/fernandodr19/mybank-tx/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
)

type Transaction struct {
	ID        vos.TransactionID
	AccountID vos.AccountID
	Operaion  operations.Operation
	Amount    vos.Money
	CreatedAt time.Time
	UpdateAt  time.Time
}
