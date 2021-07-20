// A command line library for prc
package cmd

import (
    "os"
    "fmt"
)

// Parses & validates args. Calls execute functions.
func ArgParse(args []string) {
    if len(args) < 2 { Error("Please specify a command\n") }

    // Execute based on args
    switch args[1] {
    case "server"   : ExecuteServer()
    case "price"    :
        if len(args) < 3 {
            Error("Please specify a ticker in which to get price data\n")
        }
        executePrices(args[2:])
    case "rebalance":
        if len(args) < 3 {
            Error("Please specify a file to parse for portfolio data\n")
        }
        ExecuteRebalance(args[2])
    default:
        err := fmt.Sprintf("%s is not a valid command\n", args[1])
        Error(err)
    }
}

func Usage() {
    fmt.Printf("PRC - A tool for recalculating your stock porfolio\n")
    fmt.Printf("Usage:\n")
    fmt.Printf("    prc <command> [arguments]\n")
    fmt.Printf("Commands:\n")
    fmt.Printf("    rebalance <file>    Rebalance a portfolio defined in a yaml or json file\n")
    fmt.Printf("    server              Start a server instance listening for REST calls\n")
}

// Adds color to messages printed to the command line
func Colorize(message interface{}, color string) {
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
        Colorize(message, "red")
    }
    Usage()
    os.Exit(1)
}
