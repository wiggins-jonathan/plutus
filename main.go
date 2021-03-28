package main

import (
    "fmt"
    "os"

    "prc/ingest"
    "prc/cmd"

    "github.com/piquette/finance-go/quote"
)

func main() {
    validArgs := cmd.ArgParse(os.Args)
    file := validArgs[1]
    data := ingest.Parse(file)

    fmt.Println(data)
    for k, _ := range data.Tickers {
        q, err := quote.Get(k)
        if err != nil {
            fmt.Println("Error getting ticker data from Yahoo Finance", err)
        }
        fmt.Println(q)
    }
}
