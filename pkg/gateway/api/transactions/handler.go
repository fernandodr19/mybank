package transactions

import (
	"net/http"

	"github.com/fernandodr19/mybank-tx/pkg/domain/usecases/transactions"
	"github.com/fernandodr19/mybank-tx/pkg/gateway/api/middleware"

	"github.com/gorilla/mux"
)

// Handler handles account related requests
type Handler struct {
	Usecase transactions.Usecase
}

// NewHandler builds accounts handler
func NewHandler(public *mux.Router, admin *mux.Router, usecase transactions.Usecase) *Handler {
	h := &Handler{
		Usecase: usecase,
	}

	public.Handle("/transactions",
		middleware.Handle(h.Transact)).
		Methods(http.MethodPost)

	return h
}
