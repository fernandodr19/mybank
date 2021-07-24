package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/fernandodr19/mybank-tx/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/api/transactions"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/grpc/accounts"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Transact(t *testing.T) {
	testTable := []struct {
		Name               string
		Req                transactions.TransactionRequest
		Setup              func()
		ExpectedStatusCode int
	}{
		{
			Name: "deposit happy path",
			Req: transactions.TransactionRequest{
				AccountID:   "899c24d2-31a8-44c7-ab7b-8f681aa42e0a",
				OperationID: operations.Debit,
				Amount:      10,
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name: "invalid acc id",
			Req: transactions.TransactionRequest{
				AccountID:   "123",
				OperationID: operations.Debit,
				Amount:      10,
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name: "invalid amount",
			Req: transactions.TransactionRequest{
				AccountID:   "dd84836d-d627-4dff-9525-d7303fbec2fb",
				OperationID: operations.Debit,
				Amount:      -10,
			},
			ExpectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			Name: "invalid operation",
			Req: transactions.TransactionRequest{
				AccountID:   "dd84836d-d627-4dff-9525-d7303fbec2fb",
				OperationID: 0,
				Amount:      10,
			},
			ExpectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			Name: "deposit acc server unknown error",
			Req: transactions.TransactionRequest{
				AccountID:   "dd84836d-d627-4dff-9525-d7303fbec2fb",
				OperationID: operations.Payment,
				Amount:      10,
			},
			Setup: func() {
				testEnv.AccountsServer.OnDeposit = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return &accounts.Response{}, status.New(codes.Unknown, "unknown error").Err()
				}
			},
			ExpectedStatusCode: http.StatusInternalServerError,
		},
		{
			Name: "withdrawal insufficient balance",
			Req: transactions.TransactionRequest{
				AccountID:   "dd84836d-d627-4dff-9525-d7303fbec2fb",
				OperationID: operations.Withdrawal,
				Amount:      10,
			},
			Setup: func() {
				testEnv.AccountsServer.OnWithdrawal = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return &accounts.Response{}, status.New(codes.InvalidArgument, "err::insufficient_balance").Err()
				}
			},
			ExpectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			Name: "insufficient credit",
			Req: transactions.TransactionRequest{
				AccountID:   "dd84836d-d627-4dff-9525-d7303fbec2fb",
				OperationID: operations.Credit,
				Amount:      10,
			},
			Setup: func() {
				testEnv.AccountsServer.OnReserve = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return &accounts.Response{}, status.New(codes.InvalidArgument, "err::insufficient_credit").Err()
				}
			},
			ExpectedStatusCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			defer truncatePostgresTables()

			// prepare
			if tt.Setup != nil {
				tt.Setup()
			}

			target := testEnv.Server.URL + "/api/v1/transactions"
			body, err := json.Marshal(tt.Req)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(body))
			require.NoError(t, err)

			// test
			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			// assert
			require.Equal(t, tt.ExpectedStatusCode, resp.StatusCode)

		})
	}
}
