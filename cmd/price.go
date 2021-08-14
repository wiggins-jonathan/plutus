// Get a current price quote for a given ticker
package cmd

import (
    "fmt"

    "gitlab.com/wiggins.jonathan/plutus/api"

    "github.com/docopt/docopt-go"
)

// Range over a slice of tickers, calling getPrice & printing to terminal
func getPrices() {
    usage := `plutus price - Get a price quote for a space-separated list of tickers

Usage:
    plutus price <tickers>...
`
    args, err := docopt.ParseDoc(usage)
    if err != nil {
        Error(err)
    }

    tickers := args["<tickers>"].([]string)
    for _, ticker := range tickers {
        price := api.GetPrice(ticker)
        fmt.Printf("%s - $%.2f\n", ticker, price)
    }
}
