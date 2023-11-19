package currency

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertToUsDollarSuccess(t *testing.T) {
	amount := "2.86"
	exchangeRate := "1.324"
	convertAmount, err := convertToUsDollarAndRoundOff(amount, exchangeRate)
	require.Equal(t, "3.79", convertAmount)
	require.NoError(t, err)
}

func TestGetExchangeRateSuccess(t *testing.T) {
	_, err := getExchangeRate("Australia")
	require.NoError(t, err)
}

func TestGetExchangeRateFailed_InvalidCountry(t *testing.T) {
	_, err := getExchangeRate("ABC")
	require.Error(t, err)
}
