// Code for REST API
package cmd

import (
    "fmt"
    "log"
    "net/http"
)

func ExecuteServer() {
    http.HandleFunc("/", helloWorld)
    log.Fatal(http.ListenAndServe(":5000", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
    fmt.Println("Endpoint Hit: HomePage")
}
