package httpserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateStoreTransactionRequest(t *testing.T) {

	storeTrxRequestValid := StoreTransactionRequest{
		Description:     "description",
		TransactionDate: "2022-10-13",
		USDollarAmount:  "1.46",
	}

	storeTrxRequestInvalidDate := StoreTransactionRequest{
		Description:     "description",
		TransactionDate: "10-10-13",
		USDollarAmount:  "1.46",
	}

	storeTrxRequestInvalidAmount := StoreTransactionRequest{
		Description:     "description",
		TransactionDate: "10-10-13",
		USDollarAmount:  "abc",
	}
	tests := []struct {
		name      string
		arg       StoreTransactionRequest
		assertion assert.BoolAssertionFunc
	}{
		{"invalid date", storeTrxRequestInvalidDate, assert.False},
		{"valid request", storeTrxRequestValid, assert.True},
		{"invalid amount", storeTrxRequestInvalidAmount, assert.False},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, validateStoreTransactionRequest(tt.arg))
		})
	}

}
