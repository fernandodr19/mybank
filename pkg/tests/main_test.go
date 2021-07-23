package tests

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"

	"github.com/fernandodr19/mybank-tx/pkg/gateway/grpc/accounts"
	"github.com/fernandodr19/mybank-tx/pkg/instrumentation/logger"
	"github.com/fernandodr19/mybank-tx/pkg/tests/servers"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type _testEnv struct {
	AccountsServer *servers.FakeAccountsServer
	AccoutsClient  *accounts.Client
}

var testEnv _testEnv

func TestMain(m *testing.M) {
	teardown := setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func setup() func() {
	testEnv.AccountsServer = servers.NewFakeAccountsServer()

	grpcConn, err := grpc.Dial(":9000", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return testEnv.AccountsServer.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		logger.Default().WithError(err).Fatalln("failed connecting grpc")
	}

	testEnv.AccoutsClient = accounts.NewClient(grpcConn)

	return func() {}
}

func Test_Deposit(t *testing.T) {
	testTable := []struct {
		Name        string
		Setup       func()
		ExpectError bool
	}{
		{
			Name: "deposit happy path",
		},
		{
			Name: "deposit error",
			Setup: func() {
				testEnv.AccountsServer.OnDeposit = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
					return nil, errors.New("deposit unknown error")
				}
			},
			ExpectError: true,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Setup != nil {
				tt.Setup()
			}
			err := testEnv.AccoutsClient.Deposit(context.Background(), "123", 999)
			assert.Equal(t, tt.ExpectError, err != nil)
		})
	}
}
