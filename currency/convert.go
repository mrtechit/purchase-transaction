package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"strings"
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

// GetExchangeRate gets the exchange rate from API using default paginated parameters of 1,100
func GetExchangeRate(country, date string) (string, error) {

	dateFormat := "2006-01-02" // YYYY-MM-DD
	parsedTime, err := time.Parse(dateFormat, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return "", err
	}
	sixMonthsAgo := parsedTime.AddDate(0, -6, 0)
	formattedDate := sixMonthsAgo.Format(dateFormat)

	// Capitalize the first character
	capitalizedCountryString := capitalizeFirst(country)

	TreasuryUrl := TreasuryUrl + formattedDate + countryEq + capitalizedCountryString + sortByDate
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

func ConvertToUsDollarAndRoundOff(amount, exchangeRate string) (string, error) {
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

func capitalizeFirst(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
