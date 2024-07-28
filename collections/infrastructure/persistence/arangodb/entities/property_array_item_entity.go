package entities

type PropertyArrayItem struct {
	Type    string `json:"type"`
	Maximum int    `json:"maximum,omitempty"`
}
