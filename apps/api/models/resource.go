package models

type Resource struct {
	Key     string      `json:"key"`
	Value   interface{} `json:"value"`
	Expires int64       `json:"expires,string"`
}
