package format

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Ok    bool        `json:"ok"`
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

func DefaultOk(data interface{}) []byte {
	v, _ := json.Marshal(Response{
		Ok:   true,
		Data: data,
	})

	fmt.Println(data)

	return v
}

func FormatError(w http.ResponseWriter, code int, msg string) {

	// @todo handle
	v, _ := json.Marshal(Response{
		Ok:    false,
		Error: msg,
	})

	FormatResponse(w, code, v)
}

func FormatResponse(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Clacks-Overhead", "pending")
	w.Header().Set("X-Powered-By", "pending")
	w.WriteHeader(code)
	w.Write(payload)
}
