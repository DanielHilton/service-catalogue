package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"service-catalogue/catalogue"
)

var port = 4321
var cache *catalogue.Cache

func main() {
	fmt.Printf("Welcome to service catalogue backend.\n")

	var err error
	cache, err = catalogue.NewCache(os.Getenv("CATALOGUE_URL"))
	if cache == nil {
		fmt.Printf("fuck")
		os.Exit(1)
	}
	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/refresh", RefreshCacheHandler)
	loggingRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), loggingRouter)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Failed to listen on %d", port))
	}
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	var buf bytes.Buffer
	for _, e := range cache.Entries {
		buf.WriteString(fmt.Sprintf("%s\n", e.Name))
	}

	fmt.Fprintf(w, buf.String())
}

func RefreshCacheHandler(w http.ResponseWriter, _ *http.Request) {
	var err error
	cache, err = catalogue.NewCache(os.Getenv("CATALOGUE_URL"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to refresh cache")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cache updated")
}
