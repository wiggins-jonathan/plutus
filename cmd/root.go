package cmd

import (
    "os"
    "fmt"
    "path"
    "path/filepath"
    "strings"
)

// Parses & validates args. Calls execute functions.
func ArgParse(args []string) {
    if len(args) < 1 { Error("Please specify a command\n") }

    commands, err := getCommandFiles("cmd")
    if err != nil {
        Error("Cannnot detect commands in /cmd directory")
    }
    _, found := func(slice []string, val string) (int, bool) {
        for i, item := range slice {
            if item == val {
                return i, true
            }
        }
        return -1, false
    }(commands, args[0])
    if !found {
        err := fmt.Sprintf("%s is not a valid command\n", args[0])
        Error(err)
    }


    // Execute based on args
    switch args[0] {
    case "server"   : serve()
    case "price"    :
        if len(args) < 2 {
            Error("Please specify a ticker in which to get price data\n")
        }
        getPrices(args[1:])
    case "rebalance":
        if len(args) < 2 {
            Error("Please specify a file to parse for portfolio data\n")
        }
        rebalance(args[1])
    case "help", "-h", "--help": Usage()
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
    Usage()
    os.Exit(1)
}

// Walk a filepath & return slice of all command files
func getCommandFiles(dir string) ([]string, error) {
    var paths []string
    err := filepath.WalkDir(dir, func(file string, info os.DirEntry, err error) error {
        if strings.HasPrefix(info.Name(), ".") {
            if info.IsDir() {
                return filepath.SkipDir
            }
            return err
        }

        if !info.IsDir() {
            if file == "cmd/root.go" {
                return err
            }
            file := path.Base(file)
            file = strings.TrimSuffix(file, filepath.Ext(file))
            paths = append(paths, file)
        }
        return err

    })
    return paths, err
}
