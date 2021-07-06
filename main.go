package main

import (
    _ "fmt"
    "os"

    "prc/ingest"
    "prc/cmd"
)

func main() {
    validArgs := cmd.ArgParse(os.Args)
    file := validArgs[1]
    p := ingest.FileParse(file)
    p.GetTickerData()
    p.DoMath()
}
