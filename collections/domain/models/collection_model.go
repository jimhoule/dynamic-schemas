package models

type Collection struct {
	Name       string      `json:"name"`
	Properties []*Property `json:"properties"`
}
