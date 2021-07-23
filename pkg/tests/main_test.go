package tests

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"

	"github.com/fernandodr19/mybank/pkg/gateway/grpc/accounts"
	"github.com/fernandodr19/mybank/pkg/instrumentation/logger"
	"github.com/fernandodr19/mybank/pkg/tests/servers"
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
		Name string
	}{
		{
			Name: "deposit happy path",
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			err := testEnv.AccoutsClient.Deposit(context.Background(), "123", 999)
			assert.NoError(t, err)

			testEnv.AccountsServer.OnDeposit = func(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
				return nil, errors.New("asa")
			}
			err = testEnv.AccoutsClient.Deposit(context.Background(), "123", 999)
			assert.Error(t, err)

		})
	}
}
