package routes

import (
	"fmt"
	"github.com/DanielHilton/service-catalogue/models"
	"net/http"
	"os"
)

var Cache *models.Catalogue

func RefreshCatalogueHandler(w http.ResponseWriter, _ *http.Request) {
	err := fetchCache()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to refresh Cache")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Catalogue updated")
}

func GetCatalogueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if Cache == nil {
			fetchCache()
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Method Not Allowed")
}

func fetchCache() error {
	var err error
	Cache, err = models.NewCatalogue(os.Getenv("CATALOGUE_URL"))
	return err
}
