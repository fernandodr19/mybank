package tests

import (
	"context"
	"testing"

	"github.com/fernandodr19/mybank-tx/pkg/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_TransactionSave(t *testing.T) {
	testTable := []struct {
		Name        string
		Tx          entities.Transaction
		ExpectError bool
	}{
		{
			Name: "save tx happy path",
			Tx: entities.Transaction{
				AccountID: "79924661-86b6-4160-bd9e-48f3fbc4b489",
				Operaion:  4,
				Amount:    100,
			},
		},
		{
			Name: "save negative amount tx happy path",
			Tx: entities.Transaction{
				AccountID: "79924661-86b6-4160-bd9e-48f3fbc4b489",
				Operaion:  1,
				Amount:    -100,
			},
		},
		{
			Name: "invalid acc id uuid",
			Tx: entities.Transaction{
				AccountID: "123",
				Operaion:  1,
				Amount:    -100,
			},
			ExpectError: true,
		},
		{
			Name: "fk violation, invalid operation",
			Tx: entities.Transaction{
				AccountID: "79924661-86b6-4160-bd9e-48f3fbc4b489",
				Operaion:  -1,
				Amount:    -100,
			},
			ExpectError: true,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			defer truncatePostgresTables()
			txID, err := testEnv.TxRepo.SaveTransaction(context.Background(), tt.Tx)
			if tt.ExpectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			//testing tx id (uuid) parse
			_, err = uuid.Parse(txID.String())
			assert.NoError(t, err)
		})
	}
}
