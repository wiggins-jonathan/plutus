// Parse & execute CLI commands
package cmd

import (
    "os"
    "fmt"

    "gitlab.com/wiggins.jonathan/plutus/server"

    "github.com/docopt/docopt-go"
)

var usage string = `plutus - A financial services tool.

Usage:
    plutus [options] <command> [<args>...]

Options:
    -h, --help  Print this help dialogue

Commands:
    price       Get a price quote for a space-separated list of tickers.
    rebalance   Rebalance a portfolio defined in a yaml or json file.
    server      Start a server to access the API over the net.
`

// Parses & validates args. Calls execute functions.
func ArgParse() {
    parser := &docopt.Parser{ OptionsFirst: true }
    opts, err := parser.ParseArgs(usage, nil, "")
    if err != nil {
        Error(err)
    }

    cmd     := opts["<command>"].(string)
    switch cmd {
    case "price"    : getPrices()
    case "rebalance": rebalance()
    case "server"   : server.Serve()
    default         :
        err = fmt.Errorf("%s is not a command." , cmd)
        Error(err)
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
    fmt.Printf(usage)
    os.Exit(1)
}
