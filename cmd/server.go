// Code for web API
package cmd

import (
    "fmt"
    "log"
    "net/http"
)

func serve() {
    http.HandleFunc("/p", getDailyPrice)
    fmt.Println("Server listening on port 5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

func getDailyPrice(w http.ResponseWriter, r *http.Request) {
    // Take in the ticker & call GetPrice()
    price := getPrice("AMZN")
    fmt.Fprintf(w, "Price is %g", price)
}
