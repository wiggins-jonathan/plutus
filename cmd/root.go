// Entrypoint for CLI + utility functions
package cmd

import (
    "os"
    "fmt"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use     : "plutus",
    Short   : "A financial services tool",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

// Adds color to messages printed to the command line
func colorize(message interface{}, color string) {
    switch color {
    case "red"      : color = "\033[31m"
    case "green"    : color = "\033[32m"
    case "yellow"   : color = "\033[33m"
    case "blue"     : color = "\033[34m"
    case "purple"   : color = "\033[35m"
    case "cyan"     : color = "\033[36m"
    case "gray"     : color = "\033[37m"
    case "white"    : color = "\033[97m"
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
