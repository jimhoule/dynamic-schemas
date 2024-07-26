package models

type Property struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	IsRequired bool   `json:"isRequired"`
	ItemType   string `json:"itemType,omitempty"`
}
