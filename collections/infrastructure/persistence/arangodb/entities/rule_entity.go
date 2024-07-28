package entities

type Rule struct {
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}
