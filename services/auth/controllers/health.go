package controllers

import (
	"net/http"
	"os"

	"github.com/exbotanical/gouache/models"
)

// Health is a liveness check that returns the server's current status.
func Health(w http.ResponseWriter, r *http.Request) {

	if name, err := os.Hostname(); err != nil {
		models.FormatError(w, http.StatusBadRequest, err.Error(), "Health check failed", 0)
	} else {

		models.FormatResponse(w, http.StatusOK, models.DefaultOk(map[string]string{"server": name}))
	}

}
