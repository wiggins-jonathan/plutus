// Get a current price quote for a given ticker
package cmd

import (
    "fmt"

    "github.com/piquette/finance-go/quote"
)

// Range over a slice of tickers, calling getPrice & printing to terminal
func getPrices(tickers []string) {
    for _, ticker := range tickers {
        price := getPrice(ticker)
        fmt.Printf("%s - $%.2f\n", ticker, price)
    }
}

// Get current price data for a single ticker
func getPrice(ticker string) float64 {
    q, err := quote.Get(ticker)
    if err != nil {
        Error("Error getting ticker data from Yahoo Finance", err)
    }
    return q.RegularMarketPrice
}
