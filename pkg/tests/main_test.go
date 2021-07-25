package tests

import (
	"context"
	"fmt"
	"net"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	app "github.com/fernandodr19/mybank-tx/pkg"
	"github.com/fernandodr19/mybank-tx/pkg/config"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/api"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/db/postgres"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/grpc/accounts"
	"github.com/fernandodr19/mybank-tx/pkg/instrumentation/logger"
	"github.com/fernandodr19/mybank-tx/pkg/tests/servers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/testcontainers/testcontainers-go"
	"google.golang.org/grpc"
)

type _testEnv struct {
	// Server
	Server *httptest.Server

	// 3rd party fake Servers
	AccountsServer *servers.FakeAccountsServer

	// Clients
	AccoutsClient *accounts.Client

	// DB
	Conn   *pgx.Conn
	TxRepo *postgres.TransactionsRepository
}

var testEnv _testEnv

func TestMain(m *testing.M) {
	teardown := setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func setup() func() {
	log := logger.Default()
	log.Info("setting up integration tests env")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.WithError(err).Fatal("failed loading config")
	}

	err = setupDockerTest()
	if err != nil {
		log.WithError(err).Fatal("failed setting up docker")
	}

	// Setup postgres
	cfg.Postgres.DBName = "test"
	cfg.Postgres.Port = "5436"
	dbConn, err := setupPostgresTest(cfg.Postgres)
	if err != nil {
		log.WithError(err).Fatal("failed setting up postgres")
	}

	testEnv.Conn = dbConn
	testEnv.TxRepo = postgres.NewTransactionRepository(dbConn)

	testEnv.AccountsServer = servers.NewFakeAccountsServer()

	grpcConn, err := grpc.Dial(cfg.AccountsClient.URL, grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return testEnv.AccountsServer.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatalln("failed connecting grpc")
	}

	testEnv.AccoutsClient = accounts.NewClient(grpcConn)

	app := app.BuildApp(dbConn, grpcConn)

	apiHandler := api.BuildHandler(app, cfg)

	testEnv.Server = httptest.NewServer(apiHandler)

	return func() {}
}

func setupDockerTest() error {
	running, err := isDockerRunning([]string{
		"pg-test",
	})
	if err != nil {
		return err
	}

	if running {
		logger.Default().Infoln("necessary containers already running...")
		return nil
	}

	compose := testcontainers.NewLocalDockerCompose(
		[]string{"./docker-compose.yml"},
		strings.ToLower(uuid.New().String()),
	)
	execErr := compose.WithCommand([]string{"up", "-d"}).Invoke()
	if execErr.Error != nil {
		return execErr.Error
	}
	return nil
}

func isDockerRunning(expectedImages []string) (bool, error) {
	stdout, err := exec.Command("docker", "ps").Output()
	if err != nil {
		return false, err
	}

	ps := string(stdout)
	if err != nil {
		return false, err
	}

	running := true
	for _, image := range expectedImages {
		if !strings.Contains(ps, image) {
			running = false
			break
		}
	}
	return running, nil
}

func setupPostgresTest(cfg config.Postgres) (*pgx.Conn, error) {
	done := make(chan bool, 1)
	var dbConn *pgx.Conn
	var err error

	// tries to connect within 5 seconds timeout
	go func() {
		for {
			dbConn, err = postgres.NewConnection(context.Background(), cfg)
			if err != nil {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			break
		}
		close(done)
	}()

	select {
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("timed out trying to set up postgres: %w", err)
	case <-done:
	}

	return dbConn, nil
}

func truncatePostgresTables() {
	testEnv.Conn.Exec(context.Background(),
		`TRUNCATE TABLE 
			transactions
		CASCADE`,
	)
}
