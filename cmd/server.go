// Start a server to access the API over the net
package cmd

import (
    "fmt"
    "log"
    "net/http"
)

func serve() {
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

    ticker := r.URL.Query().Get("ticker")
    if ticker == "" {
        http.Error(w, "Please provide ticker=<ticker>", http.StatusBadRequest)
        return
    }

    price := getPrice(ticker)
    fmt.Fprintf(w, "Price for %s is $%g", ticker, price)
}
