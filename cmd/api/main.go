package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fernandodr19/mybank/pkg/gateway/grpc/accounts"
	"github.com/fernandodr19/mybank/pkg/instrumentation/logger"

	"google.golang.org/grpc"
)

func main() {
	log := logger.Default()
	log.Infoln("=== My Bank API ===")

	// GRPC
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Errorln("failed connecting grpc")
	}
	defer conn.Close()
	accClient := accounts.NewClient(conn)

	acc, err := accClient.GetAccountDetails(context.Background(), "")
	if err != nil {
		log.WithError(err).Fatal("failed getting acc")
	}
	log.Infoln("UUHUUUL")
	log.Infoln(acc)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	signal.Stop(signals)
}
