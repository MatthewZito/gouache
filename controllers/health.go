package controllers

import (
	"net/http"
	"os"

	format "github.com/MatthewZito/gouache/format"
)

// Health is a liveness check that returns the server's current status.
func Health(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()

	if err != nil {
		format.FormatError(w, http.StatusBadRequest, "Health check failed")
	}

	format.FormatResponse(w, http.StatusOK, map[string]string{"server": name, "result": "success"})
}
