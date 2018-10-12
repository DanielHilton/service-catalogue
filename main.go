package main

import (
	"fmt"
	"github.com/DanielHilton/service-catalogue/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var port = 4321

func main() {
	fmt.Printf("Welcome to service models backend.\n")

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/catalogue", routes.GetCatalogueHandler)
	router.HandleFunc("/catalogue/refresh", routes.RefreshCatalogueHandler)

	// Logging middleware
	loggingRouter := handlers.CombinedLoggingHandler(os.Stdout, router)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggingRouter)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Failed to listen on %d", port))
	}
}
