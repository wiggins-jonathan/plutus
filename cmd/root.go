// Parse & instantiate CLI commands
package cmd

import (
    "os"
    "fmt"
    "path"

    "gitlab.com/wiggins.jonathan/plutus/server"
)

// Parses & validates args. Calls execute functions.
func ArgParse(args []string) {
    if len(args) < 1 { Error("Please specify a command") }

    // Execute command based on user input
    switch args[0] {
    case "server"   : server.Serve()
    case "price"    :
        if len(args) < 2 {
            Error("Please specify a ticker in which to get price data")
        }
        getPrices(args[1:])
    case "rebalance":
        if len(args) < 2 {
            Error("Please specify a file to parse for portfolio data")
        }
        rebalance(args[1])
    case "help", "-h", "--help": Usage()
    default:
        err := fmt.Sprintf("%s is not a valid command\n", args[0])
        Error(err)
    }
}

func Usage() {
    basename := path.Base(os.Args[0])
    fmt.Printf("%s - A financial services tool\n\n", basename)
    fmt.Printf("Usage:\n")
    fmt.Printf("    %s <command> [arguments]\n\n", basename)
    fmt.Printf("Commands:\n")
    fmt.Printf("    server      Start a server to access the API over the net\n")
    fmt.Printf("    price       Get a price quote for a space-separated list of tickers\n")
    fmt.Printf("    rebalance   Rebalance a portfolio defined in a yaml or json file\n")
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
    fmt.Printf("\n")
    Usage()
    os.Exit(1)
}
