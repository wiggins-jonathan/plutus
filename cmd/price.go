// Code for the price cmd
package cmd

import (
    "fmt"

    "github.com/piquette/finance-go/quote"
)

// Get concurrent price data
func getPrice(ticker string) float64 {
    q, err := quote.Get(ticker)
    if err != nil {
        Error("Error getting ticker data from Yahoo Finance", err)
    }
    fmt.Printf("%s - %.2f\n", ticker, q.RegularMarketPrice)
    return q.RegularMarketPrice
}

// Range over all tickers passed in from the CLI args
func executePrices(tickers []string) {
    for _, ticker := range tickers {
        getPrice(ticker)
    }
}
