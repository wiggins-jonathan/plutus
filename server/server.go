// Start a server to access the API over the net
package server

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "gitlab.com/wiggins.jonathan/plutus/api"
)

func Serve() {
    http.HandleFunc("/p", priceHandler)
    fmt.Println("Server listening on port 5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

// Take in the ticker & call GetPrice()
func priceHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/p" {
        http.NotFound(w, r)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        return
    }

    tickers := r.URL.Query().Get("tickers")
    if tickers == "" {
        http.Error(w, "Please provide tickers=<ticker1>,<tickerN>", http.StatusBadRequest)
        return
    }

    for _, ticker := range strings.Split(tickers, ",") {
        price, err := api.GetPrice(ticker)
        if err != nil {
            fmt.Fprintf(w, "Error retrieving price for %s\n", ticker)
        }
        fmt.Fprintf(w, "Price for %s is $%g\n", ticker, price)
    }
}
