package routes

import (
	"encoding/json"
	"fmt"
	"github.com/DanielHilton/service-catalogue/models"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var Catalogue *models.Catalogue

func RefreshCatalogueHandler(w http.ResponseWriter, _ *http.Request) {
	err := fetchCache()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to refresh Catalogue")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Catalogue updated")
}

func GetCatalogueHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if Catalogue == nil {
		err = fetchCache()
	}

	if err == nil {
		w.WriteHeader(http.StatusOK)
		catalogueJson, _ := json.Marshal(Catalogue)
		fmt.Fprint(w, string(catalogueJson))
		return
	}

	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprintf(w, "Failed to get catalogue.")
}

func GetServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serviceName := vars["serviceName"]
	service := models.GetEntryFromCatalogue(serviceName, Catalogue)

	serviceJson, _ := json.Marshal(service)
	fmt.Fprintf(w, string(serviceJson))
}

func fetchCache() error {
	var err error
	Catalogue, err = models.NewCatalogue(os.Getenv("CATALOGUE_URL"))
	return err
}
