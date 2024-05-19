// Get a current price quote for a given ticker
package api

import (
	"fmt"

	yf "github.com/shoenig/yahoo-finance"
)

// Get current price data for a single ticker
func GetPrice(ticker string) (float64, error) {
	client := yf.New(nil)
	data, err := client.Lookup(ticker)
	if err != nil {
		return 0, fmt.Errorf("Could not obtain data from yahoo finance: %w", err)
	}

	price := data.Price()
	if price <= 0 {
		return 0, fmt.Errorf("No price data found for %s: %w", ticker, err)
	}

	return price, nil
}
