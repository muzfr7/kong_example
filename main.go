package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Attach an elegant path with handler
	r.HandleFunc("/health/check", HealthcheckHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// HealthcheckHandler is a handler for healthcheck endpoint.
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, time.Now().String())
}
