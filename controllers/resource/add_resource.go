package controllers

import (
	"encoding/json"
	"net/http"

	resource "github.com/MatthewZito/gouache/models"
)

func AddResource(w http.ResponseWriter, r *http.Request) {
	rs := resource.Resource{}

	err := json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
