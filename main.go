package main

import (
    "fmt"
    "os"

    "prc/ingest"

    _ "github.com/piquette/finance-go/quote"
)

func main() {
    file := os.Args[1]
    data := ingest.Parse(file)

    fmt.Println(data)
    for k, v := range data.Tickers {
        fmt.Println(k, v)

    //    //q, err := quote.Get(i)
    //    //if err != nil { panic(err) }
    //    //fmt.Println(q)
    }
}
