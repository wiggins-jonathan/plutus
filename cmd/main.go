// A command line library for prc
package cmd

import (
    "os"
    "fmt"
    _ "flag"
)

func Usage() {
    fmt.Printf("PRC - This is generic usage for prc\n")
    fmt.Printf("    This continues the usage\n")
}

// Checks command line args & return the args if valid, otherwise call Usage()
func ArgParse(args []string) []string {
    if len(args) < 2 {
        Colorize("Please specify a command", "red")
        Usage()
        os.Exit(1)
    }
    return args
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
    os.Exit(1)
}
