package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrtechit/purchase-transaction/currency"
	"github.com/mrtechit/purchase-transaction/model"
	"github.com/shopspring/decimal"
	"net/http"
	"time"
)

const (
	apiPath = "/v1/api/transaction"
)

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

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
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
				response := ErrorResponse{ErrorMessage: "Error decoding JSON"}
				jsonResponse, _ := json.Marshal(response)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonResponse)
				return
			}
			fmt.Println("Store request received", storeTransactionRequest)
			if !validateStoreTransactionRequest(storeTransactionRequest) {
				response := ErrorResponse{ErrorMessage: "Invalid request body"}
				jsonResponse, _ := json.Marshal(response)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(jsonResponse)
				return
			}
			apiHandler.handleStoreTrx(w, storeTransactionRequest)
		} else if r.Method == http.MethodGet {

			transactionID := r.URL.Query().Get("transaction_id")
			country := r.URL.Query().Get("country")

			if transactionID == "" || country == "" {
				response := ErrorResponse{ErrorMessage: "Missing request params"}
				jsonResponse, _ := json.Marshal(response)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(jsonResponse)
				return
			}
			fmt.Println("Retrieve request received for transactionID : ", transactionID)
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
		response := ErrorResponse{ErrorMessage: "Error encoding JSON"}
		jsonResponse, _ = json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// handleRetrieveTrx http handler of GET request which retrieves trx
func (apiHandler *ApiHandler) handleRetrieveTrx(w http.ResponseWriter, transactionID, country string) {

	trx, err := apiHandler.Db.RetrieveTrx(transactionID)
	if err != nil {
		response := ErrorResponse{ErrorMessage: "Trx not found"}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResponse)
		return
	}
	exchangeRate, err := currency.GetExchangeRate(country, trx.TransactionDate)
	if err != nil {
		response := ErrorResponse{ErrorMessage: "Error fetching exchangeRate"}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
		return
	}
	convertedAmount, err := currency.ConvertToUsDollarAndRoundOff(trx.USDollarAmount, exchangeRate)
	if err != nil {
		response := ErrorResponse{ErrorMessage: "Error converting currency"}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
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
		response := ErrorResponse{ErrorMessage: "Error encoding json"}
		jsonResponse, _ = json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func validateStoreTransactionRequest(storeTransactionRequest StoreTransactionRequest) bool {
	transactionDate := storeTransactionRequest.TransactionDate
	dateFormat := "2006-01-02" // YYYY-MM-DD
	_, err := time.Parse(dateFormat, transactionDate)
	if err != nil {
		return false
	}
	amount := storeTransactionRequest.USDollarAmount
	_, err = decimal.NewFromString(amount)
	if err != nil {
		return false
	}
	return true
}
