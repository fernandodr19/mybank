package transactions

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fernandodr19/mybank/pkg/domain/entities/operations"
	usecase "github.com/fernandodr19/mybank/pkg/domain/usecases/transactions"
	"github.com/fernandodr19/mybank/pkg/domain/vos"
	"github.com/fernandodr19/mybank/pkg/gateway/api/middleware"
	"github.com/fernandodr19/mybank/pkg/gateway/api/responses"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ProcessTransaction(t *testing.T) {
	const (
		routePattern = "/api/v1/transactions"
		target       = "/api/v1/transactions"
	)

	request := func(body []byte) *http.Request {
		return httptest.NewRequest(http.MethodPost, target, bytes.NewReader(body))
	}

	testTable := []struct {
		Name                 string
		Handler              Handler
		Req                  TransactionRequest
		ExpectedStatusCode   int
		ExpectedErrorPayload responses.ErrorPayload
	}{
		{
			Name:    "transact happy path",
			Handler: transactHandler(nil),
			Req: TransactionRequest{
				AccountID:   "4442011c-e531-4703-b264-c3d700750f81",
				OperationID: 1,
				Amount:      1000,
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:    "acc not found",
			Handler: transactHandler(usecase.ErrAccountNotFound),
			Req: TransactionRequest{
				AccountID:   "4442011c-e531-4703-b264-c3d700750f81",
				OperationID: operations.Payment,
				Amount:      1000,
			},
			ExpectedStatusCode:   http.StatusNotFound,
			ExpectedErrorPayload: responses.ErrAccountNotFound,
		},
		{
			Name:    "acc not found",
			Handler: transactHandler(usecase.ErrAccountNotFound),
			Req: TransactionRequest{
				AccountID:   "4442011c-e531-4703-b264-c3d700750f81",
				OperationID: operations.Payment,
				Amount:      1000,
			},
			ExpectedStatusCode:   http.StatusNotFound,
			ExpectedErrorPayload: responses.ErrAccountNotFound,
		},
		{
			Name:    "insufficient balance",
			Handler: transactHandler(usecase.ErrInsufficientBalance),
			Req: TransactionRequest{
				AccountID:   "4442011c-e531-4703-b264-c3d700750f81",
				OperationID: operations.Debit,
				Amount:      999999999999999,
			},
			ExpectedStatusCode:   http.StatusUnprocessableEntity,
			ExpectedErrorPayload: responses.ErrInsufficientBalance,
		},
		{
			Name:    "insufficient credit",
			Handler: transactHandler(usecase.ErrInsufficientCredit),
			Req: TransactionRequest{
				AccountID:   "4442011c-e531-4703-b264-c3d700750f81",
				OperationID: operations.Credit,
				Amount:      999999999999999,
			},
			ExpectedStatusCode:   http.StatusUnprocessableEntity,
			ExpectedErrorPayload: responses.ErrInsufficientCredit,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			// prepare
			response := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc(routePattern, middleware.Handle(tt.Handler.Transact)).Methods(http.MethodPost)

			body, err := json.Marshal(tt.Req)
			require.NoError(t, err)

			// test
			router.ServeHTTP(response, request(body))

			//assert
			assert.Equal(t, tt.ExpectedStatusCode, response.Code)
			assert.Equal(t, "application/json", response.Header().Get("content-type"))

			if response.Code != http.StatusCreated {
				var errPayload responses.ErrorPayload
				err = json.NewDecoder(response.Body).Decode(&errPayload)
				require.NoError(t, err)
				assert.Equal(t, tt.ExpectedErrorPayload, errPayload)
			}
		})
	}
}

func transactHandler(err error) Handler {
	return Handler{
		Usecase: &usecase.TransactionsMockUsecase{
			TransactFunc: func(ctx context.Context, accID vos.AccountID, op operations.Operation, amount vos.Money) (vos.TransactionID, error) {
				return vos.TransactionID(uuid.NewString()), err
			},
		},
	}
}
