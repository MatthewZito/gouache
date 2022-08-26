package models

import (
	"encoding/json"
	"net/http"
)

// GouacheResponse represents a normalized, formatted response object.
type GouacheResponse struct {
	Internal string      `json:"internal,omitempty"`
	Friendly string      `json:"friendly,omitempty"`
	Data     interface{} `json:"data"`
	Flags    int         `json:"flags,omitempty"`
}

// DefaultOk generates a fallback 200 OK `GouacheResponse`.
func DefaultOk(data interface{}) []byte {
	v, _ := json.Marshal(GouacheResponse{
		Data:     data,
		Internal: "",
		Friendly: "",
		Flags:    0,
	})

	return v
}

// DefaultError generates a fallback erroneous `GouacheResponse`.
func DefaultError(internal string, friendly string, flags int) []byte {
	v, _ := json.Marshal(GouacheResponse{
		Data:     nil,
		Internal: internal,
		Friendly: friendly,
		Flags:    flags,
	})

	return v
}

// FormatError formats error metadata into an erroneous `GouacheResponse`.
func FormatError(w http.ResponseWriter, code int, internal string, friendly string, flags int) {

	// @todo handle
	v := DefaultError(internal, friendly, flags)

	FormatResponse(w, code, v)
}

// FormatResponse formats response metadata into a successful `GouacheResponse`.
func FormatResponse(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "gouache/auth")
	w.WriteHeader(code)
	w.Write(payload)
}
