package controllers

import (
	"net/http"
	"os"

	format "github.com/MatthewZito/gouache/format"
)

// Health is a liveness check that returns the server's current status.
func Health(w http.ResponseWriter, r *http.Request) {

	if name, err := os.Hostname(); err != nil {
		format.FormatError(w, http.StatusBadRequest, "Health check failed")
	} else {

		format.FormatResponse(w, http.StatusOK, format.DefaultOk(map[string]string{"server": name}))
	}

}
