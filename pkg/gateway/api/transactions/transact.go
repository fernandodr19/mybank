package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/fernandodr19/mybank/pkg/domain"
	"github.com/fernandodr19/mybank/pkg/domain/entities/operations"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
	"github.com/fernandodr19/mybank/pkg/gateway/api/responses"
)

// Transact executes a transaction
// @Summary Process a transaction
// @Description Process a transaction for a given account
// @Tags Transactions
// @Param Body body TransactionRequest true "Body"
// @Accept json
// @Produce json
// @Success 200 {object} TransactionResponse
// @Failure 400 "Could not parse request"
// @Failure 404 "Account not found"
// @Failure 422 "Could not process transaction due to lack of balance or available credit"
// @Failure 500 "Internal server error"
// @Router /accounts [post]
func (h Handler) Transact(r *http.Request) responses.Response {
	operation := "transactions.Handler.Transact"

	ctx := r.Context()
	var body TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return responses.BadRequest(domain.Error(operation, err), responses.ErrInvalidBody)
	}

	txID, err := h.Usecase.Transact(ctx, body.AccountID, body.OperationID, body.Amount)
	if err != nil {
		return responses.ErrorResponse(domain.Error(operation, err))
	}

	return responses.OK(TransactionResponse{
		TransactionID: txID,
	})
}

// TransactionRequest payload
type TransactionRequest struct {
	AccountID   vos.AccountID        `json:"account_id"`
	OperationID operations.Operation `json:"operation_type_id"`
	Amount      vos.Money            `json:"amount"`
}

// TransactionResponse payload
type TransactionResponse struct {
	TransactionID vos.TransactionID `json:"transaction_id"`
}
