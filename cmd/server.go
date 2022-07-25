package cmd

import (
	"gitlab.com/wiggins.jonathan/plutus/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Serve the finance-go API",
	Long:    "Serve the piquette/finance-go API over http",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve()
	},
}
