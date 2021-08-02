// Get a current price quote for a given ticker
package cmd

import (
    "fmt"

    "github.com/piquette/finance-go/quote"
)

// Get current price data for a single ticker
func getPrice(ticker string) float64 {
    q, err := quote.Get(ticker)
    if err != nil {
        Error("Error getting ticker data from Yahoo Finance", err)
    }
    fmt.Printf("%s - $%.2f\n", ticker, q.RegularMarketPrice)
    return q.RegularMarketPrice
}

// Range over a slice of tickers, calling getPrice
func getPrices(tickers []string) {
    for _, ticker := range tickers {
        getPrice(ticker)
    }
}
