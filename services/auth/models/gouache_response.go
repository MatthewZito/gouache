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

// ToOk generates a fallback 200 OK `GouacheResponse`.
func ToOk(data interface{}) []byte {
	v, _ := json.Marshal(GouacheResponse{
		Data:     data,
		Internal: "",
		Friendly: "",
		Flags:    0,
	})

	return v
}

// ToException generates a fallback erroneous `GouacheResponse`.
func ToException(internal string, friendly string, flags int) []byte {
	v, _ := json.Marshal(GouacheResponse{
		Data:     nil,
		Internal: internal,
		Friendly: friendly,
		Flags:    flags,
	})

	return v
}

// SendGouacheException formats error metadata into an erroneous `GouacheResponse`.
func SendGouacheException(w http.ResponseWriter, code int, internal string, friendly string, flags int) {
	// @todo handle
	v := ToException(internal, friendly, flags)

	SendGouacheResponse(w, code, v)
}

// SendGouacheResponse formats response metadata into a successful `GouacheResponse`.
func SendGouacheResponse(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "gouache/auth")
	w.WriteHeader(code)
	w.Write(payload)
}
