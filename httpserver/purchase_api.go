package httpserver

import (
	"github.com/mrtechit/purchase-transaction/model"
	"net/http"
)

type ApiHandler struct {
	Db TransactionManager
}

func NewApiHandler(transactionManager TransactionManager) *ApiHandler {
	return &ApiHandler{Db: transactionManager}
}

type TransactionManager interface {
	StoreTrx(trx *model.StoreTransaction) error
}

func (apiHandler *ApiHandler) Handler() {

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			apiHandler.Db.StoreTrx(nil)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
