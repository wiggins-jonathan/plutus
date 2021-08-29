// Get a current price quote for a given ticker
package api

import (
    "github.com/piquette/finance-go/quote"
)

// Get current price data for a single ticker
func GetPrice(ticker string) (float64, error) {
    q, err := quote.Get(ticker)
    if err != nil {
        return 0, err
    }
    return q.RegularMarketPrice, nil
}
