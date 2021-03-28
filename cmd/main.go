package cmd

import (
    "os"
    "fmt"
    _ "flag"
)

func ArgParse(args []string) []string {
    if len(args) < 2 {
        fmt.Println("You need at least one arg")
        os.Exit(1)
    }
    return args
}
