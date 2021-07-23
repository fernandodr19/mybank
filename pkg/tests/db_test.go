package tests

import (
	"context"
	"testing"

	"github.com/fernandodr19/mybank-tx/pkg/domain/entities"
	"github.com/stretchr/testify/assert"
)

func Test_TransactionSave(t *testing.T) {
	testTable := []struct {
		Name string
		Tx   entities.Transaction
	}{
		{
			Name: "save tx happy path",
			Tx: entities.Transaction{
				AccountID: "79924661-86b6-4160-bd9e-48f3fbc4b489",
				Operaion:  1,
				Amount:    100,
			},
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			txID, err := testEnv.TxRepo.SaveTransaction(context.Background(), tt.Tx)
			assert.NoError(t, err)
			assert.NotEmpty(t, txID)
		})
	}
}
