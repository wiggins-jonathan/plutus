// Rebalance a portfolio define in a yaml or json file
package cmd

import (
    "fmt"
    "sync"

    "gitlab.com/wiggins.jonathan/plutus/ingest"

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

func rebalance(file string) {
    data := ingest.FileParse(file)
    p := newPortfolio(data)
    p.getTickerData()
    p.doMath()
}

// Validate data & transform to Portfolio struct
func newPortfolio(data map[string]interface{}) *Portfolio {
    var p Portfolio
    t := make(map[string]*Ticker)
    var sumTotal, sumPercents float64
    for key, value := range data {
        if key == "addition" {
            switch value.(type) {
            case int:   // Transform to float64
                value := value.(int)
                p.Addition = float64(value)
            case float64:
                p.Addition = value.(float64)
            default:
                Error("The <addition> field must be a number")
            }
            continue
        }

        // More type assertions
        value := value.(map[string]interface{})
        c := value["current"].(float64)
        d := value["desired"].(float64)

        if c < 0 {
            err := fmt.Sprintf("The <current> field for %s must be greater than 0", key)
            Error(err)
        }

        sumTotal    += c
        sumPercents += d

        t[key] = &Ticker{
            Current: c,
            Desired: d,
        }
        p.Tickers = t
    }
    p.Total = sumTotal

    if sumPercents != 100 {
        Error("The sum of all <desired> fields must equal 100")
    }

    return &p
}

// Concurrently get ticker price from finance-go & embed in Portfolio struct
// We might want to think about just wholly embedding q into p & then creating
// multiple methods to return specific data
func (p *Portfolio) getTickerData() {
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
func (p *Portfolio) doMath() {
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
