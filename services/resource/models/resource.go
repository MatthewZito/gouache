package models

import (
	"time"

	"github.com/lib/pq"
)

// Resource represents a state record in the resource service.
type Resource struct {
	Id        string         `json:"id"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Tags      pq.StringArray `json:"tags"`
}

// NewResourceTemplate represents the necessary fields to create a new resource record.
type NewResourceTemplate struct {
	Title string
	Tags  []string
}

// UpdateResourceTemplate represents the necessary fields to update an existing resource record.
type UpdateResourceTemplate struct {
	Id    string
	Title string
	Tags  []string
}
