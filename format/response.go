package format

import (
	"encoding/json"
	"net/http"
)

func DefaultOk() map[string]bool {
	return map[string]bool{"ok": true}
}

func FormatError(w http.ResponseWriter, code int, msg string) {
	FormatResponse(w, code, map[string]string{"error": msg})
}

func FormatResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Clacks-Overhead", "pending")
	w.Header().Set("X-Powered-By", "pending")
	w.WriteHeader(code)
	w.Write(response)
}
