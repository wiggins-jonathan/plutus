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
    data := ingest.FileParse(file)
    p := ingest.NewPortfolio(data)
    p.GetTickerData()
    p.DoMath()
}
