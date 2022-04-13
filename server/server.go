// Start a server to access the API over the net
package server

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "encoding/json"

    "gitlab.com/wiggins.jonathan/plutus/api"
)

func Serve() {
    http.HandleFunc("/p", priceHandler)
    fmt.Println("Server listening on port 8000. Ctrl+c to quit.")
    log.Fatal(http.ListenAndServe(":8000", nil))
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
        // Return data in anonymous struct
        data := &struct{
            Name    string
            Price   float64
        }{
            Name    : ticker,
            Price   : price,
        }
        err = sendResponse(w, 200, data)
        if err != nil { fmt.Fprintf(w, "Error sending payload") }
    }
}

// Send a JSON response given a ResponseWriter, an HTTP status code, & payload
func sendResponse(w http.ResponseWriter, code int, payload interface{}) error {
    response, err := json.Marshal(payload)
    if err != nil { return err }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(code)
    w.Write(response)
    return nil
}
