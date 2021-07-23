package servers

import (
	"context"

	"github.com/fernandodr19/mybank-tx/pkg/gateway/grpc/accounts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type FakeAccountsServer struct {
	*bufconn.Listener
	accounts.UnimplementedAccountsServiceServer

	// Functions to be mocked if needed
	OnDeposit    func(ctx context.Context, req *accounts.Request) (*accounts.Response, error)
	OnWithdrawal func(ctx context.Context, req *accounts.Request) (*accounts.Response, error)
	OnReserve    func(ctx context.Context, req *accounts.Request) (*accounts.Response, error)
}

func NewFakeAccountsServer() *FakeAccountsServer {
	bufSize := 1024 * 1024
	listener := bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	s := &FakeAccountsServer{Listener: listener}
	accounts.RegisterAccountsServiceServer(grpcServer, s)
	go func() {
		grpcServer.Serve(listener)
	}()
	return s
}

func (s *FakeAccountsServer) Deposit(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
	if s.OnDeposit != nil {
		defer func() { s.OnDeposit = nil }()
		return s.OnDeposit(ctx, req)
	}
	return &accounts.Response{Success: true}, nil
}

func (s *FakeAccountsServer) Withdrawal(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
	if s.OnWithdrawal != nil {
		defer func() { s.OnWithdrawal = nil }()
		return s.OnWithdrawal(ctx, req)
	}
	return &accounts.Response{Success: true}, nil
}

func (s *FakeAccountsServer) ReserveCreditLimit(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
	if s.OnReserve != nil {
		defer func() { s.OnReserve = nil }()
		return s.OnReserve(ctx, req)
	}
	return &accounts.Response{Success: true}, nil
}
