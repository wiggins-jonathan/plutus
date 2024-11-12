package cmd

import (
	"fmt"
	"os"

	"github.com/wiggins-jonathan/plutus/ingest"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(parseCmd)
}

var parseCmd = &cobra.Command{
	Use:   "parse [directory]",
	Short: "Parse beancount files",
	Long:  "Parse beancount files recursively",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd() // default to current working directory
		if err != nil {
			return fmt.Errorf("Errog getting cwd: %w", err)
		}
		if len(args) > 0 { // change default if passed in from CLI
			dir = args[0]
		}

		ingest.ReadFiles(dir, ".bean")
		return nil
	},
}
