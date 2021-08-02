// Get a current price quote for a given ticker
package api

import (
    "fmt"

    "github.com/piquette/finance-go/quote"
)

// Get current price data for a single ticker
func GetPrice(ticker string) float64 {
    q, err := quote.Get(ticker)
    if err != nil {
        fmt.Printf("Error getting ticker data from Yahoo Finance", err)
    }
    return q.RegularMarketPrice
}
