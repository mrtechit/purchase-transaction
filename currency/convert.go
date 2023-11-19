package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"time"
)

const (
	TreasuryUrl    = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=record_date,exchange_rate,country&filter=record_date:gte:"
	sortByDate     = "&sort=-record_date"
	countryEq      = ",country:eq:"
	roundOffPlaces = 2
)

type ExchangeRateResponse struct {
	Data []struct {
		ExchangeRate string `json:"exchange_rate"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// getExchangeRate gets the exchange rate from API using default paginated parameters of 1,100
func getExchangeRate(country string) (string, error) {
	currentTime := time.Now()
	sixMonthsAgo := currentTime.AddDate(0, -6, 0)
	dateFormat := "2006-01-02" // YYYY-MM-DD
	formattedDate := sixMonthsAgo.Format(dateFormat)

	TreasuryUrl := TreasuryUrl + formattedDate + countryEq + country + sortByDate

	response, err := http.Get(TreasuryUrl)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return "", err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	var exchangeRateResponse ExchangeRateResponse
	err = json.Unmarshal(body, &exchangeRateResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "", err
	}

	if exchangeRateResponse.Meta.Count == 0 {
		return "", errors.New("no record found")
	}

	fmt.Println("Response from the API:")
	fmt.Println(string(body))

	return exchangeRateResponse.Data[0].ExchangeRate, nil
}

func convertToUsDollarAndRoundOff(amount, exchangeRate string) (string, error) {
	amountPrecise, err := decimal.NewFromString(amount)
	if err != nil {
		return "", err
	}
	exchangeRatePrecise, err := decimal.NewFromString(exchangeRate)
	if err != nil {
		return "", err
	}
	convertedAmount := amountPrecise.Mul(exchangeRatePrecise)
	convertedAmountAndRounded := convertedAmount.Round(roundOffPlaces)
	return convertedAmountAndRounded.String(), nil
}
