package main

import (
    "fmt"
    "os"

    "prc/ingest"

    _ "github.com/piquette/finance-go/quote"
)

func main() {
    file        := os.Args[1]
    yamlData    := ingest.Parse(file)
    fmt.Println(yamlData)

    //for _, i := range yamlData.Ticker {
    //    fmt.Println(i)

        //q, err := quote.Get(i)
        //if err != nil { panic(err) }
        //fmt.Println(q)
    //}
}
