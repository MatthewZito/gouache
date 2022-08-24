package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Internal string      `json:"internal,omitempty"`
	Friendly string      `json:"friendly,omitempty"`
	Data     interface{} `json:"data"`
	Flags    int         `json:"flags,omitempty"`
}

func DefaultOk(data interface{}) []byte {
	v, _ := json.Marshal(Response{
		Data: data,
	})

	fmt.Println(data)

	return v
}

func FormatError(w http.ResponseWriter, code int, internal string, friendly string, flags int) {

	// @todo handle
	v, _ := json.Marshal(Response{
		Internal: internal,
		Friendly: friendly,
		Flags:    flags,
	})

	FormatResponse(w, code, v)
}

func FormatResponse(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "gouache")
	w.WriteHeader(code)
	w.Write(payload)
}
