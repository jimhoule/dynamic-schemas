package models

type Document struct {
	Key  string         `json:"key"`
	Body map[string]any `json:"body"`
}