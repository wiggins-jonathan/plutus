package main

import (
    "fmt"
    "os"
    "sync"

    "prc/ingest"
    "prc/cmd"

    "github.com/piquette/finance-go/quote"
)

func main() {
    validArgs := cmd.ArgParse(os.Args)
    file := validArgs[1]
    data := ingest.Parse(file)

    wg := sync.WaitGroup{}
    for k, _ := range data.Tickers {
        wg.Add(1)
        go func(k string) {
            q, err := quote.Get(k)
            if err != nil {
                fmt.Println("Error getting ticker data from Yahoo Finance", err)
            }
            //fmt.Printf("%+v", q) // Shows q struct fields
            fmt.Println(q.ShortName)
            fmt.Println(q.RegularMarketPrice)
            wg.Done()
        }(k)
    }
    wg.Wait()
}
