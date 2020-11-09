package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	port    = 8080
	timeout = 2 * time.Second
)

func hello(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	log.Printf("Serving request with name as input: %s", name)
	fmt.Fprintf(w, "Hello, %s", name)
}

func hostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error fetching hostname")
		panic(err)
	}
	fmt.Fprintf(w, "%s", hostname)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/host", hostname)

	s := http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		ReadHeaderTimeout: timeout,
	}
	log.Printf("Starting server on port %d", port)
	log.Fatal(s.ListenAndServe())
}
