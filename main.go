package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"service-catalogue/catalogue"
)

func main() {
	fmt.Printf("Welcome to service catalogue backend.\n")
	c := catalogue.NewCache(os.Getenv("CATALOGUE_URL"))
	if c == nil {
		fmt.Printf("fuck")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":4321", router)
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}
