package main

import (
	"context"
	"net/http"
	"time"

	app "github.com/fernandodr19/mybank-tx/pkg"
	"github.com/fernandodr19/mybank-tx/pkg/config"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/api"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/db/postgres"
	"github.com/fernandodr19/mybank-tx/pkg/instrumentation/logger"

	"google.golang.org/grpc"
)

func main() {
	log := logger.Default()
	log.Infoln("=== My Bank API ===")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.WithError(err).Fatal("failed loading config")
	}

	ctx := context.Background()

	// Setup postgres
	dbConn, err := postgres.NewConnection(ctx, cfg.Postgres)
	if err != nil {
		log.WithError(err).Fatal("failed setting up postgres")
	}

	// gRPC
	grpcConn, err := grpc.Dial(cfg.AccountsClient.URL, grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatalln("failed connecting grpc")
	}
	defer grpcConn.Close()

	// Build app
	app, err := app.BuildApp(dbConn, grpcConn)
	if err != nil {
		log.WithError(err).Fatalln("failed to build app")
	}

	// Build API handler
	apiHandler, err := api.BuildHandler(app, cfg)
	if err != nil {
		log.WithError(err).Fatalln("failed to build api")
	}

	serveApp(apiHandler, cfg)
}

func serveApp(apiHandler http.Handler, cfg *config.Config) {
	server := &http.Server{
		Handler:      apiHandler,
		Addr:         cfg.API.Address(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Default().WithField("address", cfg.API.Address()).Info("server starting...")
	logger.Default().Fatal(server.ListenAndServe())
}
