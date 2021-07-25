package main

import (
	"context"
	"net/http"
	"time"

	_ "github.com/fernandodr19/mybank-tx/docs/swagger"
	_ "github.com/joho/godotenv/autoload"

	app "github.com/fernandodr19/mybank-tx/pkg"
	"github.com/fernandodr19/mybank-tx/pkg/config"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/api"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/db/postgres"
	"github.com/fernandodr19/mybank-tx/pkg/instrumentation/logger"

	"google.golang.org/grpc"
)

// @title Swagger Mybank API
// @version 1.0
// @host localhost:3000
// @basePath /api/v1
// @schemes http
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @description Documentation Mybank API
func main() {
	log := logger.Default()
	log.Infoln("=== My Bank TX ===")

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
	app := app.BuildApp(dbConn, grpcConn)

	// Build API handler
	apiHandler := api.BuildHandler(app, cfg)

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
