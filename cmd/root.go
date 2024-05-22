// Entrypoint for CLI + utility functions
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "development"
	debug   bool
)

var rootCmd = &cobra.Command{
	Use:     "plutus",
	Short:   "The financial services tool",
	Long:    "plutus - The Financial Services Tool",
	Version: version, // overriden by ldflags at build time
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			fmt.Println("Debug mode enabled!")
		}
	},
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.PersistentFlags().BoolVarP(
		&debug, "debug", "d", false, "Enable Debug mode",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// Adds color to messages printed to the command line
func colorize(message interface{}, color string) {
	switch color {
	case "red":
		color = "\033[31m"
	case "green":
		color = "\033[32m"
	case "yellow":
		color = "\033[33m"
	case "blue":
		color = "\033[34m"
	case "purple":
		color = "\033[35m"
	case "cyan":
		color = "\033[36m"
	case "gray":
		color = "\033[37m"
	case "white":
		color = "\033[97m"
	}
	reset := "\033[0m"

	fmt.Printf("%v%v%v\n", color, message, reset)
}

// Prints a red error message to the command line & exits
func Error(messages ...interface{}) {
	for _, message := range messages {
		colorize(message, "red")
	}
	os.Exit(1)
}
