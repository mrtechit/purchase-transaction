package httpserver

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mrtechit/purchase-transaction/currency"
	"github.com/mrtechit/purchase-transaction/model"
	"net/http"
)

const (
	apiPath = "/v1/api/transaction"
)

type RetrieveTransactionRequest struct {
	TransactionID string `json:"transaction_id"`
	Country       string `json:"country"`
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
			var storeTransactionRequest StoreTransactionRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&storeTransactionRequest)
			if err != nil {
				http.Error(w, "Error decoding JSON", http.StatusBadRequest)
				return
			}
			apiHandler.handleStoreTrx(w, storeTransactionRequest)
		} else if r.Method == http.MethodGet {

			transactionID := r.URL.Query().Get("transaction_id")
			country := r.URL.Query().Get("country")

			if transactionID == "" || country == "" {
				http.Error(w, "Missing request params", http.StatusBadRequest)
				return
			}

			apiHandler.handleRetrieveTrx(w, transactionID, country)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}

// handleStoreTrx http handler for POST request which stores the trx
func (apiHandler *ApiHandler) handleStoreTrx(w http.ResponseWriter, storeTransactionRequest StoreTransactionRequest) {

	transactionID := uuid.New().String()

	storeTrx := &model.StoreTransaction{
		TransactionID:   transactionID,
		Description:     storeTransactionRequest.Description,
		TransactionDate: storeTransactionRequest.TransactionDate,
		USDollarAmount:  storeTransactionRequest.USDollarAmount,
	}

	err := apiHandler.Db.StoreTrx(storeTrx)
	response := StoreTransactionResponse{TransactionID: transactionID}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)

}

// handleRetrieveTrx http handler of GET request which retrieves trx
func (apiHandler *ApiHandler) handleRetrieveTrx(w http.ResponseWriter, transactionID, country string) {

	trx, err := apiHandler.Db.RetrieveTrx(transactionID)
	if err != nil {
		http.Error(w, "Error fetching trx", http.StatusInternalServerError)
		return
	}
	exchangeRate, err := currency.GetExchangeRate(country, trx.TransactionDate)
	if err != nil {
		http.Error(w, "Error fetching exchangeRate", http.StatusInternalServerError)
		return
	}
	convertedAmount, err := currency.ConvertToUsDollarAndRoundOff(trx.USDollarAmount, exchangeRate)
	if err != nil {
		http.Error(w, "Error converting currency", http.StatusInternalServerError)
		return
	}
	response := RetrieveTransactionResponse{
		TransactionID:   trx.TransactionID,
		Description:     trx.Description,
		TransactionDate: trx.TransactionDate,
		USDollarAmount:  trx.USDollarAmount,
		ExchangeRate:    exchangeRate,
		ConvertedAmount: convertedAmount,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}
