package routes

import (
	"bytes"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	var buf bytes.Buffer
	for _, e := range Catalogue.Entries {
		buf.WriteString(fmt.Sprintf("%s\n", e.Name))
	}

	fmt.Fprintf(w, buf.String())
}
