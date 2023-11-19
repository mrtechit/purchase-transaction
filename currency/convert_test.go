package currency

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertToUsDollarSuccess(t *testing.T) {
	amount := "2.86"
	exchangeRate := "1.324"
	convertAmount, err := ConvertToUsDollarAndRoundOff(amount, exchangeRate)
	require.Equal(t, "3.79", convertAmount)
	require.NoError(t, err)
}

func TestGetExchangeRateSuccess(t *testing.T) {
	_, err := GetExchangeRate("Australia", "2023-10-12")
	require.NoError(t, err)
}

func TestGetExchangeRateFailed_InvalidCountry(t *testing.T) {
	_, err := GetExchangeRate("ABC", "2025-10-12")
	require.Error(t, err)
}
