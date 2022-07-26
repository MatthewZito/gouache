package controllers

import (
	"fmt"
	"net/http"
	"time"

	format "github.com/MatthewZito/gouache/format"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixNano()
	format.FormatResponse(w, http.StatusOK, map[string]string{"time": fmt.Sprint(now)})
}
