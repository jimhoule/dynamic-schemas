package entities

type Property struct {
	Type  string            `json:"type"`
	Items PropertyArrayItem `json:"items,omitempty"`
}
