package httpserver

import (
	"github.com/mrtechit/purchase-transaction/model"
	"net/http"
)

const (
	apiPath = "/v1/api/transaction"
)

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
