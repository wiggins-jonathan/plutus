package ingest

import (
    "io/ioutil"
    "path/filepath"
    "fmt"
    "encoding/json"
    "os"
    "sync"

    "prc/cmd"

    "gopkg.in/yaml.v3"
    "github.com/piquette/finance-go/quote"
)

type Portfolio struct {
    Total   float64
    Tickers map[string]struct {
        Current float64
        Desired int
    }
}

// Parse a json or yaml file for data & return Portfolio struct
func FileParse(file string) *Portfolio {
    fileData, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("error reading", file, err)
        os.Exit(1)
    }

    // First load the data to map[string]interface{}.
    // This is done because we don't know how many Tickers there will be.
    var data map[string]interface{}
    extension := filepath.Ext(file)
    switch extension {
    case ".yml", ".yaml"    : err = yaml.Unmarshal(fileData, &data)
    default                 : err = json.Unmarshal(fileData, &data)
    }

    if err != nil {
        fmt.Println("Error unmarshalling", file, err)
        os.Exit(1)
    }

    // Perform type assertions on our data & transform to Portfolio struct
    var p Portfolio
    t := make(map[string]struct {
        Current float64
        Desired int
    })
    for key, value := range data {
        if key == "total" {
            switch value.(type) {
            case int:   // Transform to float64
                value := value.(int)
                p.Total = float64(value)
            case float64:
                p.Total = value.(float64)
            default:
                cmd.Usage("The <total> field must be a number")
            }
            continue
        }

        // More type assertions
        value := value.(map[string]interface{})
        c := value["current"].(float64)
        d := value["desired"].(int)

        // Assign inner map to the p.Tickers anonymous struct
        t[key] = struct {
            Current float64
            Desired int
        }{
            Current: c,
            Desired: d,
        }
        p.Tickers = t
    }

    return &p
}

// Concurrently get ticker data from finance-go & embed in Portfolio struct
func GetTickerData(p *Portfolio) {
    wg := sync.WaitGroup{}
    for ticker, _ := range p.Tickers {
        wg.Add(1)
        go func(ticker string) {
            q, err := quote.Get(ticker)
            if err != nil {
                fmt.Println("Error getting ticker data from Yahoo Finance", err)
            }
            fmt.Printf("%+v\n\n", q) // Shows q struct fields
            //fmt.Println(q.Symbol, q.ShortName)
            //fmt.Println(q.RegularMarketPrice)
            //fmt.Printf("%T\n", q.RegularMarketPrice)
            wg.Done()
        }(ticker)
    }
    wg.Wait()
}
