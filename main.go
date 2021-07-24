package main

import (
    "os"

    "gitlab.com/wiggins.jonathan/plutus/cmd"
)

func main() {
    cmd.ArgParse(os.Args)
}
