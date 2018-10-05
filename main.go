package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"service-catalogue/catalogue"
)

var port = 4321
var cache *catalogue.Cache

func main() {
	fmt.Printf("Welcome to service catalogue backend.\n")

	cache = catalogue.NewCache(os.Getenv("CATALOGUE_URL"))
	if cache == nil {
		fmt.Printf("fuck")
		os.Exit(1)
	}
	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
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
