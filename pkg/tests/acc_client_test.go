package tests

import (
	"context"
	"testing"

	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/grpc/accounts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Deposit(t *testing.T) {
	testTable := []struct {
		Name          string
		Setup         func()
		ExpectedError error
	}{
		{
			Name: "deposit happy path",
		},
		{
			Name: "account server responded acc not found",
			Setup: func() {
				testEnv.AccountsServer.OnDeposit = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return nil, status.New(codes.NotFound, "err::account_not_found").Err()
				}
			},
			ExpectedError: transactions.ErrAccountNotFound,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Setup != nil {
				tt.Setup()
			}
			err := testEnv.AccoutsClient.Deposit(context.Background(), "123", 999)
			assert.ErrorIs(t, tt.ExpectedError, err)
		})
	}
}

func Test_Withdrawal(t *testing.T) {
	testTable := []struct {
		Name          string
		Setup         func()
		ExpectedError error
	}{
		{
			Name: "withdrawal happy path",
		},
		{
			Name: "account server responded acc not found",
			Setup: func() {
				testEnv.AccountsServer.OnWithdrawal = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return nil, status.New(codes.NotFound, "err::account_not_found").Err()
				}
			},
			ExpectedError: transactions.ErrAccountNotFound,
		},
		{
			Name: "account server responded insufficient balance",
			Setup: func() {
				testEnv.AccountsServer.OnWithdrawal = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return nil, status.New(codes.InvalidArgument, "err::insufficient_balance").Err()
				}
			},
			ExpectedError: transactions.ErrInsufficientBalance,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Setup != nil {
				tt.Setup()
			}
			err := testEnv.AccoutsClient.Withdrawal(context.Background(), "123", 999)
			assert.ErrorIs(t, tt.ExpectedError, err)
		})
	}
}

func Test_Credit(t *testing.T) {
	testTable := []struct {
		Name          string
		Setup         func()
		ExpectedError error
	}{
		{
			Name: "credit happy path",
		},
		{
			Name: "account server responded acc not found",
			Setup: func() {
				testEnv.AccountsServer.OnReserve = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return nil, status.New(codes.NotFound, "err::account_not_found").Err()
				}
			},
			ExpectedError: transactions.ErrAccountNotFound,
		},
		{
			Name: "account server responded insufficient credit limit",
			Setup: func() {
				testEnv.AccountsServer.OnReserve = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return nil, status.New(codes.InvalidArgument, "err::insufficient_credit").Err()
				}
			},
			ExpectedError: transactions.ErrInsufficientCredit,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Setup != nil {
				tt.Setup()
			}
			err := testEnv.AccoutsClient.ReserveCreditLimit(context.Background(), "123", 999)
			assert.ErrorIs(t, tt.ExpectedError, err)
		})
	}
}
