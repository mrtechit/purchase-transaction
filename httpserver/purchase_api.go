package httpserver

import (
	"github.com/mrtechit/purchase-transaction/model"
	"net/http"
)

const (
	apiPath = "/v1/api/transaction"
)

type RetrieveTransactionRequest struct {
	TransactionID string `json:"transaction_id"`
	Currency      string `json:"currency"`
}

type RetrieveTransactionResponse struct {
	TransactionID   string `json:"transaction_id"`
	Description     string `json:"description"`
	TransactionDate string `json:"transaction_date"`
	USDollarAmount  string `json:"us_dollar_amount"`
	ExchangeRate    string `json:"exchange_rate"`
	ConvertedAmount string `json:"converted_amount"`
}

type StoreTransactionRequest struct {
	Description     string `json:"description"`
	TransactionDate string `json:"transaction_date"`
	USDollarAmount  string `json:"us_dollar_amount"`
}

type StoreTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
}

type ApiHandler struct {
	Db TransactionManager
}

func NewApiHandler(transactionManager TransactionManager) *ApiHandler {
	return &ApiHandler{Db: transactionManager}
}

type TransactionManager interface {
	StoreTrx(trx *model.StoreTransaction) error
	RetrieveTrx(transactionID string) (*model.StoreTransaction, error)
}

func (apiHandler *ApiHandler) Handler() {

	http.HandleFunc(apiPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleStoreTrx()
		} else if r.Method == http.MethodGet {
			handleRetrieveTrx()
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}

func handleStoreTrx() {

}

func handleRetrieveTrx() {

}
