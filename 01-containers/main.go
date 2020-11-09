
package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

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