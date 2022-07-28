package models

import (
	"time"

	"github.com/lib/pq"
)

type Resource struct {
	Key     string      `json:"key"`
	Value   interface{} `json:"value"`
	Expires int64       `json:"expires,string"`
}

type ResourceV2 struct {
	Id        string         `json:"id"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Tags      pq.StringArray `json:"tags"`
}

type NewResourceTemplate struct {
	Title string
	Tags  []string
}

type UpdateResourceTemplate struct {
	Id    string
	Title string
	Tags  []string
}
