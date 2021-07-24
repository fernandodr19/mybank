package transactions

import (
	"context"
	"net/http"

	"github.com/fernandodr19/mybank-tx/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
	"github.com/fernandodr19/mybank-tx/pkg/domain/vos"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/api/middleware"

	"github.com/gorilla/mux"
)

//go:generate moq -skip-ensure -stub -out mocks.gen.go . Usecase:TransactionsMockUsecase

var _ Usecase = &transactions.Usecase{}

// Usecase of transactions
type Usecase interface {
	Transact(ctx context.Context, accID vos.AccountID, op operations.Operation, amount vos.Money) (vos.TransactionID, error)
}

// Handler handles account related requests
type Handler struct {
	Usecase
}

// NewHandler builds accounts handler
func NewHandler(public *mux.Router, admin *mux.Router, usecase Usecase) *Handler {
	h := &Handler{
		Usecase: usecase,
	}

	public.Handle("/transactions",
		middleware.Handle(h.Transact)).
		Methods(http.MethodPost)

	return h
}
