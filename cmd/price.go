// Get a current price quote for a given ticker
package cmd

import (
	"fmt"

	"github.com/wiggins-jonathan/plutus/api"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(priceCmd)
}

var priceCmd = &cobra.Command{
	Use:     "price <tickers...>",
	Aliases: []string{"p", "prices"},
	Short:   "Get a price quote",
	Long:    "Get a price quote for a space-separated list of tickers",
	Example: "plutus price swagx swisx",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, ticker := range args {
			price, err := api.GetPrice(ticker)
			if err != nil {
				Error(err)
			}

			fmt.Printf("%s - $%.2f\n", ticker, price)
		}
	},
}
