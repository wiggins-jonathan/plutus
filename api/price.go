// Get a current price quote for a given ticker
package api

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

// Get current price data for a single ticker
func GetPrice(ticker string) (float64, error) {
	q, err := quote.Get(ticker)
	if err != nil {
		return 0, err
	}

	if q == nil {
		return 0, fmt.Errorf("%s is an invalid ticker", ticker)
	}

	return q.RegularMarketPrice, nil
}
