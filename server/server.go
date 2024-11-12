// Start a server to access the API over the net
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/wiggins-jonathan/plutus/api"
)

type ServerOption func(*server)

type server struct {
	Port  int
	Debug bool
}

// NewServer returns a server with optional defaults
func NewServer(options ...ServerOption) *server {
	server := &server{Port: 8080, Debug: false}

	for _, option := range options {
		option(server)
	}

	return server
}

// Serve serves the plutus API
func (s *server) Serve() error {
	http.HandleFunc("/p", priceHandler)

	port := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("Server listening at http://localhost%s - Ctrl+c to quit.\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		return fmt.Errorf("Failed to start server: %w", err)
	}

	return nil
}

func WithPort(port int) ServerOption {
	return func(s *server) {
		s.Port = port
	}
}

func WithDebug(debug bool) ServerOption {
	return func(s *server) {
		s.Debug = debug
	}
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
		data := &struct {
			Name  string
			Price float64
		}{
			Name:  ticker,
			Price: price,
		}
		err = sendResponse(w, 200, data)
		if err != nil {
			fmt.Fprintf(w, "Error sending payload")
		}
	}
}

// Send a JSON response given a ResponseWriter, an HTTP status code, & payload
func sendResponse(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}
