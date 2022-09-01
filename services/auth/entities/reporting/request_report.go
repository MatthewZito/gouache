package entities

type RequestReport struct {
	Path       string                 `json:"path"`
	Method     string                 `json:"method"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Error      string                 `json:"error,omitempty"`
}
