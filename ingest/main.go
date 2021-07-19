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

type Ticker struct {
    Current             float64
    Desired             float64
    RegularMarketPrice  float64
}

type Portfolio struct {
    Addition    float64
    Tickers     map[string]*Ticker
    Total       float64
}

// Parse a json or yaml file
func FileParse(file string) map[string]interface{} {
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

    return data
}

// Validate data from file & transform to Portfolio struct
func NewPortfolio(data map[string]interface{}) *Portfolio {
    var p Portfolio
    t := make(map[string]*Ticker)
    var sumTotal float64
    for key, value := range data {
        if key == "addition" {
            switch value.(type) {
            case int:   // Transform to float64
                value := value.(int)
                p.Addition = float64(value)
            case float64:
                p.Addition = value.(float64)
            default:
                cmd.Usage("The <addition> field must be a number")
            }
            continue
        }

        // More type assertions
        value := value.(map[string]interface{})
        c := value["current"].(float64)
        d := value["desired"].(float64)

        sumTotal = sumTotal + c

        t[key] = &Ticker{
            Current: c,
            Desired: d,
        }
        p.Tickers = t
    }
    p.Total = sumTotal

    return &p
}

// Concurrently get ticker price from finance-go & embed in Portfolio struct
// We might want to think about just wholly embedding q into p & then creating
// multiple methods to return specific data
func (p *Portfolio) GetTickerData() {
    wg := sync.WaitGroup{}
    for ticker, _ := range p.Tickers {
        wg.Add(1)
        go func(ticker string) {
            q, err := quote.Get(ticker)
            if err != nil {
                fmt.Println("Error getting ticker data from Yahoo Finance", err)
            }

            // Assign ticker data to Portfolio struct
            p.Tickers[ticker].RegularMarketPrice = q.RegularMarketPrice

            wg.Done()
        }(ticker)
    }
    wg.Wait()
}

// Calculate the proportional number of shares to buy
func (p *Portfolio) DoMath() {
    for ticker, _ := range p.Tickers {
        // Determine the actual proportion of the portfolio for each ticker
        // as a percentage
        actualPercent:= (p.Tickers[ticker].Current / p.Total) * 100

        // Determine the difference between the actual percent that each ticker
        // represents & the desired percent we want to obtain
        percentDiff := (p.Tickers[ticker].Desired - actualPercent)

        // Determine the percent amount of the total addition we need to add or
        // subtract to reach our desired percentage of each ticker in our portfolio
        targetPercent := (percentDiff + p.Tickers[ticker].Desired)

        // Translate that difference in desired percentage into a dollar amount
        // We must check if either of these are 0
        amountToChange := (targetPercent * p.Addition) / 100

        // Giving us the # of shares to buy or sell to reach our desired percentage
        // We must check if either of these are 0
        sharesToBuy := (amountToChange / p.Tickers[ticker].RegularMarketPrice)

        // Round to two sig figs & print
        atc := fmt.Sprintf("%.2f", amountToChange)
        stb := fmt.Sprintf("%.2f", sharesToBuy)
        fmt.Printf("%v - Buy $%v or %v shares\n", ticker, atc, stb)
    }
}
