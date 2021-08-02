// Get a current price quote for a given ticker
package cmd

import (
    "fmt"

    "gitlab.com/wiggins.jonathan/plutus/api"
)

// Range over a slice of tickers, calling getPrice & printing to terminal
func getPrices(tickers []string) {
    for _, ticker := range tickers {
        price := api.GetPrice(ticker)
        fmt.Printf("%s - $%.2f\n", ticker, price)
    }
}
