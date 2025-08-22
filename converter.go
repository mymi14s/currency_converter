package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Converter handles currency conversion using ExchangeRate API.
type Converter struct {
	API_KEY string
}

// Convert converts an amount from one currency to another.
// Returns a map with cleaned data and the amount in words.
func (c *Converter) Convert(fromCurrency string, amount float64, toCurrency string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/", c.API_KEY)
	url := fmt.Sprintf("%spair/%s/%s/%v", endpoint, fromCurrency, toCurrency, amount)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var convertedData map[string]interface{}
	if err := json.Unmarshal(body, &convertedData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	convertedData["from_amount"] = amount
	cleanedData := CleanData(convertedData)
	atw := AmountToWords(cleanedData["to_amount"].(float64), cleanedData["to_currency"].(string))
	convertedData["amount_in_word"] = atw["amount_in_word"]

	return convertedData, nil
}
