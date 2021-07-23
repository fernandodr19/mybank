package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fernandodr19/mybank-tx/pkg/gateway/grpc/accounts"
	"github.com/fernandodr19/mybank-tx/pkg/instrumentation/logger"

	"google.golang.org/grpc"
)

func main() {
	log := logger.Default()
	log.Infoln("=== My Bank API ===")

	// GRPC
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatalln("failed connecting grpc")
	}
	defer conn.Close()
	accClient := accounts.NewClient(conn)

	err = accClient.Deposit(context.Background(), "", 11212)
	if err != nil {
		log.WithError(err).Fatal("failed deposit")
	}
	err = accClient.Withdrawal(context.Background(), "", 11212)
	if err != nil {
		log.WithError(err).Fatal("failed withdraw")
	}
	err = accClient.ReserveCreditLimit(context.Background(), "", 11212)
	if err != nil {
		log.WithError(err).Fatal("failed reserve")
	}
	log.Infoln("UUHUUUL")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	signal.Stop(signals)
}
